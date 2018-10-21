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
func (adapter *HTTP) ProcessResult() {

	marshaledPayload, err := json.Marshal(adapter.Result)
	if err != nil {
		// We need proper error handling here we need to understand how the system all links together first
		log.Fatalln(err)
	}

	// service, ok := viper.Get("service").(string)
	// if ok != true {
	// 	log.Fatalln("service was not decoded correctly")
	// }

	resp, err := http.Post(adapter.PingRequest.CallbackURL, "application/json", bytes.NewBuffer(marshaledPayload))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
}

// SetResult - Sets the result value
func (adapter *HTTP) SetResult(result pinger.Result) {
	adapter.Result = result
}

// GetResult - Gets result value
func (adapter *HTTP) GetResult() pinger.Result {
	return adapter.Result
}

// SetPingRequest - Sets ping request
func (adapter *HTTP) SetPingRequest(request pinger.PingRequest) {
	adapter.PingRequest = request
}

// GetPingRequest - Gets ping request
func (adapter *HTTP) GetPingRequest() pinger.PingRequest {
	return adapter.PingRequest
}

// Run - Runs a ping process and sets updates the result struct
func (adapter *HTTP) Run() {
	if adapter.Ping == nil {
		// If pinger has not been set manually
		adapter.SetPingDependency(ping.NewPinger(adapter.GetPingRequest().Target))
	}
	adapter.Ping.Run()
	adapter.Result = pinger.Result{JobID: adapter.PingRequest.JobID, Statistics: adapter.Ping.Statistics()}
}

// SetPingDependency - Sets the ping dependency
func (adapter *HTTP) SetPingDependency(x *ping.Pinger, err error) {
	if err != nil {
		// TODO: We should handle this: this does an os.exit behind the scenes, we want to handle all errors as this is a client
		log.Fatalln(err)
	}
	x.SetPrivileged(true)
	x.Count = adapter.PingRequest.Count
	adapter.Ping = x
}

// GetPingDependency - Gets the ping dependency struct
func (adapter *HTTP) GetPingDependency() *ping.Pinger {
	return adapter.Ping
}
