package windows

import (
	"github.com/gorilla/mux"
	"github.com/mattrax/Mattrax/platform/windows/endpoints"
)

func (w windowsMDM) Routes(r *mux.Router) error {
	// TODO: Must Not be sent Chunked -> Fix With Middleware maybe
	// TODO: Use Common Naming Scheme For Endpoints
	r.HandleFunc("/EnrollmentServer/Discovery.svc", endpoints.DiscoveryGet()).Methods("GET")                // The Device Checking The Server Exists
	r.HandleFunc("/EnrollmentServer/Discovery.svc", endpoints.DiscoveryPost(w.AuthService)).Methods("POST") // Returns The Servers Details
	r.HandleFunc("/EnrollmentServer/PolicyService.svc", endpoints.EnrollmentPolicyEndpoint()).Methods("POST")

	return nil
}
