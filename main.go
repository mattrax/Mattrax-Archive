/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is The Core. It Loads The Database, Logging and The Webserver Components.
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	// External Deps
	"github.com/gorilla/handlers" // HTTP Handlers
	"github.com/gorilla/mux"      // HTTP Router

	// Internal Functions
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database"      //Mattrax Database
	mlg "github.com/mattrax/Mattrax/internal/logging"       //Mattrax Logging

	// Internal Modules
	"github.com/mattrax/Mattrax/appleMDM" // The Apple MDM Module
	//"github.com/mattrax/Mattrax/windowsMDM" // The Windows MDM Module
)

var ( // Get The Internal State
	pgdb = mdb.GetDatabase()
	log = mlg.GetLogger()
	config = mcf.GetConfig()
	srv *http.Server // The Webserver
)

// This Function Handles The Webserver And Cleanup When Exitting The Application
func main() {
	//Load The Modules
	appleMDM.Init()
	//windowsMDM.Init()

	//Webserver Routes
	router := mux.NewRouter()
	r := router.Host(config.Domain).Subrouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/enroll", enrollmentHandler).Methods("GET")
	appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())

	//Start The Webserver (In The Background)
	go func() { startWebserver(router) }()
	log.Info("The Mattrax Webserver Is Listening At " + fmt.Sprintf("%v:%v", "0.0.0.0", config.Port))

	//Wait Until Shutting Down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//Cleanup
	log.Info("Mattrax is Shutting Down...")
	mdb.Cleanup() //Shutdown The Database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx) // Shutdown The Webserver
	os.Exit(0)        //Exit Successfuly
}

/* Database Initialisation */ //FIXME: Make This Entire Section Work
func correctSchema() bool {
	if _, err := pgdb.Exec("SELECT * FROM devices"); err != nil {
		//Find Out If Error Was Database Table If Not Log Fatal
		return false
	}
	return true
}

func initDatabaseSchema() {
	panic("This will create the database schema")
}

/* The Webserver */
func startWebserver(router *mux.Router) {
	var handler http.Handler
	if config.Verbose {
		handler = verboseRequestLogger(handlers.CORS()(router))
	} else {
		handler = handlers.CORS()(router)
	}

	srv = &http.Server{
		Addr:         fmt.Sprintf("%v:%v", "0.0.0.0", config.Port), //FIXME Configurable Listen IP (Optional)
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func verboseRequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Request: " + r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		handler.ServeHTTP(w, r)
	})
}

/* HTTP Handling Routes */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mattrax MDM Server!")
}

func enrollmentHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: This Will Show Interface To Guide User Through Enrollment
	http.Redirect(w, r, "/apple/enroll", 301)
}

/* The End */
//Now

// TOMORROW:
// Redo The errors.go, apns.go and appleAPI.go
// /server route


//Remove All Evnly Formatted (Whietspaced) Structs
// APNS Make Device update Every * Days -> Configurable Checkin Timeout Cause Big Deployments Will Need Longer
// Does os.Exit(int) Run The Cleanup Functions If Not make it
// For Test, Test "gofmt -s -w ." And Break If It Does't Parse 100%
// FUTURE FEATURE: Redo Separator Between Blocks Of Function -> They Don't Stand Out Enought
// TODO: Contant Pinging Database To Stop HTTP Soon As It Stops Connecting
// TODO: Add "* Package Description: Something" To The Header of All Of The Files In a Package

// FIXME: Handle Device Removing From MDM Without Being In The Database

// Log Wiping After Restart (Fix That)
//      Redo Logging For Subfiles To Use The Features Of The New System
// TODO: Log File Roation So The Log Files Doesn't Get To Big
//      Config Options For External Logging Server
// TODO: Godoc Documentation Throughout Code -> On All Functions And Structs
// TODO: GoDEP Package Management
// TODO: Check Line Ending for ; and Remove (Maybe Add Test To Check For Them)
// TODO: Redo/Add Separators Between Part Of Code That Are Clean And Easy To See
//   Func Based (struct, err) Error Handling Insead Of *struct (Remove All Of *struct)
// Using The New Logging Information Log Structs/Varibles/State To mMake Debugging Easier (On log.Debug Only)
//TODO: Track Device Events (When They Enrolled, When They Checkin)
// Add Time Register To The Device Information
//  Lint/Test Check That Stuff Isn't Exports (Capitalised) Unless Used
// TODO: Update The Documentation URL In Each File To Match What is in It
// Disable File Logging, Log To Syslog

//Far In The Future
//   Build Tests -> For All Function And Routes (Fake Device Requests/Response Verifying)
//	 Optimisng Performance
// IDEA: HTTPS And HTTP Support With Automatic Redirection Between
//      Certbot ACME Built In For Automatically Issuing And Renewing Cert
// IDEA: Built In IP Whitelist For Access To The Admin Area
