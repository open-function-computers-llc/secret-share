package server

import (
	"net/http"
)

func (s *Server) bindRoutes() {
	// API routes
	http.Handle("/api/recipients", http.HandlerFunc(s.handleRecipients()))
	http.Handle("/api/secret", s.handleShowSecret())
	http.Handle("/api/store", s.handleStoreSecret())
	http.Handle("/api/theme", s.handleTheme())

	// branding files, all configured through the env
	http.Handle("/theme.css", serveThemeFile(s.theme.cssPath))
	http.Handle("/logo", serveThemeFile(s.theme.logoPath))
	http.Handle("/favicon.ico", serveThemeFile(s.theme.faviconPath))

	frontendFS := http.FileServer(http.FS(s.filesystem))
	http.Handle("/", frontendFS)
}

func (s *Server) Serve() {
	if s.port == "" {
		s.port = "8844"
	}
	s.log("Starting server on port " + s.port)
	http.ListenAndServe(":"+s.port, nil)
}
