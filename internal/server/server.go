package server

import (
	"net/http"

	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"

	appleMDM "github.com/mattrax/Mattrax/platform/apple"
	appleHttp "github.com/mattrax/Mattrax/platform/apple/http"
	windowsMDM "github.com/mattrax/Mattrax/platform/windows"
	windowsHttp "github.com/mattrax/Mattrax/platform/windows/http"
)

// Server holds the dependencies for the HTTP server.
type server struct {
	AppleMDM   appleMDM.Service
	WindowsMDM windowsMDM.Service

	Logger kitlog.Logger
	router chi.Router
}

// New returns a new HTTP server.
func New(as appleMDM.Service, ws windowsMDM.Service, logger kitlog.Logger) *server {
	s := &server{
		AppleMDM:   as,
		WindowsMDM: ws,
		Logger:     logger,
	}

	r := chi.NewRouter()

	//r.NotFound(handlerFn) // TODO:
	//r.MethodNotAllowed(handlerFn)

	//TODO: Copy the ------ r.Use(accessControl)

	r.Route("/apple", func(r chi.Router) {
		h := appleHttp.Endpoints{s.AppleMDM, s.Logger}
		r.Mount("/", h.Routes())
	})

	r.Route("/windows", func(r chi.Router) {
		h := windowsHttp.Endpoints{s.WindowsMDM, s.Logger}
		r.Mount("/", h.Routes())
	})

	s.router = r
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
