package server

import (
	"os"
	"strings"
)

func (s *Server) setUpRecipients() {
	recipientsRaw := os.Getenv("RECIPIENTS")
	s.recipients = parseRecipients(recipientsRaw)
}

func parseRecipients(s string) map[string]string {
	output := map[string]string{}
	individuals := strings.Split(s, ",")
	for _, i := range individuals {
		info := strings.Split(i, ":")
		if len(info) != 2 {
			continue
		}
		output[info[0]] = info[1]
	}
	return output
}
