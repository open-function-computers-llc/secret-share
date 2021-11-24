package server

import (
	"net/http"
	"strconv"

	"github.com/open-function-computers-llc/secret-share/secret"
)

func (s *Server) handleStoreSecret() http.HandlerFunc {
	type incomingPayload struct {
		Secret         string
		NumberOfShares string
		ShareWith      []string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			output := map[string]string{
				"error": err.Error(),
			}
			s.send500Json(w, output)
			return
		}

		s.log(r.PostForm.Get("numberOfShares"))
		numShares, err := strconv.Atoi(r.PostForm.Get("numberOfShares"))
		if err != nil {
			output := map[string]string{
				"error": err.Error(),
			}
			s.send500Json(w, output)
			return
		}
		secret, err := secret.StoreNewSecret(r.PostForm.Get("secret"), numShares)
		if err != nil {
			output := map[string]string{
				"error": err.Error(),
			}
			s.send500Json(w, output)
			return
		}
		s.secrets[secret.ID] = secret

		output := map[string]interface{}{
			"message": "Secret stored!",
			"secret":  secret,
		}
		s.sendJson(w, output)
	}
}
