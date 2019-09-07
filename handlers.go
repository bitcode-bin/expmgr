package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleBalanceGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log.WithFields(map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.String(),
		}).Info("")

		bal := s.wallet.Balance()
		data := map[string]interface{}{
			"balance": bal,
		}

		s.respond(w, r, http.StatusOK, data)
	}
}

func (s *server) handleIncomePost() http.HandlerFunc {
	type request struct {
		Income int `json:"income"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log := s.log.WithFields(map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.String(),
		})

		log.Info("")

		var req request
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			log.WithFields(map[string]interface{}{
				"error": err,
			}).Error("failed to decode request")

			data := map[string]interface{}{"error": "invalid json"}
			s.respond(w, r, http.StatusBadRequest, data)
			return
		}

		s.wallet.AddIncome(req.Income)
		bal := s.wallet.Balance()

		data := map[string]interface{}{
			"balance": bal,
		}

		s.respond(w, r, http.StatusOK, data)
	}
}

func (s *server) handleExpensePost() http.HandlerFunc {
	type request struct {
		Expense int `json:"expense"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log := s.log.WithFields(map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.String(),
		})

		log.Info("")

		var req request
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			log.WithFields(map[string]interface{}{
				"error": err,
			}).Error("failed to decode request")

			data := map[string]interface{}{"error": "invalid json"}
			s.respond(w, r, http.StatusBadRequest, data)
			return
		}

		s.wallet.AddExpense(req.Expense)
		bal := s.wallet.Balance()

		data := map[string]interface{}{
			"balance": bal,
		}

		s.respond(w, r, http.StatusOK, data)
	}
}
