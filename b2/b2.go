package b2

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/pcunning/hamcall/data"
	"github.com/puzpuzpuz/xsync"
	"gopkg.in/kothar/go-backblaze.v0"
)

type b2 struct {
	b2b     *backblaze.Bucket
	xht     *xsync.Map
	workers int
}

type hashTable map[string]string

type upload struct {
	FileName    string
	FileContent *[]byte
}

func New(keyID, applicationKey string, workers int) (*b2, error) {
	b, err := backblaze.NewB2(backblaze.Credentials{
		KeyID:          keyID,
		ApplicationKey: applicationKey,
	})
	if err != nil {
		return nil, err
	}

	b2b, err := b.Bucket("hamcall")
	if err != nil {
		return nil, err
	}

	xht := xsync.NewMap()

	return &b2{b2b, xht, workers}, nil
}

func (b b2) Write(calls *map[string]data.HamCall, osSigExit chan bool) error {

	start := time.Now()

	err := loadHashTable(b.xht, b.b2b)
	if err != nil {
		return err
	}

	go func() {
		// Create workers
		fmt.Printf("%s: creating workers\n", time.Since(start))
		uploadTasks := make(chan upload, b.workers*2)
		group := sync.WaitGroup{}
		for i := 0; i < b.workers; i++ {
			group.Add(1)
			go uploadWorker(uploadTasks, b.xht, b.b2b, &group)
		}

		fmt.Printf("%s: Writing %d files to disk\n", time.Since(start), len(*calls))
		i := 1
		batchTime := time.Now()
		for _, v := range *calls {
			writeCall(&v, b.xht, uploadTasks)

			if i%10000 == 0 {
				rps := 10000 / time.Since(batchTime).Seconds()
				etr := time.Duration((float64(len(*calls)-i) / rps)) * time.Second
				fmt.Printf("%s: %d... %s = %.2f/second etr: %s\n", time.Since(start), i, time.Since(batchTime), rps, etr)
				batchTime = time.Now()
			}
			i++
		}

		close(uploadTasks)
		group.Wait()

		fmt.Printf("%s: ", time.Since(start))
		osSigExit <- true
	}()

	<-osSigExit

	fmt.Printf("\n\nexiting... please wait while saving hash table\n\n")
	saveHashTable(b.xht, b.b2b)
	return nil
}

func writeCall(data *data.HamCall, xht *xsync.Map, task chan upload) {
	// fmt.Print(data.Callsign, " ")
	call := strings.ToLower(strings.Replace(data.Callsign, "/", "-", -1))
	filename := "callsigns/" + call + ".json"
	file, _ := json.Marshal(data)
	task <- upload{FileName: filename, FileContent: &file}
}

func uploadWorker(task chan upload, xht *xsync.Map, b2b *backblaze.Bucket, wg *sync.WaitGroup) {
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

func loadHashTable(xht *xsync.Map, b2b *backblaze.Bucket) error {
	_, hashReader, err := b2b.DownloadFileByName("hash.gob")
	if err != nil {
		return err
	}
	fmt.Printf("downloaded hash file\n")

	var ht hashTable
	hashDecoder := gob.NewDecoder(hashReader)
	err = hashDecoder.Decode(&ht)
	if err != nil {
		return err
	}

	fmt.Printf("%d entries in hash table\n", len(ht))

	for k, v := range ht {
		xht.Store(k, v)
	}

	return nil
}

func saveHashTable(xht *xsync.Map, b2b *backblaze.Bucket) error {
	var ht hashTable
	ht = make(hashTable)
	xht.Range(func(k string, v interface{}) bool {
		ht[k] = v.(string)
		return true
	})

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
