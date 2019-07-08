package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/open-function-computers-llc/secret-share/secret"

	"github.com/dchest/uniuri"
)

func store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		io.WriteString(w, "Only POST requests can be stored in the system cache")
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, "ParseForm() err: %v"+err.Error())
		return
	}

	formValue := r.FormValue("secret")
	if formValue == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, "The secret you are trying to store can't be empty")
		return
	}

	key := uniuri.NewLen(32)
	s := secret.StorableSecret{
		Value:          formValue,
		RemainingViews: 5,
	}
	cache.Set(key, s)
	output := strings.ReplaceAll(vb.views["saved"], "%%TARGET%%", "/show/"+key)
	io.WriteString(w, output)
}
