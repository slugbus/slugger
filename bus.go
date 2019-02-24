package slugger

// Bus is a structure that
// contains the json response
// from the UCSC TAPS server.
type Bus struct {
	ID   string  `json:"id"`
	Lon  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
	Type string  `json:"type"`
}
