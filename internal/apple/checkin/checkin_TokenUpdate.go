package checkin

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/mattrax/Mattrax/models"
)

func checkinTokenUpdate(w http.ResponseWriter, request models.AppleAuthenticateDetails, db *sqlx.DB, device_exists bool, device models.Device) error {
	log.Println("Checkin TokenUpdate")

	//TODO: Store The Data To The Database
	//			Update The Last Checkedin Time

	return nil
}
