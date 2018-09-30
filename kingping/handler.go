package kingping

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	ping "github.com/sparrc/go-ping"
)

type job struct {
	JobID string `json:"job_id"`
	*ping.Statistics
}

// JobRequest struct
// TODO: Should be camelcase but we were derps when doing this the first time
type JobRequest struct {
	// https://stackoverflow.com/questions/9452897/how-to-decode-json-with-type-convert-from-string-to-float64-in-golang
	Target      string `json:"target"`
	Count       int    `json:"count"`
	JobID       string `json:"job_id"`
	CallbackURL string `json:"callback_url"`
}

type pinger interface {
	ping() job
}

// Process a job
func (jr *JobRequest) Process() {
	// jr.Count = 100
	pinger, err := ping.NewPinger(jr.Target)
	if err != nil {
		// We need proper error handling here we need to understand how the system all links together first
		log.Fatalln(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = jr.Count
	pinger.Run()
	s := pinger.Statistics()
	job := job{JobID: jr.JobID, Statistics: s}
	sendCallback(job, jr.CallbackURL)
}

func sendCallback(j job, callbackURL string) {
	marshaledPayload, err := json.Marshal(j)
	if err != nil {
		// We need proper error handling here we need to understand how the system all links together first
		log.Fatalln(err)
	}

	// service, ok := viper.Get("service").(string)
	// if ok != true {
	// 	log.Fatalln("service was not decoded correctly")
	// }

	resp, err := http.Post(callbackURL, "application/json", bytes.NewBuffer(marshaledPayload))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
}
