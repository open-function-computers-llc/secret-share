package server

import (
	"time"

	"github.com/robfig/cron/v3"
)

func main(s *Server) {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		for _, element := range s.secrets {
			if element.EndTime.Before(time.Now()) {
				delete(s.secrets, element.ID)
			}
		}
	})
}
