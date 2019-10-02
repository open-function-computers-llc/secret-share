package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/open-function-computers-llc/secret-share/config"
)

var conf config.Config

func bootstrap() {
	// set some stuff up
	conf = config.Config{}

	conf.Logger = logrus.New()
	conf.Logger.SetOutput(os.Stdout)

	// load ENV variables
	err := godotenv.Load()
	if err != nil {
		logError(err)
		logError("If you want to override system ENV variables, place a .env file in the same directory as the binary file")
	}

	conf.Mail = os.Getenv("MAIL_OUTPUT")
	validMailConfigs := []string{"smtp", "log"}
	for _, choice := range validMailConfigs {
		if os.Getenv("MAIL_OUTPUT") == choice {
			conf.Mail = os.Getenv("MAIL_OUTPUT")
			break
		}
	}
	if conf.Mail == "" {
		logError("Possible bad setting for email output. Please check your ENV for the MAIL_OUTPUT setting. Valid options include:", validMailConfigs)
		conf.Mail = "log"
	}
	if conf.Mail == "smtp" {
		conf.MailHost = os.Getenv("SMTP_HOST")
		conf.MailPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
		conf.MailUserName = os.Getenv("SMTP_USERNAME")
		conf.MailPassword = os.Getenv("SMTP_PASSWORD")
		if conf.MailHost == "" || conf.MailPassword == "" || conf.MailUserName == "" || conf.MailPort == 0 {
			logError("There is a problem with your SMTP settings. Please check the ENV values for [SMTP_HOST SMTP_PORT SMTP_USERNAME SMTP_PASSWORD]")
			conf.Mail = "log"
		}
	}
	logInfo("Email configured to send via " + conf.Mail)

	conf.Port = os.Getenv("PORT")
	if conf.Port == "" {
		conf.Port = "8000"
	}

	conf.BaseURL = os.Getenv("BASE_URL")

	conf.Logger.Info("ready")

	initCache()
}
