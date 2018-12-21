package appleHttp

import (
	"log"
	"net/http"
	"time"

	"github.com/groob/plist"
	mattrax "github.com/mattrax/Mattrax/internal/mattrax"
)

func (h *Endpoints) checkinHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd CheckinCommand
		if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
			log.Println(err)
			return
		}

		switch cmd.MessageType {
		case "Authenticate":
			device := &mattrax.Device{
				//EnrolledBy
				EnrollmentTime: time.Now(),
			}
			if err := device.GenerateID(); err != nil {
				log.Println(err) //TODO: Error Handling (Centeral)
			}

			err := h.S.Enroll(device)
			if err != nil {
				log.Println(err) //TODO: Error Handling (Centeral)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

		case "TokenUpdate":
		case "CheckOut":
		default:
			log.Println("Invalid MessageType") // TEMP
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

		w.WriteHeader(http.StatusOK)
	}
}
