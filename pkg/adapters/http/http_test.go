package http

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/sparrc/go-ping"

	"github.com/NullPrice/kingpinger/pkg/pinger"
)

func TestProcessResult(t *testing.T) {
	jobID := "b1d29348-ca8e-46f2-8630-88450636791e"
	ipaddr, _ := net.ResolveIPAddr("ip", "216.58.198.238")
	pingStatisticsStruct := &ping.Statistics{
		PacketsRecv: 5,
		PacketsSent: 5,
		PacketLoss:  0,
		IPAddr:      ipaddr,
		Addr:        "google.com",
		Rtts: []time.Duration{18004400,
			18003400,
			18003500,
			18003400,
			18004300},
		MinRtt:    18003400,
		MaxRtt:    18004400,
		AvgRtt:    18003800,
		StdDevRtt: 451,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Testing to see if callback is properly formatted.
		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected 'application/json' request, got '%s'", r.Header.Get("Content-Type"))
		}

		var testResponse pinger.Result
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&testResponse)
		if err != nil {
			t.Error(err)
		}

		if testResponse.JobID != jobID {
			t.Errorf("Expected '%s' as JobID, got '%s'", jobID, testResponse.JobID)
		}

		if !reflect.DeepEqual(pingStatisticsStruct, testResponse.Statistics) {
			t.Errorf("Decoded struct: %+v\n Does not equal expected struct: %+v\n", testResponse.Statistics, pingStatisticsStruct)
		}
	}))
	defer ts.Close()

	httpAdapter := HTTP{
		Result:      pinger.Result{JobID: jobID, Statistics: pingStatisticsStruct},
		PingRequest: pinger.PingRequest{CallbackURL: ts.URL + "/callback"},
	}
	httpAdapter.ProcessResult()

}

func TestSetResult(t *testing.T) {

	httpAdapter := HTTP{}
	result := pinger.Result{JobID: "b1d29348-ca8e-46f2-8630-88450636791e"}
	httpAdapter.SetResult(result)
	if httpAdapter.Result == (pinger.Result{}) {
		t.Errorf("Result struct in httpAdapter is empty: %+v\n", httpAdapter)
	}

	// TODO: Should add an empty check function for results
	if httpAdapter.Result.JobID != result.JobID {
		t.Errorf("Expected '%s' as JobID, got '%s'", result.JobID, httpAdapter.Result.JobID)
	}

}

func TestGetResult(t *testing.T) {

	httpAdapter := HTTP{}
	result := pinger.Result{JobID: "b1d29348-ca8e-46f2-8630-88450636791e"}
	httpAdapter.SetResult(result)
	getResult := httpAdapter.GetResult()
	if getResult.JobID != result.JobID {
		t.Errorf("Expected '%s' as JobID, got '%s'", result.JobID, getResult.JobID)
	}
}

// TestSetPingDependency - For mocks
func TestSetPingDependency(t *testing.T) {

	httpAdapter := HTTP{PingRequest: pinger.PingRequest{
		Target:      "google.com",
		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
		Count:       5,
		CallbackURL: "http://example.com/callback",
	},
	}
	httpAdapter.SetPingDependency(ping.NewPinger(httpAdapter.GetPingRequest().Target))

	if httpAdapter.Ping == nil {
		t.Errorf("Ping dependency has failed to be set")
	}
}

//TestGetPingDependency - For mocks
func TestGetPingDependency(t *testing.T) {

	httpAdapter := HTTP{PingRequest: pinger.PingRequest{
		Target:      "google.com",
		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
		Count:       5,
		CallbackURL: "http://example.com/callback",
	},
	}
	httpAdapter.SetPingDependency(ping.NewPinger(httpAdapter.GetPingRequest().Target))

	if httpAdapter.GetPingDependency() == nil {
		t.Errorf("Failed to get ping dependency using get function")
	}
}

func TestSetPingRequest(t *testing.T) {

	httpAdapter := HTTP{}
	pingRequest := pinger.PingRequest{
		Target:      "google.com",
		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
		Count:       5,
		CallbackURL: "http://example.com/callback",
	}
	httpAdapter.SetPingRequest(pingRequest)

	if httpAdapter.PingRequest == (pinger.PingRequest{}) {
		t.Errorf("Ping request has failed to be set")
	}
}

func TestGetPingRequest(t *testing.T) {

	httpAdapter := HTTP{}
	pingRequest := pinger.PingRequest{
		Target:      "google.com",
		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
		Count:       5,
		CallbackURL: "http://example.com/callback",
	}
	httpAdapter.SetPingRequest(pingRequest)
	getPingRequest := httpAdapter.GetPingRequest()

	if !reflect.DeepEqual(pingRequest, getPingRequest) {
		t.Errorf("Struct: %+v\n Does not equal expected struct: %+v\n", getPingRequest, pingRequest)
	}
}

func TestRunPing(t *testing.T) {
	httpAdapter := HTTP{}
	pingRequest := pinger.PingRequest{
		Target:      "google.com",
		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
		Count:       5,
		CallbackURL: "http://example.com/callback",
	}
	httpAdapter.SetPingRequest(pingRequest)
	httpAdapter.Run()
	if httpAdapter.Result == (pinger.Result{}) {
		t.Errorf("Results have not been recorded")
	}
}

func TestRunPingWithExternalPingDependency(t *testing.T) {
	httpAdapter := HTTP{}
	pingRequest := pinger.PingRequest{
		Target:      "google.com",
		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
		Count:       5,
		CallbackURL: "http://example.com/callback",
	}
	httpAdapter.SetPingRequest(pingRequest)
	httpAdapter.SetPingDependency(ping.NewPinger(httpAdapter.GetPingRequest().Target))
	httpAdapter.Run()
	if httpAdapter.Result == (pinger.Result{}) {
		t.Errorf("Results have not been recorded")
	}
}

// TODO: After further improvements to error handling
// func TestSetPingDependencyWithError(t *testing.T) {

// 	httpAdapter := HTTP{PingRequest: pinger.PingRequest{
// 		Target:      "google.com",
// 		JobID:       "b1d29348-ca8e-46f2-8630-88450636791e",
// 		Count:       5,
// 		CallbackURL: "http://example.com/callback",
// 	},
// 	}
// 	httpAdapter.SetPingDependency(&ping.Pinger{}, errors.New("Test"))
// }
