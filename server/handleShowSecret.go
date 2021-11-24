package server

import "net/http"

func (s *Server) handleShowSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.Form.Get("id")
		if id == "" {
			output := map[string]string{
				"error": "Invalid request. Make sure you define the `id` query variable.",
			}
			s.send404Json(w, output)
			return
		}
		foundSecret, ok := s.secrets[id]
		if !ok {
			output := map[string]string{
				"error": "Secret " + id + " not found",
			}
			s.send404Json(w, output)
			return
		}
		foundSecret.RemainingViews = foundSecret.RemainingViews - 1
		if foundSecret.RemainingViews < 0 {
			delete(s.secrets, foundSecret.ID)
			output := map[string]string{
				"error": "Secret " + id + " not found",
			}
			s.send404Json(w, output)
			return
		}
		s.secrets[foundSecret.ID] = foundSecret

		s.sendJson(w, foundSecret)
	}
}
