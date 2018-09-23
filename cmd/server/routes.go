package main

import (
	"net/http"

	"github.com/gorilla/mux"
	apple_checkin "github.com/mattrax/Mattrax/internal/apple/checkin"
	apple_enroll "github.com/mattrax/Mattrax/internal/apple/enroll"
	apple_scep "github.com/mattrax/Mattrax/internal/apple/scep"
	apple_server "github.com/mattrax/Mattrax/internal/apple/server"
	"upper.io/db.v3/lib/sqlbuilder"
)

func routes(router *mux.Router, config map[string]string, db sqlbuilder.Database) {
	r := router.Host(config["domain"]).Subrouter()

	//TODO: Caching For These Assets
	//vue.Handle("/", httpHandler(IndexHandler))
	//vue.Handle("/err", httpHandler(ErrorHandler))

	r.PathPrefix("/enroll/apple").Methods("GET").Handler(httpHandler(apple_enroll.Handler(config)))
	r.PathPrefix("/apple/scep").Methods("GET").Handler(httpHandler(apple_scep.GetHandler())) //TODO: Require Special Message In the URL Of These For It To Work (Make Proifile Forging Harder) + Check The SCEP Challenge Key
	r.PathPrefix("/apple/scep").Methods("POST").Handler(httpHandler(apple_scep.PostHandler()))

	r.PathPrefix("/apple/checkin").Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin").Handler(httpHandler(apple_checkin.Handler()))
	r.PathPrefix("/apple/server").Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm").Handler(httpHandler(apple_server.Handler()))

	//r.Handle("/enroll/apple", httpHandler(apple_enroll.Handler(config))).Methods("GET")
	//r.Handle("/apple/scep", httpHandler(apple_scep.GetHandler())).Methods("GET") //TODO: Revoking Devices Cert On Checkout/Removal & Handle SCEP Cert Renewal Without Requiring Reenrollment
	//r.Handle("/apple/scep", httpHandler(apple_scep.PostHandler())).Methods("POST")
	//r.Methods("GET", "POST").Path("/apple/scep").HandlerFunc(ScepHandler) //TEMP
	//r.Handle("/apple/checkin", httpHandler(apple_checkin.Handler())).Methods("PUT") //.Headers("Content-Type", "application/x-apple-aspen-mdm-checkin")
	//r.Handle("/apple/server", httpHandler(apple_server.Handler())).Methods("PUT")   //.Headers("Content-Type", "application/x-apple-aspen-mdm")

	r.PathPrefix("/").Handler(httpHandler(NotFound()))

	// TODO: Certficiate/Header Checking Middleweare For SCEP Verification On The Server And Checkin Routes

	// TODO: Custom Error 404 Page
}

func NotFound() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(404)
		w.Write([]byte("Not Found Here. Keep On Moving. \n"))
		return nil
	}
}

//TODO: Move The Var "operations" From URL Splitting To mux.Vars(r)["operation"]
//TODO: Add Auto Gzipping To httpHandler
