package radioid

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Downloading radioid data")
	err := downloader.FetchHttp("dmrid.dat", "https://www.radioid.net/static/dmrid.dat")
	if err != nil {
		log.Fatalf("Error downloading RadioID data: %v", err)
	}
	return nil
}

func Process(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing radioID")

	f, err := os.Open("dmrid.dat")
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		call := record[1]
		item, c := (*calls)[call]
		if c {
			item.DMRID = append(item.DMRID, strconv.Itoa(id))
		} else {
			item = data.HamCall{
				Callsign: call,
				DMRID:    []string{strconv.Itoa(id)},
			}

		}
		(*calls)[call] = item
	}
	fmt.Printf(" ... %s\n", time.Since(start).String())
}
