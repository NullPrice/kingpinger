package kingping

import (
	"testing"
)

func TestJobRequest_Process(t *testing.T) {
	tests := []struct {
		name string
		jr   *JobRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.jr.Process()
		})
	}
}

func Test_sendCallback(t *testing.T) {
	type args struct {
		j           job
		callbackURL string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendCallback(tt.args.j, tt.args.callbackURL)
		})
	}
}
