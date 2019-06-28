package main

import (
	"net/http"

	"github.com/urfave/cli"
)

func runApp(c *cli.Context) error {
	buildViews()
	for path, handler := range webRoutes() {
		http.HandleFunc(path, handler)
	}
	http.ListenAndServe(":8000", nil)
	return nil
}

func webRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/"] = home
	routes["/show/"] = lookup

	return routes
}
