package main

import (
	"io"
	"net/http"
)

func bundledAsset(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	data, err := Asset(path[1:])
	if err != nil {
		http.Error(w, "Could not find valid bundled file", 500)
		return
	}

	// manual mime type fixes (css is not text/pain)
	if path[len(path)-3:] == "css" {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}

	io.WriteString(w, string(data))
}
