package server

import (
	"strings"
)

func (s *Server) sendNotifications(userList, url string) {
	s.log(userList)
	users := strings.Split(userList, ",")

	for name, email := range s.recipients {
		for _, u := range users {
			if u == name {
				s.mailer.Send(email, url)
			}
		}
	}
}
