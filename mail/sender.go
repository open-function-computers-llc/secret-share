package mail

import (
	"errors"

	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

func New(l *logrus.Logger, o string, host string, port int, user string, pass string) (*Mailer, error) {
	m := Mailer{
		logger:   l,
		output:   o,
		host:     host,
		port:     port,
		username: user,
		password: pass,
	}
	if message, ok := m.Validate(); !ok {
		return &m, errors.New(message)
	}

	return &m, nil
}

// Send - check package configuration and send an email or log it
func (m *Mailer) Send(to, url string) error {
	if m.output == "log" {
		m.logger.Info("Sent email to: " + to)
		m.logger.Info(notificiationEmail(url))
		return nil
	}

	if m.output == "smtp" {
		go func() {
			m.logger.Info("Sending email to: " + to)
			message := gomail.NewMessage()
			message.SetHeader("From", "bot@openfunctioncomputers.com")
			message.SetHeader("To", to)
			message.SetHeader("Subject", "A new secret was shared with you")
			message.SetBody("text/html", notificiationEmail(url))

			d := gomail.NewDialer(m.host, m.port, m.username, m.password)
			if err := d.DialAndSend(message); err != nil {
				m.logger.Error(err)
			}
		}()
		return nil
	}

	return errors.New("Invalid state to send emails...")
}
