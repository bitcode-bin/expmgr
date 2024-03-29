package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		s.log.WithFields(map[string]interface{}{
			"error": err,
		}).Error("failed to encode")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(buf)
}
