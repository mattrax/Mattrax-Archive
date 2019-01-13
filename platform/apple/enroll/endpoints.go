package enroll

import (
	"net/http"

	"github.com/kataras/muxie"
	"github.com/rs/zerolog/log"
)

// MountEndpoints mounts the http endpoints for the service
func (svc *Service) MountEndpoints(mux *muxie.Mux) {
	mux.Handle("/apple/enroll", muxie.Methods().
		HandleFunc(http.MethodGet, getEnrollmentProfile(svc)))
}

func getEnrollmentProfile(svc *Service) http.HandlerFunc {
	profile, err := svc.SignedEnrollmentProfile()
	if err != nil {
		log.Fatal().Err(err).Msg("Error Creating The Enrollment Profile")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-apple-aspen-config")
		w.Write(profile)
	}
}
