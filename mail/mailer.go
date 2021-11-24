package mail

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Mailer struct {
	logger   *logrus.Logger
	baseURL  string
	output   string
	host     string
	port     int
	username string
	password string
}

func (m *Mailer) Validate() (string, bool) {
	if m.logger == nil {
		return "Can't instantiate the mailer without a valid logger", false
	}

	validMailerOutputTypes := []string{"log", "smtp"}
	isValidOutputType := false
	for _, o := range validMailerOutputTypes {
		if o == m.output {
			isValidOutputType = true
		}
	}
	if !isValidOutputType {
		return "Invalid output type, must be one of these: " + strings.Join(validMailerOutputTypes, ", "), false
	}

	if m.output == "smtp" {
		if m.host == "" || m.username == "" || m.password == "" || m.port == 0 {
			return "Invalid SMTP settings", false
		}
	}

	m.logger.Info("Mail config validated. Mail output: " + m.output)
	return "", true
}
