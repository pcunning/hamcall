package data

type HamCall struct {
	Callsign   string    `json:"callsign"`
	Name       string    `json:"name"`
	FirstName  string    `json:"first_name"`
	Mi         string    `json:"mi,omitempty"`
	LastName   string    `json:"last_name"`
	Class      string    `json:"class"`
	Address    string    `json:"address,omitempty"`
	City       string    `json:"city,omitempty"`
	State      string    `json:"state,omitempty"`
	Zip        string    `json:"zip,omitempty"`
	PoBox      string    `json:"po_box,omitempty"`
	Grant      string    `json:"grant,omitempty"`
	Effective  string    `json:"effective,omitempty"`
	Expiration string    `json:"expiration,omitempty"`
	FRN        string    `json:"frn,omitempty"`
	FileNumber string    `json:"file_number,omitempty"`
	LOTW       string    `json:"last_lotw,omitempty"`
	LicenseKey string    `json:"license_key,omitempty"`
	DMRID      []int     `json:"dmr_id,omitempty"`
	Location   *Location `json:"location,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}
