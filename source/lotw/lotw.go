package lotw

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Downloading lotw data")
	err := downloader.FetchHttp("tmp/lotw.csv", "https://lotw.arrl.org/lotw-user-activity.csv")
	if err != nil {
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall) {

	fmt.Println("processing LOTW")

	f, err := os.Open("tmp/lotw.csv")
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

		hc := data.HamCall{
			Callsign: record[0],
			LOTW:     record[1] + record[2],
		}
		data.UpdateMap(calls, &hc, record[0])
	}
}
