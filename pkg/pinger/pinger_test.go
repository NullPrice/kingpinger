package pinger

import "testing"

func TestProcess(t *testing.T) {
	type args struct {
		a Adapter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Process(tt.args.a)
		})
	}
}
