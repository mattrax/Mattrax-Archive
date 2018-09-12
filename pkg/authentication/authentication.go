package authentication

import (
	"encoding/json"
	"net/http"
)

func LoginHandler() http.HandlerFunc {
  type loginResponse struct {
    success bool `json:"success"`
    error_msg string `json:"error_msg"`
  }

	return func(w http.ResponseWriter, r *http.Request) {

		//io.WriteString(w, "User Login API")
    json.NewEncoder(w).Encode(loginResponse{
      success: true,
      error_msg: "",
    })
	}
}
