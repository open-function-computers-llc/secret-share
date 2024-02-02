package server

import (
	"net/http"
	"os"
	"strconv"

	"github.com/open-function-computers-llc/secret-share/secret"
)

func (s *Server) handleStoreSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			output := map[string]string{
				"error": err.Error(),
			}
			s.send500Json(w, output)
			return
		}

		numShares, err := strconv.Atoi(r.PostForm.Get("numberOfShares"))
		if err != nil {
			output := map[string]string{
				"error": err.Error(),
			}
			s.send500Json(w, output)
			return
		}
		viewTime, err := strconv.Atoi(r.PostForm.Get("time"))
		secret, err := secret.StoreNewSecret(r.PostForm.Get("secret"), numShares, viewTime)
		if err != nil {
			output := map[string]string{
				"error": err.Error(),
			}
			s.send500Json(w, output)
			return
		}
		s.secrets[secret.ID] = secret

		newSecretURL := os.Getenv("BASE_URL") + "/#/show/" + secret.ID
		s.sendNotifications(r.PostForm.Get("shareWith"), newSecretURL)

		output := map[string]interface{}{
			"message": "Secret stored!",
			"secret":  secret,
			"url":     newSecretURL,
		}
		s.sendJson(w, output)
	}
}
