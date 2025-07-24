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

	// Check the ULS processing mode
	ulsMode := os.Getenv("ULS_MODE")

	switch ulsMode {
	case "partial":
		wg.Add(2)
		go DownloadDailyLicenses(wg)
		go DownloadDailyApplications(wg)
	case "full":
		wg.Add(4)
		go DownloadLicenses(wg)
		go DownloadApplications(wg)
		go DownloadDailyLicenses(wg)
		go DownloadDailyApplications(wg)
	default:
		// Default to weekly processing
		wg.Add(2)
		go DownloadLicenses(wg)
		go DownloadApplications(wg)
	}

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

// getPreviousBusinessDay returns the day code for the most recent business day with daily files
// Daily files are published the day after, so Monday's file is available on Tuesday after noon
func getPreviousBusinessDay() string {
	now := time.Now()
	
	// Get yesterday's day
	yesterday := now.AddDate(0, 0, -1)
	dayOfWeek := yesterday.Weekday()
	
	// Handle weekends - if yesterday was Sunday or Saturday, get Friday
	switch dayOfWeek {
	case time.Sunday:
		// Yesterday was Sunday, get Friday's file (2 days back)
		yesterday = yesterday.AddDate(0, 0, -2)
	case time.Saturday:
		// Yesterday was Saturday, get Friday's file (1 day back)
		yesterday = yesterday.AddDate(0, 0, -1)
	}
	
	// Convert to 3-letter lowercase day code
	day := yesterday.Format("Mon")
	return strings.ToLower(day[:3])
}

// getAllDailysSinceWeekly returns all day codes for daily files since the last weekly
// Weekly files are published on Sunday, so get Monday through the previous business day
func getAllDailysSinceWeekly() []string {
	var days []string
	now := time.Now()
	
	// Find the most recent Sunday (when weekly was published)
	daysBack := int(now.Weekday())
	if daysBack == 0 {
		daysBack = 7 // If today is Sunday, go back to previous Sunday
	}
	lastSunday := now.AddDate(0, 0, -daysBack)
	
	// Collect all business days from Monday after last Sunday to yesterday
	for d := lastSunday.AddDate(0, 0, 1); d.Before(now); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			dayStr := d.Format("Mon")
			days = append(days, strings.ToLower(dayStr[:3]))
		}
	}
	
	return days
}
func DownloadDailyLicenses(wg *sync.WaitGroup) error {
	defer wg.Done()

	ulsMode := os.Getenv("ULS_MODE")
	var daysToDownload []string
	
	if ulsMode == "full" {
		fmt.Println("Downloading ULS Daily License data (full mode - all dailies since weekly)")
		daysToDownload = getAllDailysSinceWeekly()
		if len(daysToDownload) == 0 {
			fmt.Println("No daily files to download in full mode")
			return nil
		}
	} else {
		fmt.Println("Downloading ULS Daily License data (partial mode)")
		day := getPreviousBusinessDay()
		daysToDownload = []string{day}
	}

	for _, day := range daysToDownload {
		dailyUrl := fmt.Sprintf("ftp://wirelessftp.fcc.gov:21/pub/uls/daily/l_am_%s.zip", day)
		dailyFileName := fmt.Sprintf("l_am_%s.zip", day)

		fmt.Printf("Downloading daily license file for %s...\n", day)
		err := downloader.FetchFtp(dailyFileName, dailyUrl)
		if err != nil {
			fmt.Printf("Failed to download daily license file for %s, skipping: %v\n", day, err)
			continue
		}

		// For full mode, extract to separate directory per day, then merge
		extractDir := "l_amat"
		if ulsMode == "full" {
			extractDir = fmt.Sprintf("l_amat_daily_%s", day)
		}

		files, err := downloader.Unzip(dailyFileName, extractDir)
		if err != nil {
			fmt.Printf("Failed to unzip daily license file for %s, skipping: %v\n", day, err)
			continue
		}

		fmt.Printf("Daily License files for %s unzipped:\n%s\n", day, strings.Join(files, "\n"))
	}

	// For partial mode, check if daily files contain essential data and fallback if needed
	if ulsMode == "partial" {
		if !dailyFilesContainEssentialData("l_amat") {
			fmt.Println("Daily files missing essential data, downloading weekly files...")
			wg.Add(1)
			return DownloadLicenses(wg)
		}
	}

	// For full mode, merge all daily files with the weekly files
	if ulsMode == "full" {
		err := mergeDailyFiles("l_amat", daysToDownload, "license")
		if err != nil {
			fmt.Printf("Failed to merge daily license files: %v\n", err)
		}
	}

	return nil
}

func DownloadDailyApplications(wg *sync.WaitGroup) error {
	defer wg.Done()

	ulsMode := os.Getenv("ULS_MODE")
	var daysToDownload []string
	
	if ulsMode == "full" {
		fmt.Println("Downloading ULS Daily Application data (full mode - all dailies since weekly)")
		daysToDownload = getAllDailysSinceWeekly()
		if len(daysToDownload) == 0 {
			fmt.Println("No daily application files to download in full mode")
			return nil
		}
	} else {
		fmt.Println("Downloading ULS Daily Application data (partial mode)")
		day := getPreviousBusinessDay()
		daysToDownload = []string{day}
	}

	for _, day := range daysToDownload {
		dailyUrl := fmt.Sprintf("ftp://wirelessftp.fcc.gov:21/pub/uls/daily/a_am_%s.zip", day)
		dailyFileName := fmt.Sprintf("a_am_%s.zip", day)

		fmt.Printf("Downloading daily application file for %s...\n", day)
		err := downloader.FetchFtp(dailyFileName, dailyUrl)
		if err != nil {
			fmt.Printf("Failed to download daily application file for %s, skipping: %v\n", day, err)
			continue
		}

		// For full mode, extract to separate directory per day, then merge
		extractDir := "a_amat"
		if ulsMode == "full" {
			extractDir = fmt.Sprintf("a_amat_daily_%s", day)
		}

		files, err := downloader.Unzip(dailyFileName, extractDir)
		if err != nil {
			fmt.Printf("Failed to unzip daily application file for %s, skipping: %v\n", day, err)
			continue
		}

		fmt.Printf("Daily Application files for %s unzipped:\n%s\n", day, strings.Join(files, "\n"))
	}

	// For full mode, merge all daily files with the weekly files
	if ulsMode == "full" {
		err := mergeDailyFiles("a_amat", daysToDownload, "application")
		if err != nil {
			fmt.Printf("Failed to merge daily application files: %v\n", err)
		}
	}

	return nil
}

// mergeDailyFiles merges daily files with weekly files in full mode
func mergeDailyFiles(baseDir string, days []string, fileType string) error {
	essentialFiles := []string{"AM.dat", "EN.dat", "HD.dat"}
	if fileType == "application" {
		essentialFiles = []string{"EN.dat", "HS.dat"}
	}
	
	for _, filename := range essentialFiles {
		baseFile := baseDir + "/" + filename
		
		// Open base file for appending
		baseFileHandle, err := os.OpenFile(baseFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Warning: Could not open base file %s for appending: %v\n", baseFile, err)
			continue
		}
		defer baseFileHandle.Close()
		
		// Append each daily file
		for _, day := range days {
			dailyFile := fmt.Sprintf("%s_daily_%s/%s", baseDir, day, filename)
			
			dailyFileHandle, err := os.Open(dailyFile)
			if err != nil {
				fmt.Printf("Warning: Could not open daily file %s: %v\n", dailyFile, err)
				continue
			}
			
			// Copy daily file content to base file
			_, err = io.Copy(baseFileHandle, dailyFileHandle)
			dailyFileHandle.Close()
			
			if err != nil {
				fmt.Printf("Warning: Could not merge daily file %s: %v\n", dailyFile, err)
			} else {
				fmt.Printf("Merged daily file %s into %s\n", dailyFile, baseFile)
			}
		}
	}
	
	return nil
}

func dailyFilesContainEssentialData(dir string) bool {
	// Check if essential files exist and have content
	essentialFiles := []string{"AM.dat", "EN.dat", "HD.dat"}
	
	for _, filename := range essentialFiles {
		filepath := dir + "/" + filename
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			fmt.Printf("Essential file %s missing from daily download\n", filename)
			return false
		}
		
		// Check if file has content
		info, err := os.Stat(filepath)
		if err != nil || info.Size() == 0 {
			fmt.Printf("Essential file %s is empty or unreadable\n", filename)
			return false
		}
	}
	
	return true
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
