package authentication

import (
	"encoding/json"
	"net/http"
)

func LoginHandler() http.HandlerFunc {
  type loginResponse struct {
    Success bool `json:"Success"`
    ErrorCode int `json:"ErrorCode,omitempty"`
    Error string `json:"Error,omitempty"`
  }

	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    /*json.NewEncoder(w).Encode(loginResponse{
      Success: true,
      Error: "",
    })*/
    json.NewEncoder(w).Encode(loginResponse{
      Success: false,
      ErrorCode: 1,
      Error: "Could Not Create The User!",
    })
	}
}
