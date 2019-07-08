package main

import (
	"fmt"
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

	fmt.Println("key: " + key)
	var output string

	s, err := cache.Get(key)
	sCasted := s.(secret.StorableSecret)
	if sCasted.RemainingViews < 1 {
		output = vb.views["notFound"]
	} else {
		sCasted.RemainingViews--
		cache.Set(key, sCasted)
		output = strings.ReplaceAll(vb.views["show"], "%%SECRET%%", sCasted.Value)
		output = strings.ReplaceAll(output, "%%REMAININGVIEWS%%", strconv.Itoa(sCasted.RemainingViews))
	}

	if err != nil {
		output = vb.views["notFound"]
	}

	io.WriteString(w, output)
}
