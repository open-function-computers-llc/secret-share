package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/open-function-computers-llc/secret-share/secret"
)

func lookup(w http.ResponseWriter, r *http.Request) {
	secretParts := strings.SplitAfter(r.RequestURI, "/")
	key := secretParts[2]

	if key == "" {
		http.Redirect(w, r, "/", 301)
		return
	}

	s, err := cache.Get(key)

	if err != nil {
		handleNotFound(w)
		return
	}

	sCasted := s.(secret.StorableSecret)
	if sCasted.RemainingViews < 1 {
		handleNotFound(w)
		return
	}

	data, err := Asset("views/show.tpl")
	if err != nil {
		handleNotFound(w)
		return
	}

	sCasted.RemainingViews--
	cache.Set(key, sCasted)
	output := strings.ReplaceAll(string(data), "%%SECRET%%", sCasted.Value)
	output = strings.ReplaceAll(output, "%%REMAININGVIEWS%%", strconv.Itoa(sCasted.RemainingViews))

	io.WriteString(w, buildView(output))
}
