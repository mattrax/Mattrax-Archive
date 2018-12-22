package appleHttp

import (
	"github.com/go-chi/chi"
	logging "github.com/op/go-logging"

	appleMDM "github.com/mattrax/Mattrax/platform/apple"
)

var log = logging.MustGetLogger("mattrax_apple_http")

// Endpoints contains all the dependencies for the Apple MDM HTTP endpoints
type Endpoints struct {
	S appleMDM.Service
}

// Routes defines all of the HTTP web routes for the Endpoints. It also mounts the required middlewear
func (h *Endpoints) Routes() chi.Router {
	r := chi.NewRouter()

	//TODO: Certificate (SCEP) Checking Middlewear

	r.Put("/checkin", h.checkinHandler())

	return r
}
