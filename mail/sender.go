package mail

import (
	"github.com/open-function-computers-llc/secret-share/config"
	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

// log - pointer to a shared logger
var log *logrus.Logger

// Send - check package configuration and send an email or log it
func Send(c config.Config, key string) {
	if c.Logger == nil || c.Mail == "" {
		panic("Logging and Config not set correctly")
	}
	log = c.Logger

	if c.Mail == "log" {
		log.Info(notificiationEmail(c, key))
		return
	}

	go func() {
		log.Info("Sending email..." + notificiationEmail(c, key))
		m := gomail.NewMessage()
		m.SetHeader("From", "bot@openfunctioncomputers.com")
		m.SetHeader("To", "kurtis@openfunctioncomputers.com")
		m.SetHeader("Subject", "A new secret was shared with you")
		m.SetBody("text/html", notificiationEmail(c, key))

		d := gomail.NewDialer(c.MailHost, c.MailPort, c.MailUserName, c.MailPassword)
		if err := d.DialAndSend(m); err != nil {
			log.Error(err)
		}
	}()
}
