package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pcunning/hamcall/b2"
	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/source/lotw"
	"github.com/pcunning/hamcall/source/radioid"
	"github.com/pcunning/hamcall/source/uls"
)

func main() {
	start := time.Now()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	osSigExit := make(chan bool, 1)

	go func() {
		<-sigs
		osSigExit <- true
	}()

	keyID := os.Getenv("B2_KEYID")
	applicationKey := os.Getenv("B2_APPKEY")
	uploadWorkers := 200

	downloadFiles()

	calls := make(map[string]data.HamCall)
	process(&calls)

	writeToB2(&calls, keyID, applicationKey, uploadWorkers, osSigExit)

	fmt.Printf("total runtime %s\n", time.Since(start).String())
}

func downloadFiles() {
	var wg sync.WaitGroup

	wg.Add(3)

	go uls.Download(&wg)
	go radioid.Download(&wg)
	go lotw.Download(&wg)

	wg.Wait()
}

func process(calls *map[string]data.HamCall) {
	uls.Process(calls)
	radioid.Process(calls)
	lotw.Process(calls)
}

func writeToB2(calls *map[string]data.HamCall, keyID, applicationKey string, uploadWorkers int, osSigExit chan bool) {

	b, err := b2.New(keyID, applicationKey, uploadWorkers)
	if err != nil {
		fmt.Println("error creating b2 client: ", err)
		return
	}

	err = b.Write(calls, osSigExit)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
