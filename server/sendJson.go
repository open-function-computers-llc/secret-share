package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) sendJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		bytes = []byte("Error: " + err.Error())
	}
	w.Write(bytes)
}

// wrap the standard json response in a 404 header
func (s *Server) send404Json(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusNotFound)
	s.sendJson(w, data)
}

// wrap the standard json response in a 404 header
func (s *Server) send500Json(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	s.sendJson(w, data)
}
