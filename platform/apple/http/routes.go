package appleHttp

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"

	appleMDM "github.com/mattrax/Mattrax/platform/apple"
)

// Endpoints contains all the dependencies for the Apple MDM HTTP endpoints
type Endpoints struct {
	S      appleMDM.Service
	Logger log.Logger
}

// Routes defines all of the HTTP web routes for the Endpoints. It also mounts the required middlewear
func (h *Endpoints) Routes() chi.Router {
	r := chi.NewRouter()

	r.Put("/checkin", h.checkinHandler())

	return r
}
