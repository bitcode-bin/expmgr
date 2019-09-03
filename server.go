package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	wallet *Wallet
	router chi.Router
	log    Logger
}

func (s *server) Init() {
	s.routes()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.router = chi.NewRouter()
	s.router.Route("/api", func(r chi.Router) {
		r.Get("/balance", s.handleBalanceGet())
	})
}

func (s *server) handleBalanceGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log.Infof("%s %s", r.Method, r.URL.String())

		bal := s.wallet.Balance()
		data := map[string]interface{}{
			"balance": bal,
		}

		s.respond(w, r, http.StatusOK, data)
	}
}
