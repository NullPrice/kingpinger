package kingping

type job struct {
	jobID      string
	host       string
	sent       string
	receieved  string
	packetLoss string
	min        string
	avg        string
	max        string
	jitter     string
}

// JobRequest struct
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

// func (Job) ping(c) {

// }
