package checkin

import (
	"log"
	"net/http"
)

func checkout(w http.ResponseWriter, r *http.Request) error {
	log.Println("Checkout")
	return nil
}
