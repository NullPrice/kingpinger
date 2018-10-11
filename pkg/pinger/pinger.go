package pinger

import (
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
	SetPingDependency(x *ping.Pinger, err error)
	GetPingDependency() *ping.Pinger
	Run()
}

// Process a job
func Process(adapter Adapter) {
	// Set the ping dependency -- Could be done elsewhere
	// adapter.SetPingDependency(ping.NewPinger(adapter.GetPingRequest().Target))
	adapter.Run()
	adapter.ProcessResult()
}
