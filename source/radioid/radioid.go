package radioid

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

func Download(wg *sync.WaitGroup, backup *downloader.BackupDownloader) error {
	defer wg.Done()
	fmt.Println("Downloading radioid data")
	
	var err error
	if backup != nil {
		err = downloader.FetchWithBackup("dmrid.dat", "https://www.radioid.net/static/dmrid.dat", "dmrid.dat", backup)
	} else {
		err = downloader.FetchHttp("dmrid.dat", "https://www.radioid.net/static/dmrid.dat")
	}
	
	if err != nil {
		fmt.Printf("Warning: Error downloading RadioID data: %v\n", err)
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall, partialMode bool) {
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
			// Update existing callsign
			item.DMRID = append(item.DMRID, id)
			(*calls)[call] = item
		} else if !partialMode {
			// Only create new callsign if not in partial mode
			item = data.HamCall{
				Callsign: call,
				DMRID:    []int{id},
			}
			(*calls)[call] = item
		}
		// In partial mode, skip callsigns that don't exist in ULS data
	}
	fmt.Printf(" ... %s\n", time.Since(start).String())
}
