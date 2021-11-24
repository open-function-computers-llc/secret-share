package server

import (
	"net/http"
	"sort"
)

func (s *Server) handleRecipients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipients := []string{}
		for name := range s.recipients {
			recipients = append(recipients, name)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(recipients)))

		s.sendJson(w, recipients)
	}
}
