package server

import (
	"net/http"
	"testing"
)

//go:generate mockgen -package mocks -source server.go -destination mocks/server_mock.go

func TestServer_createVoting(t *testing.T) {
	type fields struct {
		svc Serviser
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				svc: tt.fields.svc,
			}
			s.createVoting(tt.args.w, tt.args.r)
		})
	}
}
