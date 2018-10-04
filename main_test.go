package main

import (
	"net/http"
	"testing"
)

func Test_performPing(t *testing.T) {
	type args struct {
		a Pinger
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			performPing(tt.args.a)
		})
	}
}

func Test_handler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.w, tt.args.r)
		})
	}
}
