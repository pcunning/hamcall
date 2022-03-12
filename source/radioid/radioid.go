package radioid

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Downloading radioid data")
	err := downloader.FetchHttp("tmp/dmrid.dat", "https://www.radioid.net/static/dmrid.dat")
	if err != nil {
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall) {
	f, err := os.Open("tmp/dmrid.dat")
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

		call := record[1]
		item, c := (*calls)[call]
		if c {
			id, err := strconv.Atoi(record[0])
			if err == nil {
				item.DMRID = append(item.DMRID, id)
				(*calls)[call] = item
			}
		}
	}
}
