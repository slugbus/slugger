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
	"fmt"
	"testing"
)

func mergeTest(t *testing.T) {
	t.Run()
	testBus1 := Bus{"Bus One", 60.0, 60.0, "type"}
	updatedTestBus := UpdatedBus{testBus1, 30.0, 30.0}

	busOneMap := BusMap{}
	busOneMap[testBus1.ID] = testBus1
	updatedBusOneMap := UpdatedBusMap{}
	updatedBusOneMap[updatedTestBus.ID] = updatedTestBus

	mergedUpdatedBusMap := mergeWithState(busOneMap, 30.0, updatedBusOneMap)

	fmt.Printf("ID: %s\n, Speed: %f\n", mergedUpdatedBusMap[updatedTestBus.ID].ID, mergedUpdatedBusMap[updatedTestBus.ID].Speed)

}

func TestQuery(t *testing.T) {
	// The default query should be okay.
	t.Run("OK Query", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}
		if got == nil {
			t.Errorf("got no response")
		}
	})

	// These urls should fail.
	tURLs := []string{"https://www.random.org/bad_url", "bad_url", "https://jsonplaceholder.typicode.com/todos"}
	for _, url := range tURLs {
		t.Run(fmt.Sprintf("Bad Query: %s", url), func(t *testing.T) {
			s := NewSource(url)
			got, err := s.Query()
			if err == nil {
				t.Errorf("got no error, wanted error")
			}
			if got != nil {
				t.Errorf("got a response")
			}
		})
	}

	// Reseting the URL should make the
	// first test pass, despite the changes above.
	t.Run("OK Query", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}
		if got == nil {
			t.Errorf("got no response")
		}
	})

	// If we get a successful query, test to make
	// sure none of the results are their 0 values.
	t.Run("Default Check", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}

		for _, bus := range got {

			if bus.ID == "" {
				t.Errorf("got default value for ID, got %q , did not want %q", bus.ID, bus.ID)
			}

			if bus.Type == "" {
				t.Errorf("got default value for Type, got %q , did not want %q", bus.Type, bus.Type)
			}

			if bus.Lat == 0 {
				t.Errorf("got default value for Lat, got %v , did not want %v", bus.Lat, bus.Lat)
			}

			if bus.Lon == 0 {
				t.Errorf("got default value for Lon, got %v , did not want %v", bus.Lon, bus.Lon)
			}

		}

	})
}

func TestQueryAsMap(t *testing.T) {

	t.Run("OK Query", func(t *testing.T) {
		tbus, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted nil", err)
		}
		mbus, err := QueryAsMap()
		if err != nil {
			t.Errorf("got err: %v, wanted nil", err)
		}

		if len(tbus) != len(mbus) {
			t.Errorf("lengths do no math %v != %v", len(tbus), len(mbus))
		}

		for _, bus := range tbus {
			if _, ok := mbus[bus.ID]; !ok {
				t.Errorf("bus in query is not in map, %v", bus)
			}
		}

	})

	t.Run("Possibly OK Query", func(t *testing.T) {
		s := NewSource("http://0.0.0.0:9191/location/get")
		tbus, err := s.Query()
		if err != nil {
			// This isn't an error it just means the mock server is not
			// running
			return
		}
		mbus, err := s.QueryAsMap()
		if err != nil {
			// Same reason as above.
			return
		}

		if len(tbus) != len(mbus) {
			t.Errorf("lengths do no math %v != %v", len(tbus), len(mbus))
		}

		for _, bus := range tbus {
			if _, ok := mbus[bus.ID]; !ok {
				t.Errorf("bus in query is not in map, %v", bus)
			}
		}

	})

	t.Run("Bad Query", func(t *testing.T) {
		s := NewSource("BAD")
		_, err := s.QueryAsMap()
		if err == nil {
			t.Errorf("got nil, wanted error")
		}

	})

}
