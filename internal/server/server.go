package server

import (
	"net/http"

	"github.com/go-chi/chi"
	logging "github.com/op/go-logging"

	appleMDM "github.com/mattrax/Mattrax/platform/apple"
	appleHttp "github.com/mattrax/Mattrax/platform/apple/http"
	windowsMDM "github.com/mattrax/Mattrax/platform/windows"
	windowsHttp "github.com/mattrax/Mattrax/platform/windows/http"
)

// Server holds the dependencies for the HTTP server.
type server struct {
	AppleMDM   appleMDM.Service
	WindowsMDM windowsMDM.Service

	router chi.Router
}

var log = logging.MustGetLogger("mattrax_http")

// New returns a new HTTP server.
func New(as appleMDM.Service, ws windowsMDM.Service) *server { //TEMP logger *logging.Logger
	s := &server{
		AppleMDM:   as,
		WindowsMDM: ws,
	}

	r := chi.NewRouter()

	//r.NotFound(handlerFn) // TODO:
	//r.MethodNotAllowed(handlerFn)

	//TODO: Copy the ------ r.Use(accessControl)

	r.Route("/apple", func(r chi.Router) {
		h := appleHttp.Endpoints{s.AppleMDM}
		r.Mount("/", h.Routes())
	})

	r.Route("/EnrollmentServer", func(r chi.Router) { // Windows requires some routes with that prefix
		h := windowsHttp.Endpoints{s.WindowsMDM}
		r.Mount("/", h.EnrollmentRoutes())
	})

	s.router = r
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
