package main

import (
	"github.com/gorilla/mux"
	apple_checkin "github.com/mattrax/Mattrax/internal/apple/checkin"
	apple_scep "github.com/mattrax/Mattrax/internal/apple/scep"
)

func routes(router *mux.Router) {
	r := router.Host("mdm.otbeaumont.me").Subrouter()
	//vue := r.Host("mdm.otbeaumont.me").Methods("GET").Subrouter() //TODO: Load The Domain From The Config

	//vue.Handle("/", httpHandler(IndexHandler))
	//vue.Handle("/err", httpHandler(ErrorHandler))

	//appleMDM.Handle("/enroll/apple", httpHandler(enroll.Handler()))
	r.Handle("/apple/scep", httpHandler(apple_scep.GetHandler())).Methods("GET")
	r.Handle("/apple/scep", httpHandler(apple_scep.PostHandler())).Methods("POST")
	r.Handle("/apple/checkin", httpHandler(apple_checkin.Handler())).Methods("PUT").Headers("Content-Type", "application/x-apple-aspen-mdm-checkin")
	//r.Handle("/apple/server", httpHandler(apple_server.Handler())).Methods("PUT").Headers("Content-Type", "application/x-apple-aspen-mdm")

	// TODO: Certficiate/Header Checking Middleweare For SCEP Verification

	// TODO: Custom Error 404 Page
}
