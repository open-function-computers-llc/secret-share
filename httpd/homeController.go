package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI[1:] // git rid of the leading slash
	if path == "" {
		path = "home"
	}

	data, err := Asset("views/home.tpl")
	if err != nil {
		http.Error(w, "Could not find valid tpl file", 500)
		return
	}

	io.WriteString(w, buildView(string(data)))
}
