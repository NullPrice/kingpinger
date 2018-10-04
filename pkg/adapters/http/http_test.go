package http

import (
	"reflect"
	"testing"

	pinger "github.com/NullPrice/kingpinger/pkg/pinger"
)

func TestHTTP_ProcessResult(t *testing.T) {
	tests := []struct {
		name string
		a    *HTTP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.ProcessResult()
		})
	}
}

func TestHTTP_SetResult(t *testing.T) {
	type args struct {
		result pinger.Result
	}
	tests := []struct {
		name string
		a    *HTTP
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.SetResult(tt.args.result)
		})
	}
}

func TestHTTP_GetResult(t *testing.T) {
	tests := []struct {
		name string
		a    HTTP
		want pinger.Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetResult(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTP.GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTP_SetPingRequest(t *testing.T) {
	type args struct {
		request pinger.PingRequest
	}
	tests := []struct {
		name string
		a    *HTTP
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.SetPingRequest(tt.args.request)
		})
	}
}

func TestHTTP_GetPingRequest(t *testing.T) {
	tests := []struct {
		name string
		a    HTTP
		want pinger.PingRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetPingRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTP.GetPingRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
