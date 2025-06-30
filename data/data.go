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

// CallRedirections tracks mappings for former callsigns to current callsigns
type CallRedirections struct {
	// FRNToCurrentCall maps FRN to the most recent active callsign for that person
	FRNToCurrentCall map[string]string
	// FormerCallToFRN maps former callsigns to their FRN for redirection lookup
	FormerCallToFRN map[string]string
}

// NewCallRedirections creates a new CallRedirections instance
func NewCallRedirections() *CallRedirections {
	return &CallRedirections{
		FRNToCurrentCall: make(map[string]string),
		FormerCallToFRN:  make(map[string]string),
	}
}

// ResolveCallsign returns the current active callsign for a given callsign
// If the callsign is current/active, returns it unchanged
// If the callsign is former, returns the current callsign for that FRN
func (cr *CallRedirections) ResolveCallsign(callsign string) string {
	// Check if this is a former callsign that needs redirection
	if frn, isFormer := cr.FormerCallToFRN[callsign]; isFormer {
		if currentCall, hasCurrentCall := cr.FRNToCurrentCall[frn]; hasCurrentCall {
			return currentCall
		}
	}
	// Return the original callsign (either it's current or we don't have redirection info)
	return callsign
}
