package data

type HamCall struct {
	Callsign   string    `json:"callsign"`
	Name       string    `json:"name"`
	FirstName  string    `json:"first_name"`
	Mi         string    `json:"mi"`
	LastName   string    `json:"last_name"`
	Class      string    `json:"class"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Zip        string    `json:"zip"`
	PoBox      string    `json:"po_box"`
	Grant      string    `json:"grant"`
	Effective  string    `json:"effective"`
	Expiration string    `json:"expiration"`
	FRN        string    `json:"frn"`
	FileNumber string    `json:"file_number"`
	LOTW       string    `json:"last_lotw"`
	LicenseKey string    `json:"license_key"`
	DMRID      []string  `json:"dmr_id,omitempty"`
	Location   *Location `json:"location,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}
