package main

import (
	"github.com/gorilla/mux"
	apple_checkin "github.com/mattrax/Mattrax/internal/apple/checkin"
	apple_enroll "github.com/mattrax/Mattrax/internal/apple/enroll"
	apple_scep "github.com/mattrax/Mattrax/internal/apple/scep"
	apple_server "github.com/mattrax/Mattrax/internal/apple/server"
	"upper.io/db.v3/lib/sqlbuilder"
)

func routes(router *mux.Router, config map[string]string, db sqlbuilder.Database) {
	r := router.Host("mdm.otbeaumont.me").Subrouter()
	//vue := r.Host("mdm.otbeaumont.me").Methods("GET").Subrouter() //TODO: Load The Domain From The Confi

	//TODO: Caching For These Assets
	//vue.Handle("/", httpHandler(IndexHandler))
	//vue.Handle("/err", httpHandler(ErrorHandler))

	r.Handle("/enroll/apple", httpHandler(apple_enroll.Handler(config))).Methods("GET")
	r.Handle("/apple/scep", httpHandler(apple_scep.GetHandler())).Methods("GET")
	r.Handle("/apple/scep", httpHandler(apple_scep.PostHandler())).Methods("POST")
	r.Handle("/apple/checkin", httpHandler(apple_checkin.Handler())).Methods("PUT").Headers("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.Handle("/apple/server", httpHandler(apple_server.Handler())).Methods("PUT").Headers("Content-Type", "application/x-apple-aspen-mdm")

	// TODO: Certficiate/Header Checking Middleweare For SCEP Verification

	// TODO: Custom Error 404 Page
}

//TODO: Move The Var "operations" From URL Splitting To mux.Vars(r)["operation"]
//TODO: Add Auto Gzipping To httpHandler
