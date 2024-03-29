package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitcode-bin/expmgr/inmemory"
	"github.com/bitcode-bin/expmgr/logger"
)

const baseURL string = "/api"

func newWallet(startingBalance int) *inmemory.Wallet {
	return inmemory.NewWallet(startingBalance)
}

func newServer() *server {
	s := &server{
		log:    logger.NewNoopLogger(),
		wallet: newWallet(0),
	}
	s.Init()
	return s
}

func TestBalanceGet(t *testing.T) {
	s := newServer()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", baseURL+"/balance", nil)

	s.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("response code: %d", w.Code)
	}

	want := string([]byte(`{"balance":0}`))
	if w.Body.String() != want {
		t.Fatalf("want=%s, got=%s", want, w.Body.String())
	}
}

func TestIncomePost(t *testing.T) {
	s := newServer()

	reqBody := []byte(`{"income": 200}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", baseURL+"/income", bytes.NewBuffer(reqBody))

	s.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("response code: %d", w.Code)
	}

	want := string([]byte(`{"balance":200}`))
	if w.Body.String() != want {
		t.Fatalf("want=%s, got=%s", want, w.Body.String())
	}

	t.Run("invalidJson", func(t *testing.T) {
		reqBody = []byte(`{"income"}`)
		r := httptest.NewRequest("POST", baseURL+"/income", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("response code: %d", w.Code)
		}
	})

	t.Run("negativeAmount", func(t *testing.T) {
		reqBody = []byte(`{"income": -100}`)
		r := httptest.NewRequest("POST", baseURL+"/income", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("response code: %d", w.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			isNil(t, err)
		}

		errMsg := response["error"].(string)
		want := "negative income is not allowed"
		if errMsg != want {
			t.Fatalf("error msg: want=%s, got=%s", want, errMsg)
		}
	})
}

func TestExpensePost(t *testing.T) {
	s := newServer()

	reqBody := []byte(`{"expense": 200}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", baseURL+"/expense", bytes.NewBuffer(reqBody))

	s.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("response code: %d", w.Code)
	}

	want := string([]byte(`{"balance":-200}`))
	if w.Body.String() != want {
		t.Fatalf("want=%s, got=%s", want, w.Body.String())
	}

	t.Run("invalidJson", func(t *testing.T) {
		reqBody = []byte(`{"expense"}`)
		r := httptest.NewRequest("POST", baseURL+"/expense", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("response code: %d", w.Code)
		}
	})
	t.Run("negativeAmount", func(t *testing.T) {
		reqBody = []byte(`{"expense": -100}`)
		r := httptest.NewRequest("POST", baseURL+"/expense", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("response code: %d", w.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			isNil(t, err)
		}

		errMsg := response["error"].(string)
		want := "negative expense is not allowed"
		if errMsg != want {
			t.Fatalf("error msg: want=%s, got=%s", want, errMsg)
		}
	})
}

func TestTransactions(t *testing.T) {
	s := newServer()
	s.wallet = newWallet(500)

	// use slice, not map, since
	// these tests MUST be called in order
	tests := []struct {
		name    string
		method  string
		url     string
		reqBody []byte

		wantCode    int
		wantResBody []byte
	}{
		{
			"getBalance",
			"GET",
			"/balance",
			nil,
			http.StatusOK,
			[]byte(`{"balance":500}`),
		},

		{
			"addIncome",
			"POST",
			"/income",
			[]byte(`{"income": 700}`),
			http.StatusOK,
			[]byte(`{"balance":1200}`),
		},

		{
			"addExpense",
			"POST",
			"/expense",
			[]byte(`{"expense": 300}`),
			http.StatusOK,
			[]byte(`{"balance":900}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.method, baseURL+test.url, bytes.NewBuffer(test.reqBody))

			s.ServeHTTP(w, r)

			if w.Code != test.wantCode {
				t.Fatalf("response code: want=%d, got=%s", test.wantCode, http.StatusText(w.Code))
			}

			want := string(test.wantResBody)
			if w.Body.String() != want {
				t.Fatalf("want=%s, got=%s", want, w.Body.String())
			}
		})
	}
}
