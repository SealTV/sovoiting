package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SealTV/sovoiting/server/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -package mocks -source server.go -destination mocks/server_mock.go

func TestServer_createVoting(t *testing.T) {
	type response struct {
		Status int    `json:"status"`
		Error  string `json:"error,omitempty"`
		Data   any    `json:"data,omitempty"`
	}

	tests := []struct {
		name     string
		data     []byte
		prepare  func(svc *mocks.MockServiser)
		want     response
		wantCode int
	}{
		{
			"1. invalid request body",
			[]byte(`invalid_json`),
			func(svc *mocks.MockServiser) {},
			response{
				Status: http.StatusBadRequest,
				Error:  "Invalid request data",
			},
			http.StatusBadRequest,
		},
		{
			"2. error on create vote",
			[]byte(`{"owner_address":"some_address", "names":["name_1", "name_2"]}`),
			func(svc *mocks.MockServiser) {
				svc.EXPECT().CreateVote(gomock.Any(), "some_address", []string{"name_1", "name_2"}).
					Return("", errors.New("some error"))
			},
			response{
				Status: http.StatusInternalServerError,
				Error:  "Internal server error",
			},
			http.StatusInternalServerError,
		},
		{
			"3. success crate vote",
			[]byte(`{"owner_address":"some_address", "names":["name_1", "name_2"]}`),
			func(svc *mocks.MockServiser) {
				svc.EXPECT().CreateVote(gomock.Any(), "some_address", []string{"name_1", "name_2"}).
					Return("some_vote_id", nil)
			},
			response{
				Status: http.StatusOK,
				Data:   "some_vote_id",
			},
			http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			svc := mocks.NewMockServiser(ctrl)
			tt.prepare(svc)

			recorder := httptest.NewRecorder()

			buffer := bytes.NewBuffer(tt.data)
			req := httptest.NewRequest(http.MethodPost, "/v1/voting/", buffer)

			New(svc).GetHttpHAndler().ServeHTTP(recorder, req)

			resp := recorder.Result()
			assert.Equal(t, tt.wantCode, resp.StatusCode)

			got := response{}
			err := json.NewDecoder(resp.Body).Decode(&got)
			assert.Nil(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
