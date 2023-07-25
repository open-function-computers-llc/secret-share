package server

import (
	"io/fs"
	"os"
	"strconv"

	"github.com/open-function-computers-llc/secret-share/mail"
	"github.com/open-function-computers-llc/secret-share/secret"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     *logrus.Logger
	filesystem fs.FS
	port       string
	recipients map[string]string
	secrets    map[string]secret.StorableSecret
	mailer     *mail.Mailer
}

func New(filesystem fs.FS) (*Server, error) {
	s := Server{}
	s.secrets = map[string]secret.StorableSecret{}
	s.filesystem = filesystem

	// initiate logging
	s.logger = logrus.New()

	// initiate mailer
	host := os.Getenv("SMTP_HOST")
	user := os.Getenv("SMTP_USERNAME")
	pass := os.Getenv("SMTP_PASSWORD")
	port := os.Getenv("SMTP_PORT")
	porti, err := strconv.Atoi(port)
	if err != nil {
		return &s, err
	}
	m, err := mail.New(s.logger, os.Getenv("MAIL_OUTPUT"), host, porti, user, pass)
	if err != nil {
		return &s, err
	}
	s.mailer = m

	s.setUpRecipients()
	s.bindRoutes()
	s.port = os.Getenv("PORT")

	ticker(&s)
	return &s, nil

}

func (s *Server) log(messages ...interface{}) {
	s.logger.Info(messages...)
}
