package server

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     *logrus.Logger
	router     *mux.Router
	filesystem fs.FS
	port       string // web host port number
}

func New(filesystem fs.FS) Server {
	s := Server{}

	s.logger = logrus.New()
	s.filesystem = filesystem
	s.port = os.Getenv("PORT")

	return s
}

func (s *Server) log(messages ...interface{}) {
	s.logger.Info(messages...)
}

func (s *Server) Serve() {
	// API routes
	http.Handle("/api/test", s.handleTest())

	frontendFS := http.FileServer(http.FS(s.filesystem))
	http.Handle("/", frontendFS)

	if s.port == "" {
		s.port = "8844"
	}
	s.log("Starting server on port " + s.port)
	http.ListenAndServe(":"+s.port, nil)
}
