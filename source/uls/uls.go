package uls

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()

	// Daily files for future reference
	// https: //data.fcc.gov/download/pub/uls/daily/a_am_sun.zip
	// https: //data.fcc.gov/download/pub/uls/daily/l_am_sun.zip

	wg.Add(2)
	go DownloadLicenses(wg)
	go DownloadApplications(wg)

	return nil
}

func DownloadLicenses(wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Println("Downloading ULS License data")

	err := downloader.FetchFtp("l_amat.zip", "ftp://wirelessftp.fcc.gov:21/pub/uls/complete/l_amat.zip")
	if err != nil {
		log.Fatalf("Error downloading ULS license data: %v", err)
	}

	files, err := downloader.Unzip("l_amat.zip", "l_amat")
	if err != nil {
		return err
	}

	fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

	return nil
}

func DownloadApplications(wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Println("Downloading ULS Application data")

	err := downloader.FetchFtp("a_amat.zip", "ftp://wirelessftp.fcc.gov:21/pub/uls/complete/a_amat.zip")
	if err != nil {
		log.Fatalf("Error downloading ULS application data: %v", err)
	}

	files, err := downloader.Unzip("a_amat.zip", "a_amat")
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
	LoadFileNumbers(calls)
}

func ProcessAM(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing AM")

	f, err := os.Open("l_amat/AM.dat")
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

	f, err := os.Open("l_amat/EN.dat")
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		s := string(line)
		record := strings.Split(s, "|")

		if len(record) < 25 {
			fmt.Printf("record too short: %v\n", record)
			continue
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

	f, err := os.Open("l_amat/HD.dat")
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

func LoadFileNumbers(calls *map[string]data.HamCall) {
	start := time.Now()
	fmt.Print("processing Application Data for File Numbers")

	f, err := os.Open("a_amat/HS.dat")
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '|'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	apgrt := make(map[string]bool)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		if record[5] == "APGRT " {
			apgrt[record[2]] = true
		}
	}

	f2, err := os.Open("a_amat/EN.dat")
	if err != nil {
		return
	}
	defer f2.Close()

	r = csv.NewReader(f2)
	r.Comma = '|'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	fns := make(map[string]string)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		if len(record) < 25 {
			continue
		}

		if apgrt[record[2]] {
			fns[record[22]] = record[2]
		}
	}

	for call, c := range *calls {
		if c.FileNumber == "" {
			c.FileNumber = fns[c.FRN]
			(*calls)[call] = c
		}
	}

	fmt.Printf(" ... %s\n", time.Since(start).String())

}
