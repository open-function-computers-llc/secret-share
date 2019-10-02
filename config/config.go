package config

import (
	"github.com/sirupsen/logrus"
)

// Config the settings we need to correctly run the program
type Config struct {
	BaseURL      string
	Logger       *logrus.Logger
	Port         string
	Mail         string // set this to log|smtp
	MailPort     int
	MailHost     string
	MailUserName string
	MailPassword string
}
