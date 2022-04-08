package lotw

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Downloading lotw data")
	err := downloader.FetchHttp("lotw.csv", "https://lotw.arrl.org/lotw-user-activity.csv")
	if err != nil {
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing LOTW")

	f, err := os.Open("lotw.csv")
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

		call := record[0]
		item, c := (*calls)[call]
		if c {
			item.LOTW = record[1] + record[2]
		} else {
			item = data.HamCall{
				Callsign: call,
				LOTW:     record[1] + record[2],
			}
		}
		(*calls)[call] = item
	}

	fmt.Printf(" ... %s\n", time.Since(start).String())

}
