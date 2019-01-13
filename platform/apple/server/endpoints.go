package server

import (
	"github.com/kataras/muxie"
)

// MountEndpoints mounts the http endpoints for the service
func (svc *Service) MountEndpoints(mux *muxie.Mux) {
	//mux.Handle("/apple/scep", muxie.Methods().
	//	HandleFunc(http.MethodGet, handler()))
}
