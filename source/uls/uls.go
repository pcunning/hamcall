package uls

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Println("Downloading ULS data")

	err := downloader.FetchFtp("amat.zip", "ftp://wirelessftp.fcc.gov:21/pub/uls/complete/l_amat.zip")
	if err != nil {
		return err
	}

	files, err := downloader.Unzip("amat.zip", "amat")
	if err != nil {
		return err
	}

	fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

	return nil
}

func Process(calls *map[string]data.HamCall) {
	ProcessAM(calls)
	ProcessEN(calls)
	ProcessHD(calls)
}

func ProcessAM(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing AM")

	f, err := os.Open("amat/AM.dat")
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '|'
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

		call := record[4]
		item, c := (*calls)[call]
		if c {
			item.Callsign = record[4]
			item.Class = record[5]
		} else {
			item = data.HamCall{
				Callsign: record[4],
				Class:    record[5],
			}
		}
		(*calls)[call] = item
	}
	fmt.Printf(" ... %s\n", time.Since(start).String())
}

func ProcessEN(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing EN")

	f, err := os.Open("amat/EN.dat")
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '|'
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

		call := record[4]
		item, c := (*calls)[call]
		if c {
			item.Callsign = record[4]
			item.Name = record[7]
			item.FirstName = record[8]
			item.Mi = record[9]
			item.LastName = record[10]
			item.Address = record[15]
			item.City = record[16]
			item.State = record[17]
			item.Zip = record[18]
			item.PoBox = record[19]
			item.LicenseKey = record[1]
			item.FRN = record[22]
		} else {
			item = data.HamCall{
				Callsign:   record[4],
				Name:       record[7],
				FirstName:  record[8],
				Mi:         record[9],
				LastName:   record[10],
				Address:    record[15],
				City:       record[16],
				State:      record[17],
				Zip:        record[18],
				PoBox:      record[19],
				LicenseKey: record[1],
				FRN:        record[22],
			}
		}
		(*calls)[call] = item
	}

	fmt.Printf(" ... %s\n", time.Since(start).String())
}

func ProcessHD(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing HD")

	f, err := os.Open("amat/HD.dat")
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '|'
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

		call := record[4]
		item, c := (*calls)[call]
		if c {
			item.Callsign = record[4]
			item.Grant = record[7]
			item.Expiration = record[8]
			item.FileNumber = record[2] // this is missing in the ULS data for some HV records - Is it in applications download?
			item.Effective = record[42]
		} else {
			item = data.HamCall{
				Callsign:   record[4],
				Grant:      record[7],
				Expiration: record[8],
				FileNumber: record[2],
				Effective:  record[42],
			}
		}
		(*calls)[call] = item
	}
	fmt.Printf(" ... %s\n", time.Since(start).String())

}
