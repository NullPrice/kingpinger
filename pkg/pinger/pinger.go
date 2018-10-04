package pinger

import (
	"log"

	ping "github.com/sparrc/go-ping"
)

// Result struct
type Result struct {
	JobID string `json:"job_id"`
	*ping.Statistics
}

// PingRequest struct
// TODO: Should be camelcase but we were derps when doing this the first time
type PingRequest struct {
	// https://stackoverflow.com/questions/9452897/how-to-decode-json-with-type-convert-from-string-to-float64-in-golang
	Target      string `json:"target"`
	Count       int    `json:"count"`
	JobID       string `json:"job_id"`
	CallbackURL string `json:"callback_url"`
}

// Adapter interface
type Adapter interface {
	ProcessResult()
	SetResult(x Result)
	GetResult() Result
	SetPingRequest(x PingRequest)
	GetPingRequest() PingRequest
}

// Process a job
func Process(a Adapter) {
	pingRequest := a.GetPingRequest()
	pinger, err := ping.NewPinger(pingRequest.Target)
	if err != nil {
		// We need proper error handling here we need to understand how the system all links together first
		log.Fatalln(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = pingRequest.Count
	pinger.Run()

	stats := pinger.Statistics()
	a.SetResult(Result{JobID: pingRequest.JobID, Statistics: stats})
	a.ProcessResult()
}
