//   Copyright 2019 The SlugBus++ Authors.

//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at

//        http://www.apache.org/licenses/LICENSE-2.0

//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package taps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Bus is a structure that
// contains the json response
// from the UCSC TAPS server.
type Bus struct {
	ID   string  `json:"id"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Type string  `json:"type"`
}

var tapsAPIURL = "http://bts.ucsc.edu:8081/location/get"

// OverrideURL overides the default TAPS API URL (http://bts.ucsc.edu:8081/location/get)
// to the given url.
// This should be used for development purposes only.
func OverrideURL(url string) {
	tapsAPIURL = url
}

// RestoreURL sets the TAPS URL back to the default of http://bts.ucsc.edu:8081/location/get.
func RestoreURL() {
	tapsAPIURL = "http://bts.ucsc.edu:8081/location/get"
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
