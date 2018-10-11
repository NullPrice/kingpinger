package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	pinger "github.com/NullPrice/kingpinger/pkg/pinger"
	ping "github.com/sparrc/go-ping"
)

// HTTP adapter for pinger
type HTTP struct {
	Result      pinger.Result
	PingRequest pinger.PingRequest
	Ping        *ping.Pinger
}

// ProcessResult - Handles dealing with the result
func (httpAdapter *HTTP) ProcessResult() {

	marshaledPayload, err := json.Marshal(httpAdapter.Result)
	if err != nil {
		// We need proper error handling here we need to understand how the system all links together first
		log.Fatalln(err)
	}

	// service, ok := viper.Get("service").(string)
	// if ok != true {
	// 	log.Fatalln("service was not decoded correctly")
	// }

	resp, err := http.Post(httpAdapter.PingRequest.CallbackURL, "application/json", bytes.NewBuffer(marshaledPayload))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
}

// SetResult - Sets the result value
func (httpAdapter *HTTP) SetResult(result pinger.Result) {
	httpAdapter.Result = result
}

// GetResult - Gets result value
func (httpAdapter *HTTP) GetResult() pinger.Result {
	return httpAdapter.Result
}

// SetPingRequest - Sets ping request
func (httpAdapter *HTTP) SetPingRequest(request pinger.PingRequest) {
	httpAdapter.PingRequest = request
}

// GetPingRequest - Gets ping request
func (httpAdapter *HTTP) GetPingRequest() pinger.PingRequest {
	return httpAdapter.PingRequest
}

// Run - Runs a ping process and sets updates the result struct
func (httpAdapter *HTTP) Run() {
	httpAdapter.Ping.Run()
	httpAdapter.Result = pinger.Result{JobID: httpAdapter.PingRequest.JobID, Statistics: httpAdapter.Ping.Statistics()}
}

// SetPingDependency - Sets the ping dependency
func (httpAdapter *HTTP) SetPingDependency(x *ping.Pinger, err error) {
	if err != nil {
		log.Fatalln(err)
	}
	x.SetPrivileged(true)
	x.Count = httpAdapter.PingRequest.Count
	httpAdapter.Ping = x
}

// GetPingDependency - Gets the ping dependency struct
func (httpAdapter *HTTP) GetPingDependency() *ping.Pinger {
	return httpAdapter.Ping
}
