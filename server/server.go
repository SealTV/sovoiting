package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Serviser interface {
	CreateVote(ctx context.Context, ownerAddres string, names []string) (string, error)
	AddVoter(ctx context.Context, voteID, voterAddres string) error
	DelegateVote(ctx context.Context, voteID, voter, to string) error
	Vote(ctx context.Context, voteID, viterAddres string, proposalID uint) error

	GetProposalse(ctx context.Context, voteID string) ([]string, error)
	GetWinnerName(ctx context.Context, voteID string) (string, error)
}

type Server struct {
	svc Serviser
}

func New(svc Serviser) *Server {
	return &Server{
		svc: svc,
	}
}

func (s *Server) GetHttpHAndler() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/hw", s.helloWorld)

		r.Route("/voting", func(r chi.Router) {
			r.Post("/", s.createVoting)
		})
	})

	return r
}

func (s *Server) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func (s *Server) createVoting(w http.ResponseWriter, r *http.Request) {
	req := struct {
		OwnerAddress string   `json:"owner_address"`
		Names        []string `json:"names"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		JSONError(w, http.StatusBadRequest, errors.New("Invalid request data"))
		return
	}
	defer r.Body.Close()

	voteID, err := s.svc.CreateVote(r.Context(), req.OwnerAddress, req.Names)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, errors.New("Internal server error"))
		return
	}

	JSON(w, voteID)
}

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
func JSON(w http.ResponseWriter, data any) {
	respData := map[string]any{
		"status": http.StatusOK,
		"data":   data,
	}
	jsonResp(w, http.StatusOK, respData)
}

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
func JSONError(w http.ResponseWriter, status int, err error) {
	data := map[string]any{
		"status": status,
		"error":  err.Error(),
	}
	jsonResp(w, status, data)
}

func jsonResp(w http.ResponseWriter, status int, data any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	w.Write(buf.Bytes()) //nolint:errcheck
}
