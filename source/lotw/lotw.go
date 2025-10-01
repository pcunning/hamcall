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

func Download(wg *sync.WaitGroup, backup *downloader.BackupDownloader) error {
	defer wg.Done()
	fmt.Println("Downloading lotw data")
	
	var err error
	if backup != nil {
		err = downloader.FetchWithBackup("lotw.csv", "https://lotw.arrl.org/lotw-user-activity.csv", "lotw.csv", backup)
	} else {
		err = downloader.FetchHttp("lotw.csv", "https://lotw.arrl.org/lotw-user-activity.csv")
	}
	
	if err != nil {
		fmt.Printf("Warning: Error downloading LOTW data: %v\n", err)
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall, partialMode bool) {
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
			// Update existing callsign
			item.LOTW = record[1] + record[2]
			(*calls)[call] = item
		} else if !partialMode {
			// Only create new callsign if not in partial mode
			item = data.HamCall{
				Callsign: call,
				LOTW:     record[1] + record[2],
			}
			(*calls)[call] = item
		}
		// In partial mode, skip callsigns that don't exist in ULS data
	}

	fmt.Printf(" ... %s\n", time.Since(start).String())
}
