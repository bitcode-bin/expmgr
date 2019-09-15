package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleBalanceGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		reqID := r.Context().Value(CtxKeyRequestID).(string)

		var req request

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			s.log.WithFields(map[string]interface{}{
				"requestId": reqID,
				"error":     err,
			}).Error("failed to decode request")

			data := map[string]interface{}{"error": "invalid json"}
			s.respond(w, r, http.StatusBadRequest, data)
			return
		}

		bal := s.wallet.AddIncome(req.Income)

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
		reqID := r.Context().Value(CtxKeyRequestID).(string)

		var req request

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			s.log.WithFields(map[string]interface{}{
				"requestId": reqID,
				"error":     err,
			}).Error("failed to decode request")

			data := map[string]interface{}{"error": "invalid json"}
			s.respond(w, r, http.StatusBadRequest, data)
			return
		}

		bal := s.wallet.AddExpense(req.Expense)

		data := map[string]interface{}{
			"balance": bal,
		}

		s.respond(w, r, http.StatusOK, data)
	}
}
