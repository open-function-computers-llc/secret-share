package server

import "net/http"

func (s *Server) handleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log("test hit!")
		w.Write([]byte("test!"))
	}
}
