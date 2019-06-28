package main

import (
	"io"
	"net/http"
	"strings"
)

func lookup(w http.ResponseWriter, r *http.Request) {
	secretParts := strings.SplitAfter(r.RequestURI, "/")

	output := strings.ReplaceAll(vb.views["show"], "%%SECRET%%", secretParts[2])
	io.WriteString(w, output)
}
