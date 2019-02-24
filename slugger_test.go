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
