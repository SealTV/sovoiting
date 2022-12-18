package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Serviser interface {
	CreateVote(ctx context.Context, voteID, ownerAddres string, names []string) (string, error)
	AddVoter(ctx context.Context, voteID, voterAddres string) error
	DelegateVote(ctx context.Context, voteID, voter, to string) error
	Vote(ctx context.Context, voteID, viterAddres string, proposalID uint) error

	GetProposalse(ctx context.Context, voteID string) ([]string, error)
	GetWinnerName(ctx context.Context, voteID string) (string, error)
}

type Server struct{}

func New() *Server {
	return &Server{}
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
	})

	return r
}

func (s *Server) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
