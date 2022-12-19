package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
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
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	voteID, err := s.svc.CreateVote(r.Context(), req.OwnerAddress, req.Names)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result := map[string]any{
		"status": http.StatusOK,
		"data":   voteID,
	}
	render.JSON(w, r, result)
}
