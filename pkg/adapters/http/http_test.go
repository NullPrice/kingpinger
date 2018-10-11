package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sparrc/go-ping"

	"github.com/NullPrice/kingpinger/pkg/pinger"
)

func TestProcessResult(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Testing to see if callback is properly formatted.
		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected 'application/json' request, got '%s'", r.Header.Get("Content-Type"))
		}

		fmt.Println(r.Header)
	}))
	defer ts.Close()

	httpAdapter := HTTP{
		Result:      pinger.Result{"12345-67890", &ping.Statistics{}},
		PingRequest: pinger.PingRequest{CallbackURL: ts.URL + "/callback"},
	}
	httpAdapter.ProcessResult()

}
