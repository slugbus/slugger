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
// Updated Bus is a structure that
// contains the updated json response
// from the UCSC Taps server
type UpdatedBus struct {
	Bus 	// contains the original bus struct 
	Speed float64 `json:"speed"`
	Angle float64 `json:"angle"`
}

// BusMap is a collection of buses
// where the key is an bus.ID and the
// value is the corresponding Bus struct.
type BusMap map[string]Bus

// Updated BusMap is a collection of updated buses
// where the key is an updatedbus.bus.id and the 
// value is the corresponding UpdatedBus struct
type UpatedBusMap map[string]UpdatedBus

// Source is a string that defines
// what URL should be called
type Source string

const defaultSource = Source("http://bts.ucsc.edu:8081/location/get")

// Source functions:

// NewSource returns a new source struct
// with url: url.
func NewSource(url string) Source {
	return Source(url)
}

// Query calls the Source URL, and
// returns a slice of Buses if successful.
func (s Source) Query() ([]Bus, error) {
	url := string(s)
	// Query the tapsAPIURL.
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "could not query %s", url)
	}
	// Remember to close the body.
	defer resp.Body.Close()
	// Check for successful http code.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got invalid status code (%d) from %s", resp.StatusCode, url)
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

// QueryAsMap calls the source API similar to Query and returns
// a BusMap if successful.
func (s Source) QueryAsMap() (BusMap, error) {
	tbus, err := s.Query()
	if err != nil {
		return nil, err
	}
	mbus := MapFromQuery(tbus)
	return mbus, nil
}

// Package Functions

// MapFromQuery transforms a slice of Bus to a BusMap.
func MapFromQuery(tbus []Bus) BusMap {
	mbus := BusMap{}
	for _, bus := range tbus {
		mbus[bus.ID] = bus
	}
	return mbus
}

// Query calls the official TAPS API URL, and
// returns a slice of Buses if successful.
func Query() ([]Bus, error) {
	return defaultSource.Query()
}

// QueryAsMap calls the taps API similar to Query and returns
// a BusMap if successful.
func QueryAsMap() (BusMap, error) {
	return defaultSource.QueryAsMap()
}
