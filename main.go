package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha1"
	"encoding/csv"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/imdario/mergo"
	"github.com/jlaffaye/ftp"
	"github.com/puzpuzpuz/xsync"
	"gopkg.in/kothar/go-backblaze.v0"
)

type HamCall struct {
	Callsign   string `json:"callsign"`
	Name       string `json:"name"`
	FirstName  string `json:"first_name"`
	Mi         string `json:"mi"`
	LastName   string `json:"last_name"`
	Class      string `json:"class"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zip        string `json:"zip"`
	PoBox      string `json:"po_box"`
	Grant      string `json:"grant"`
	Effective  string `json:"effective"`
	Expiration string `json:"expiration"`
	FRN        string `json:"frn"`
	FileNumber string `json:"file_number"`
	LOTW       string `json:"last_lotw"`
	LicenseKey string `json:"license_key"`
}

type HashTable map[string]string
type Upload struct {
	FileName    string
	FileContent *[]byte
}

func main() {

	uploadWorkersCount := 200

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	start := time.Now()

	keyID := os.Getenv("B2_KEYID")
	applicationKey := os.Getenv("B2_APPKEY")

	b2, err := backblaze.NewB2(backblaze.Credentials{
		KeyID:          keyID,
		ApplicationKey: applicationKey,
	})
	if err != nil {
		log.Fatal(err)
	}

	b2b, err := b2.Bucket("hamcall")
	if err != nil {
		log.Fatal(err)
	}

	var ht HashTable
	err = LoadHashTable(&ht, b2b)
	if err != nil {
		fmt.Printf("Error loading hash table: %s\n", err)
		ht = make(HashTable)
	}

	xht := xsync.NewMap()
	for k, v := range ht {
		xht.Store(k, v)
	}

	go func() {
		<-sigs
		done <- true
	}()

	go func() {
		fmt.Printf("%s: ", time.Since(start))
		fmt.Printf("%d entries in hash table\n", len(ht))

		var wg sync.WaitGroup

		wg.Add(1)
		go DownloadFile("lotw.csv", "https://lotw.arrl.org/lotw-user-activity.csv", &wg)

		wg.Add(1)
		go DownloadFile("dmrid.dat", "https://www.radioid.net/static/dmrid.dat", &wg)

		wg.Add(1)
		go DownloadFTPFile("amat.zip", "ftp://wirelessftp.fcc.gov:21/pub/uls/complete/l_amat.zip", &wg)

		wg.Wait()

		fmt.Println(time.Since(start))
		files, err := Unzip("amat.zip", "amat")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

		calls := make(map[string]HamCall)

		wg.Add(1)
		fmt.Printf("%s: ", time.Since(start))
		ProcessAM(&calls, &wg)

		wg.Add(1)
		fmt.Printf("%s: ", time.Since(start))
		ProcessEN(&calls, &wg)

		wg.Add(1)
		fmt.Printf("%s: ", time.Since(start))
		ProcessHD(&calls, &wg)

		wg.Add(1)
		fmt.Printf("%s: ", time.Since(start))
		ProcessLOTW(&calls, &wg)

		wg.Wait()

		// Create workers
		fmt.Printf("%s: creating workers\n", time.Since(start))
		uploadTasks := make(chan Upload, uploadWorkersCount*2)
		group := sync.WaitGroup{}
		for i := 0; i < uploadWorkersCount; i++ {
			group.Add(1)
			go UploadWorker(uploadTasks, xht, b2b, &group)
		}

		fmt.Printf("%s: Writing %d files to disk\n", time.Since(start), len(calls))
		i := 1
		batchTime := time.Now()
		for _, v := range calls {
			err = WriteCall(&v, xht, uploadTasks)
			if err != nil {
				log.Fatal(err)
			}
			if i%10000 == 0 {
				rps := 10000 / time.Since(batchTime).Seconds()
				etr := time.Duration((float64(len(calls)-i) / rps)) * time.Second
				fmt.Printf("%s: %d... %s = %.2f/second etr: %s\n", time.Since(start), i, time.Since(batchTime), rps, etr)
				batchTime = time.Now()
			}
			i++
		}

		close(uploadTasks)
		group.Wait()

		fmt.Printf("%s: ", time.Since(start))
		done <- true

	}()

	<-done
	fmt.Printf("\n\nexiting... please wait while saving hash table\n\n")
	xht.Range(func(k string, v interface{}) bool {
		ht[k] = v.(string)
		return true
	})

	SaveHashTable(&ht, b2b)
	fmt.Println(time.Since(start))
}

func DownloadFile(filepath string, u string, wg *sync.WaitGroup) {
	// DownloadFile will download a url to a local file. It's efficient because it will
	// write as it downloads and not load the whole file into memory.
	// From https://golangcode.com/download-a-file-from-a-url/
	defer wg.Done()

	// Get the data
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadFTPFile(filepath string, u string, wg *sync.WaitGroup) {
	defer wg.Done()

	pu, _ := url.Parse(u)

	c, err := ftp.Dial(pu.Host, ftp.DialWithTimeout(5120*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Retr(pu.Path)
	if err != nil {
		log.Fatal(err)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

func Unzip(src string, dest string) ([]string, error) {
	// Unzip will decompress a zip archive, moving all files and folders
	// within the zip file (parameter 1) to an output directory (parameter 2).
	// https://golangcode.com/unzip-files-in-go/

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func ProcessAM(calls *map[string]HamCall, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing AM")

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

		hc := HamCall{
			Callsign: record[4],
			Class:    record[5],
		}

		// fmt.Print("AM-", record[4], " ")
		updateMap(calls, &hc, record[4])
	}
}

func ProcessEN(calls *map[string]HamCall, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing EN")

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

		hc := HamCall{
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
		updateMap(calls, &hc, record[4])
	}

}

func ProcessHD(calls *map[string]HamCall, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing HD")

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

		hc := HamCall{
			Callsign:   record[4],
			Grant:      record[7],
			Expiration: record[8],
			FileNumber: record[2],
			Effective:  record[42],
		}

		// fmt.Print("HD-", record[4], " ")
		updateMap(calls, &hc, record[4])
	}

}

func ProcessLOTW(calls *map[string]HamCall, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing LOTW")

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

		hc := HamCall{
			Callsign: record[0],
			LOTW:     record[1] + record[2],
		}
		updateMap(calls, &hc, record[0])
	}
}

func updateMap(calls *map[string]HamCall, hc *HamCall, call string) {
	// this isn't thread safe right now

	// if it exists merge with current record
	_, c := (*calls)[call]
	if c {
		if err := mergo.Merge(hc, (*calls)[call]); err != nil {
			fmt.Printf("Error Merging: %v", err)
		}
	}

	(*calls)[hc.Callsign] = *hc
}

func WriteCall(data *HamCall, xht *xsync.Map, task chan Upload) error {
	// fmt.Print(data.Callsign, " ")
	call := strings.ToLower(strings.Replace(data.Callsign, "/", "-", -1))
	filename := "callsigns/" + call + ".json"
	file, _ := json.Marshal(data)
	task <- Upload{FileName: filename, FileContent: &file}
	return nil
}

func UploadWorker(task chan Upload, xht *xsync.Map, b2b *backblaze.Bucket, wg *sync.WaitGroup) {
	for file := range task {

		hasher := sha1.New()
		hasher.Write(*file.FileContent)
		hs := hasher.Sum(nil)
		hash := hex.EncodeToString(hs)

		h, exists := xht.Load(file.FileName)
		if !exists || h.(string) != hash {

			meta, err := b2b.UploadFile(file.FileName, nil, bytes.NewReader((*file.FileContent)))
			if err != nil {
				fmt.Println(err)
				continue
			}
			xht.Store(file.FileName, meta.ContentSha1)
			// fmt.Printf("uploaded file %s, %s = %s\n", meta.Name, meta.ContentSha1, hash)
		}
	}
	wg.Done()
}

func LoadHashTable(ht *HashTable, b2b *backblaze.Bucket) error {
	_, hashReader, err := b2b.DownloadFileByName("hash.gob")
	if err != nil {
		return err
	}
	fmt.Printf("downloaded hash file\n")

	hashDecoder := gob.NewDecoder(hashReader)
	err = hashDecoder.Decode(&ht)
	if err != nil {
		return err
	}

	return nil
}

func SaveHashTable(ht *HashTable, b2b *backblaze.Bucket) error {
	r, w := io.Pipe()
	go func() {
		encoder := gob.NewEncoder(w)
		encoder.Encode(ht)
		w.Close()
	}()

	file, err := b2b.UploadFile("hash.gob", nil, r)
	if err != nil {
		return err
	}

	fmt.Printf("uploaded hash file created at %d\n", file.UploadTimestamp)

	return nil

}
