package kingpinger

type response struct {
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

type pinger interface {
	ping() response
}
