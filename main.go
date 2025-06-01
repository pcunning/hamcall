package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/pcunning/hamcall/b2"
	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/source/geo"
	"github.com/pcunning/hamcall/source/ised"
	"github.com/pcunning/hamcall/source/lotw"
	"github.com/pcunning/hamcall/source/radioid"
	"github.com/pcunning/hamcall/source/uls"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	start := time.Now()

	downloadDataFiles := flag.Bool("dl", false, "skip downloading files")
	runMode := flag.String("m", "cli", "run mode: cli, stats, b2, or web")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	osSigExit := make(chan bool, 1)

	go func() {
		<-sigs
		osSigExit <- true
	}()

	keyID := os.Getenv("B2_KEYID")
	applicationKey := os.Getenv("B2_APPKEY")
	uploadWorkers := 200

	if *downloadDataFiles {
		downloadFiles()
		fmt.Printf("download finished at %s\n", time.Since(start).String())

	} else {
		fmt.Println("skipping downloading files (run with -dl to download)")
	}

	calls := make(map[string]data.HamCall)
	process(&calls)
	fmt.Printf("processing finished at %s\n", time.Since(start).String())

	if *runMode == "b2" || *runMode == "stats" {
		dryRun := true
		if *runMode == "b2" {
			dryRun = false
		}
		writeToB2(&calls, keyID, applicationKey, uploadWorkers, osSigExit, dryRun)
	}

	if *runMode == "cli" || *runMode == "stats" {
		cli(&calls)
	}

	if *runMode == "web" {
		web(&calls, osSigExit)
	}

	if *runMode == "sqlite" {
		qslite(&calls, osSigExit)
	}

	fmt.Printf("total runtime %s\n", time.Since(start).String())
}

func downloadFiles() {
	var wg sync.WaitGroup

	wg.Add(5)

	go uls.Download(&wg)
	go ised.Download(&wg)
	go radioid.Download(&wg)
	go lotw.Download(&wg)
	go geo.Download(&wg)

	wg.Wait()
}

func process(calls *map[string]data.HamCall) {
	uls.Process(calls)
	ised.Process(calls, "ised_data/amateur_delim.txt")
	radioid.Process(calls)
	lotw.Process(calls)
	geo.Process(calls)
}

func writeToB2(calls *map[string]data.HamCall, keyID, applicationKey string, uploadWorkers int, osSigExit chan bool, dryRun bool) {

	b, err := b2.New(keyID, applicationKey, uploadWorkers, dryRun)
	if err != nil {
		fmt.Println("error creating b2 client: ", err)
		return
	}

	err = b.Write(calls, osSigExit)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func cli(calls *map[string]data.HamCall) {
	validate := func(input string) error {
		var usCall = regexp.MustCompile(`^[AKNW][A-Z]{0,2}[0123456789][A-Z]{1,3}$`)

		var caCall = regexp.MustCompile(`^(?:VE|VA|VB)[0123456789][A-Z]{1,3}$`)

		if !usCall.MatchString(strings.ToUpper(input)) && !caCall.MatchString(strings.ToUpper(input)) {
			return errors.New("Invalid callsign")
		}

		return nil
	}

	for {
		prompt := promptui.Prompt{
			Label:    "Callsign",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		j, err := json.MarshalIndent((*calls)[strings.ToUpper(result)], "", "  ")
		if err != nil {
			log.Fatalf("error marshaling JSON: %v", err)
		}
		fmt.Printf("%s\n", string(j))
	}
}

func web(calls *map[string]data.HamCall, osSigExit chan bool) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call := r.URL.Path[len("/") : len(r.URL.Path)-len(".json")]
		fmt.Printf("%s: %s\n", r.Method, r.URL.Path)
		j, err := json.Marshal((*calls)[strings.ToUpper(call)])
		if err != nil {
			// log.Fatalf(err.Error())
		}
		fmt.Fprintf(w, "%s", j)
	})

	server := &http.Server{Addr: ":8080", Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			// handle err
		}
	}()

	<-osSigExit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// handle err
	}

}

func qslite(calls *map[string]data.HamCall, osSigExit chan bool) {
	os.Remove("hamcall.db")

	fmt.Println("Creating hamcall.db...")
	file, err := os.Create("hamcall.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("hamcall.db created")

	db, _ := sql.Open("sqlite3", "./hamcall.db")
	defer db.Close()

	create, err := db.Prepare(`CREATE TABLE hamcall (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,	
		"callsign" TEXT,
		"name" TEXT,
		"address" TEXT,
		"city" TEXT,
		"state" TEXT,
		"zip" TEXT,
		"po_box" TEXT,
		"license_key" TEXT,
		"frn" TEXT,
		"grant" TEXT,
		"expiration" TEXT,
		"file_number" TEXT,
		"effective" TEXT,
		"class" TEXT	
	  );`)
	if err != nil {
		log.Fatal(err.Error())
	}
	create.Exec()
	log.Println("created call table")

	index, err := db.Prepare(`CREATE UNIQUE INDEX callsign_index ON hamcall (callsign);`)
	if err != nil {
		log.Fatal(err.Error())
	}
	index.Exec()
	log.Println("created callsign index")

	sync, err := db.Prepare(`PRAGMA synchronous = OFF;`)
	if err != nil {
		log.Fatal(err.Error())
	}
	sync.Exec()
	log.Println("turned off synchronous mode")

	journal, err := db.Prepare(`PRAGMA journal_mode = MEMORY;`)
	if err != nil {
		log.Fatal(err.Error())
	}
	journal.Exec()
	log.Println("set journal mode to memory")

	insertStudentSQL := `INSERT INTO hamcall(
		callsign,
		name,
		address,
		city,
		state,
		zip,
		po_box,
		license_key,
		frn,
		grant,
		expiration,
		file_number,
		effective,
		class
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	insert, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	go func() {

		i := 0
		for _, call := range *calls {

			_, err = insert.Exec(
				call.Callsign,
				call.Name,
				call.Address,
				call.City,
				call.State,
				call.Zip,
				call.PoBox,
				call.LicenseKey,
				call.FRN,
				call.Grant,
				call.Expiration,
				call.FileNumber,
				call.Effective,
				call.Class,
			)
			if err != nil {
				log.Fatalln(err.Error())
			}

			i++
			if i%100000 == 0 {
				log.Printf("%d calls inserted\n", i)
			}
		}
		osSigExit <- true
	}()
	<-osSigExit
}
