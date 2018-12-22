package windowsHttp

import (
	"github.com/go-chi/chi"
	logging "github.com/op/go-logging"

	windowsMDM "github.com/mattrax/Mattrax/platform/windows"
)

var log = logging.MustGetLogger("mattrax_windows_http")

// Endpoints contains all the dependencies for the Windows MDM HTTP endpoints
type Endpoints struct {
	S windowsMDM.Service
}

// Routes defines all of the HTTP web routes for the Endpoints. It also mounts the required middlewear
func (h *Endpoints) EnrollmentRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/Discovery.svc", h.discoveryGet())
	r.Post("/Discovery.svc", h.discoveryPost())

	return r
}
