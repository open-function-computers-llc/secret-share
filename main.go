package main

import (
	"embed"
	"errors"
	"io/fs"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/open-function-computers-llc/secret-share/server"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	err := checkEnv()
	if err != nil {
		panic(err)
	}

	// static assets for Vue app
	stripped, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		panic(err)
	}

	s, err := server.New(stripped)
	if err != nil {
		panic(err)
	}
	s.Serve()
}

func checkEnv() error {
	requiredEnv := []string{
		"BASE_URL",
		"MAIL_OUTPUT",
		"SMTP_HOST",
		"SMTP_PORT",
		"SMTP_USERNAME",
		"SMTP_PASSWORD",
	}

	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			return errors.New("required env: " + env + " missing")
		}
	}
	return nil
}
