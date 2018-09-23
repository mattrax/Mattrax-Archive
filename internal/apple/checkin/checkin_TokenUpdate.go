package checkin

import (
	"log"
	"net/http"
)

func checkinTokenUpdate(w http.ResponseWriter, r *http.Request) error {
	log.Println("Checkin TokenUpdate")

	//TODO: Store The Data To The Database
	//			Update The Last Checkedin Time

	return nil
}
