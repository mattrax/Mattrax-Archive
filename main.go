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
	"github.com/mattrax/Mattrax/windowsMDM" // The Windows MDM Module
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
	windowsMDM.Init()

	//Webserver Routes
	router := mux.NewRouter()
	r := router.Host(config.Domain).Subrouter()

	//Webroutes -> Management Interface
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Mattrax MDM Server!") }).Methods("GET")
	r.HandleFunc("/enroll", enrollmentHandler).Methods("GET")

	//Webroutes -> Modules
	appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())
	windowsMDM.Mount(r.PathPrefix("/windows/").Subrouter(), router.Host(config.EEDomain).Subrouter())

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
//TODO Docs
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
//TODO Docs
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

//TODO Docs
func verboseRequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Request: " + r.RemoteAddr + " " + r.Method + " " + r.URL.String())


		var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
		//log.Info(keys)


		handler.ServeHTTP(w, r)
	})
}









//TODO Docs
func enrollmentHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: This Will Show Interface To Guide User Through Enrollment
	http.Redirect(w, r, "/apple/enroll", 301)
}

//TODO: /apple doesn;t work only /apple/ in browser
