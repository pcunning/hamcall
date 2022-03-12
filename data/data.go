package data

import (
	"fmt"

	"github.com/imdario/mergo"
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
	DMRID      []int  `json:"dmr_id,omitempty"`
}

func UpdateMap(calls *map[string]HamCall, hc *HamCall, call string) {
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
