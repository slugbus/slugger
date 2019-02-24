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

// Package slugger is a go client library
// for the UCSC TAPS API
package slugger

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	// The deefault query should be okay.
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
			OverrideURL(url)
			got, err := Query()
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
	RestoreURL()
	t.Run("OK Query", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}
		if got == nil {
			t.Errorf("got no response")
		}
	})
}
