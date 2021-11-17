package main

import (
	"embed"
	"io/fs"

	_ "github.com/joho/godotenv/autoload"
	"github.com/open-function-computers-llc/secret-share/server"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	// static assets for Vue app
	stripped, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		panic(err)
	}

	s := server.New(stripped)
	s.Serve()
}
