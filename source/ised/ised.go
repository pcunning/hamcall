package ised

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) (string, error) {
	defer wg.Done()
	fmt.Println("Downloading ISED (Canada) data")
	err := downloader.FetchHttp("ised.zip", "https://apc-cap.ic.gc.ca/datafiles/amateur_delim.zip")
	if err != nil {
		log.Fatalf("Error downloading ISED (Canada) data: %v", err)
	}

	_, err = downloader.Unzip("ised.zip", "ised_data")
	if err != nil {
		log.Fatalf("Error unzipping ISED (Canada) data: %v", err)
	}

	return "ised_data/amateur_delim.txt", nil
}

func ProcessLine(line string) data.HamCall {
	columns := strings.Split(line, ";")
	var item data.HamCall
	if columns[0] != "callsign" {
		var class string
		switch {
		case columns[7] != "" && columns[11] == "":
			class = "Basic"
		case columns[11] != "" && columns[10] == "":
			class = "Basic with Honors"
		case columns[10] != "":
			class = "Advanced"
		default:
			class = "Unknown"
		}
		if columns[9] != "" || columns[8] != "" {
			class += " (w/ Morse)"
		}
		name := columns[1] + " " + columns[2]
		item = data.HamCall{
			Callsign:  columns[0],
			Name:      name,
			FirstName: columns[1],
			LastName:  columns[2],
			Address:   columns[3],
			City:      columns[4],
			State:     columns[5],
			Zip:       columns[6],
			Class:     class,
		}
	}
	return item
}

func Process(calls *map[string]data.HamCall, path string) {
	start := time.Now()
	fmt.Print("processing ISED data")

	// Step 2: Open the new file.
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var line string
	var call data.HamCall

	// Step 3: Process Each Line.
	for scanner.Scan() {
		line = scanner.Text()
		call = ProcessLine(line)
		(*calls)[call.Callsign] = call
	}

	fmt.Printf(" ... %s\n", time.Since(start).String())
}
