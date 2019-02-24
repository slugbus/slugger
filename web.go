package slugger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var tapsAPIURL = "http://bts.ucsc.edu:8081/location/get"

// OverrideURL overides the default TAPS API URL (http://bts.ucsc.edu:8081/location/get)
// to the given url.
// This should be used for development purposes only.
func OverrideURL(url string) {
	tapsAPIURL = url
}

// Query calls the TAPS API URL (default: http://bts.ucsc.edu:8081/location/get), and
// returns a slice of Buses if successful.
func Query() ([]Bus, error) {
	// Query the tapsAPIURL.
	resp, err := http.Get(tapsAPIURL)
	if err != nil {
		return nil, errors.Wrapf(err, "could not query %s", tapsAPIURL)
	}
	// Remember to close the body.
	defer resp.Body.Close()
	// Check for successful http code.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got invalid status code (%d) from %s", resp.StatusCode, tapsAPIURL)
	}
	// Read the body of the response.
	tbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Create a slice of Buses
	tbuses := []Bus{}
	// Attempt to Unmarshal the bytes
	if err := json.Unmarshal(tbytes, &tbuses); err != nil {
		return nil, err
	}
	// If successful, return the slice.
	return tbuses, nil
}
