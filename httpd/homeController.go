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
	if vb.hasView(path) {
		io.WriteString(w, vb.views[path])
		return
	}

	io.WriteString(w, "Hello! Please feel free to share a secret:"+path)
}
