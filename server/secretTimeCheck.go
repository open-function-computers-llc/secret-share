package server

import (
	"time"
)

func ticker(s *Server) {
	ticker := time.NewTicker(1 * time.Hour)

	// Creating channel using make
	tickerChan := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChan:
				return

			case <-ticker.C:
				checkAll(s)
			}
		}
	}()
}

func checkAll(s *Server) {
	for _, element := range s.secrets {
		if element.EndTime.Before(time.Now()) {
			delete(s.secrets, element.ID)
		}
	}
}
