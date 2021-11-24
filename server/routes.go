package server

import (
	"net/http"
)

func (s *Server) bindRoutes() {
	// API routes
	http.HandleFunc("/api/recipients", s.handleRecipients())
	http.HandleFunc("/api/secret", s.handleShowSecret())
	http.HandleFunc("/api/store", s.handleStoreSecret())

	frontendFS := http.FileServer(http.FS(s.filesystem))
	http.Handle("/", frontendFS)
}
