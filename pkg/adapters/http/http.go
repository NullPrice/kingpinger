package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	pinger "github.com/NullPrice/kingpinger/pkg/pinger"
)

// HTTP adapter for pinger
type HTTP struct {
	Result      pinger.Result
	PingRequest pinger.PingRequest
}

// ProcessResult - Handles dealing with the result
func (a *HTTP) ProcessResult() {
	// TODO: callback shit
	marshaledPayload, err := json.Marshal(a.Result)
	if err != nil {
		// We need proper error handling here we need to understand how the system all links together first
		log.Fatalln(err)
	}

	// service, ok := viper.Get("service").(string)
	// if ok != true {
	// 	log.Fatalln("service was not decoded correctly")
	// }

	resp, err := http.Post(a.PingRequest.CallbackURL, "application/json", bytes.NewBuffer(marshaledPayload))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
}

// SetResult - Sets the result value
func (a *HTTP) SetResult(result pinger.Result) {
	a.Result = result
}

// GetResult - Gets result value
func (a HTTP) GetResult() pinger.Result {
	return a.Result
}

// SetPingRequest - Sets ping request
func (a *HTTP) SetPingRequest(request pinger.PingRequest) {
	a.PingRequest = request
}

// GetPingRequest - Gets ping request
func (a HTTP) GetPingRequest() pinger.PingRequest {
	return a.PingRequest
}
