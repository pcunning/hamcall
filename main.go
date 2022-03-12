package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
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
	"github.com/pcunning/hamcall/source/lotw"
	"github.com/pcunning/hamcall/source/radioid"
	"github.com/pcunning/hamcall/source/uls"
)

func main() {
	start := time.Now()

	downloadDataFiles := flag.Bool("dl", false, "skip downloading files")
	runMode := flag.String("m", "cli", "run mode: b2, cli, or web")
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
	} else {
		fmt.Println("skipping downloading files (run with -dl to download)")
	}

	calls := make(map[string]data.HamCall)
	process(&calls)

	if *runMode == "b2" {
		writeToB2(&calls, keyID, applicationKey, uploadWorkers, osSigExit)
	}

	if *runMode == "cli" {
		cli(&calls)
	}

	fmt.Printf("total runtime %s\n", time.Since(start).String())
}

func downloadFiles() {
	var wg sync.WaitGroup

	wg.Add(3)

	go uls.Download(&wg)
	go radioid.Download(&wg)
	go lotw.Download(&wg)

	wg.Wait()
}

func process(calls *map[string]data.HamCall) {
	uls.Process(calls)
	radioid.Process(calls)
	lotw.Process(calls)
}

func writeToB2(calls *map[string]data.HamCall, keyID, applicationKey string, uploadWorkers int, osSigExit chan bool) {

	b, err := b2.New(keyID, applicationKey, uploadWorkers)
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

		if !usCall.MatchString(strings.ToUpper(input)) {
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
			log.Fatalf(err.Error())
		}
		fmt.Printf("%s\n", string(j))
	}
}
