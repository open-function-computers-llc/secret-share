package server

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/open-function-computers-llc/secret-share/secret"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     *logrus.Logger
	filesystem fs.FS
	port       string
	recipients map[string]string
	secrets    map[string]secret.StorableSecret
}

func New(filesystem fs.FS) Server {
	s := Server{}

	s.secrets = map[string]secret.StorableSecret{}
	s.logger = logrus.New()
	s.filesystem = filesystem

	s.setUpRecipients()
	s.bindRoutes()
	s.port = os.Getenv("PORT")

	return s
}

func (s *Server) log(messages ...interface{}) {
	s.logger.Info(messages...)
}

func (s *Server) Serve() {
	if s.port == "" {
		s.port = "8844"
	}
	s.log("Starting server on port " + s.port)
	http.ListenAndServe(":"+s.port, nil)
}
