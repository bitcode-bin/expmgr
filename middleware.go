package main

import (
	"context"
	"net/http"

	"github.com/bitcode-bin/expmgr/logger"
	"github.com/rs/xid"
)

type Middleware func(http.Handler) http.Handler

const (
	CtxKeyRequestID string = "requestId"
)

func requestLogger(logger logger.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			logger.WithFields(map[string]interface{}{
				"method":    r.Method,
				"host":      r.Host,
				"path":      r.URL.String(),
				"requestId": r.Context().Value(CtxKeyRequestID).(string),
			}).Info("")

			h.ServeHTTP(w, r)
		})
	}
}

func requestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := xid.New()
		ctx := context.WithValue(r.Context(), CtxKeyRequestID, reqID.String())
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
