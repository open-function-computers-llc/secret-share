package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/open-function-computers-llc/secret-share/mail"
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

	s := r.FormValue("secret")
	if s == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, "The secret you are trying to store can't be empty")
		return
	}

	views := r.FormValue("viewCount")
	if views == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, "The total number of times you'd like to share this secret can't be empty")
		return
	}
	vInt, err := strconv.Atoi(views)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, "The total number of times you'd like to share this secret must be a number")
		return
	}

	key := uniuri.NewLen(32)
	obj := secret.StorableSecret{
		Value:          s,
		RemainingViews: vInt,
		Expires:        time.Now(),
	}
	cache.Set(key, obj)

	responseBody, err := Asset("views/saved.tpl")
	if err != nil {
		handleNotFound(w)
		return
	}

	output := buildView(string(responseBody))

	output = strings.ReplaceAll(output, "%%TARGET%%", "/show/"+key)
	output = strings.ReplaceAll(output, "%%TITLE%%", "Your secret was saved")
	io.WriteString(w, output)

	// send an email?
	if r.FormValue("autoshare") == "on" {
		mail.Send(conf, key)
	}
}
