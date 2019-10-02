package main

import (
	"net/http"

	"github.com/urfave/cli"
)

func runApp(c *cli.Context) error {
	bootstrap()

	for path, handler := range webRoutes() {
		http.HandleFunc(path, handler)
	}
	http.HandleFunc("/assets/", bundledAsset)

	conf.Logger.Info("Now serving on port " + conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
	return nil
}

func webRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/"] = home
	routes["/store"] = store
	routes["/show/"] = lookup

	return routes
}
