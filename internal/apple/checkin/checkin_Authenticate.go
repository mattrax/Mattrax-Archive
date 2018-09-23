package checkin

import (
	"log"
	"net/http"
)

func checkinAuthenticate(w http.ResponseWriter, r *http.Request) error {
	log.Println("Checkin Authenticate")
	return nil
}
