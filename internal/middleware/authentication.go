package middleware

import (
	"encoding/json"
	"net/http"

	mattrax "github.com/mattrax/Mattrax/internal"
)

type AuthConfig struct {
	LoginEndpoint string
	AuthService   mattrax.AuthService
}

func AuthRequireRoles(config AuthConfig, roles []string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if true { // TODO: Not Logged In
			if r.URL.String() == config.LoginEndpoint {
				h(w, r)
			}

			http.Redirect(w, r, config.LoginEndpoint, http.StatusTemporaryRedirect)
		}
	}
}

func LoginAPI(config AuthConfig) http.HandlerFunc {
	type Req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			panic(err) // TODO: Handle Error
		}

		if err := config.AuthService.VerifyLogin(req.Username, req.Password); err != nil { // TODO: Correct Login
			w.WriteHeader(http.StatusOK)
		} else if err == mattrax.ErrInvalidLogin {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			panic(err) // TODO: Handle Error
		}

		// TODO: Log The Login + IP Addr & User Agent Of Client
	}
}
