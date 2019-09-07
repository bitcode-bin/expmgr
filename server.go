package main

import (
	"net/http"

	"github.com/bitcode-bin/expmgr/logger"
	"github.com/go-chi/chi"
)

type server struct {
	wallet *Wallet
	router chi.Router
	log    logger.Logger
}

func (s *server) Init() {
	s.routes()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.router = chi.NewRouter()
	s.router.Use(
		requestID,
		requestLogger(s.log),
	)

	s.router.Route("/api", func(r chi.Router) {
		r.Get("/balance", s.handleBalanceGet())
		r.Post("/income", s.handleIncomePost())
		r.Post("/expense", s.handleExpensePost())
	})
}
