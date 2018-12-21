package windowsHttp

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"

	windowsMDM "github.com/mattrax/Mattrax/platform/windows"
)

// Endpoints contains all the dependencies for the Windows MDM HTTP endpoints
type Endpoints struct {
	S      windowsMDM.Service
	Logger log.Logger
}

// Routes defines all of the HTTP web routes for the Endpoints. It also mounts the required middlewear
func (h *Endpoints) Routes() chi.Router {
	r := chi.NewRouter()

	//r.Put("/checkin", h.checkinHandler())

	return r
}
