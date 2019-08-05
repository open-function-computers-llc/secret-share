package main

import (
	"io"
	"net/http"
)

func handleNotFound(w http.ResponseWriter) {
	data, err := Asset("views/notFound.tpl")
	if err != nil {
		http.Error(w, "Asset not found", 404)
	}

	io.WriteString(w, buildView(string(data)))
}
