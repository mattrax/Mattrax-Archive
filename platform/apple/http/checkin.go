package appleHttp

import (
	"net/http"
	"time"

	"github.com/groob/plist"
	mattrax "github.com/mattrax/Mattrax/internal/mattrax"
	appleMDM "github.com/mattrax/Mattrax/platform/apple"
)

func (h *Endpoints) checkinHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd CheckinCommand
		if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
			//log.Info(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch cmd.MessageType {
		case "Authenticate":
			now := time.Now()
			device := &mattrax.Device{
				ID:         mattrax.GenerateDeviceID(),
				Platform:   mattrax.ApplePlatform,
				PlatformID: cmd.UDID,

				PlatformData: appleMDM.PlatformData{
					//State: appleMDM. //TODO: Only Partly Enrolled
				},

				OSVersion:    cmd.OSVersion,
				OSEdition:    cmd.BuildVersion,
				DeviceName:   cmd.DeviceName,
				SerialNumber: cmd.SerialNumber,

				Policies:     []mattrax.PolicyID{},
				Applications: []mattrax.ApplicationID{},

				AssignedTo:     mattrax.UserID(mattrax.GenerateDeviceID()), //TODO: This is temp
				LatestUpdate:   now,
				EnrolledBy:     mattrax.UserID(mattrax.GenerateDeviceID()), //TODO: This is temp
				EnrollmentTime: now,
			}

			err := h.S.Enroll(device)
			if err != nil {
				//log.Error(err) //TODO: Error Handling (Centeral)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

		case "TokenUpdate":
			//TODO: Check if device exists. Cause it should! (Get it for the update func)
			//Modify The Devices PlatformSpecific Tokens
			/*err := h.S.Update(device)
			if err != nil {
				log.Println(err) //TODO: Error Handling (Centeral)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}*/

		case "CheckOut":
		default:
			//log.Error("Invalid MessageType") // TEMP
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

		w.WriteHeader(http.StatusOK)
	}
}
