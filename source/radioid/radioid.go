package radioid

import (
	"fmt"
	"sync"

	"github.com/pcunning/hamcall/data"
	"github.com/pcunning/hamcall/downloader"
)

func Download(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Downloading radioid data")
	err := downloader.FetchHttp("tmp/dmrid.dat", "https://www.radioid.net/static/dmrid.dat")
	if err != nil {
		return err
	}
	return nil
}

func Process(calls *map[string]data.HamCall) {

}
