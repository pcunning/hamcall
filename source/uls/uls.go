package uls

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Println("Downloading ULS data")

	err := downloader.FetchFtp("tmp/amat.zip", "ftp://wirelessftp.fcc.gov:21/pub/uls/complete/l_amat.zip")
	if err != nil {
		return err
	}

	files, err := downloader.Unzip("tmp/amat.zip", "tmp/amat")
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
	fmt.Println("processing AM")

	f, err := os.Open("tmp/amat/AM.dat")
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

		hc := data.HamCall{
			Callsign: record[4],
			Class:    record[5],
		}

		// fmt.Print("AM-", record[4], " ")
		data.UpdateMap(calls, &hc, record[4])
	}
}

func ProcessEN(calls *map[string]data.HamCall) {
	fmt.Println("processing EN")

	f, err := os.Open("tmp/amat/EN.dat")
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

		hc := data.HamCall{
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

		// fmt.Print("EN-", record[4], " ")
		data.UpdateMap(calls, &hc, record[4])
	}

}

func ProcessHD(calls *map[string]data.HamCall) {
	fmt.Println("processing HD")

	f, err := os.Open("tmp/amat/HD.dat")
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

		hc := data.HamCall{
			Callsign:   record[4],
			Grant:      record[7],
			Expiration: record[8],
			FileNumber: record[2], // this is missing in the ULS data for some HV records - Is it in applications download?
			Effective:  record[42],
		}

		// fmt.Print("HD-", record[4], " ")
		data.UpdateMap(calls, &hc, record[4])
	}

}
