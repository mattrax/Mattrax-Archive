package checkin

import (
	"net/http"

	"github.com/groob/plist"
	"github.com/kataras/muxie"
	"github.com/rs/zerolog/log"
)

// MountEndpoints mounts the http endpoints for the service
func (svc *Service) MountEndpoints(mux *muxie.Mux) {

	mux.Handle("/apple/checkin", muxie.Methods().
		HandleFunc(http.MethodPut, checkinHandler())) // TODO: Check The MDM Signature Middleware
}

func checkinHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/x-apple-aspen-mdm-checkin; charset=UTF-8" {
			log.Debug().Str("protocol", "apple/checkin").Str("Content-Type", r.Header.Get("Content-type")).Msg("FAILED Checkin. Incorrect 'Content-Type' header.")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var cmd CheckinCommand
		if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
			log.Debug().Str("protocol", "apple/checkin").Msg("FAILED Checkin. Invalid Request Body.")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch cmd.MessageType {
		case "Authenticate":
			Authenticate(cmd, w)
		case "TokenUpdate":
			panic("not implemented (TokenUpdate)")
		case "CheckOut":
			panic("not implemented (CheckOut)")
		default:
			log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Msg("FAILED Checkin. Invalid MessageType.")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

func Authenticate(cmd CheckinCommand, w http.ResponseWriter) {
	if cmd.OSVersion == "" || cmd.BuildVersion == "" || cmd.ProductName == "" || cmd.SerialNumber == "" || cmd.IMEI == "" || cmd.MEID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: Convert The Req to A Device Struct
	// TODO: Send That Struct To The Store
	log.Info().Interface("checkin", cmd).Msg("A Device Checked In") // TEMP

	w.WriteHeader(http.StatusOK)
}
