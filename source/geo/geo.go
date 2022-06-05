package geo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Downloading GEO data")
	err := downloader.FetchHttp("ham-stations.csv", os.Getenv("GEO_URL"))
	if err != nil {
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing GEO")

	f, err := os.Open("ham-stations.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		lat, err := strconv.ParseFloat(record[9], 64)
		if err != nil {
			continue
		}

		lon, err := strconv.ParseFloat(record[10], 64)
		if err != nil {
			continue
		}

		loc := data.Location{
			Latitude:  lat,
			Longitude: lon,
		}

		call := record[0]
		item, c := (*calls)[call]
		if c {
			item.Location = &loc
		} else {
			item = data.HamCall{
				Callsign: call,
				Location: &loc,
			}
		}
		(*calls)[call] = item

	}
	fmt.Printf(" ... %s\n", time.Since(start).String())

}
