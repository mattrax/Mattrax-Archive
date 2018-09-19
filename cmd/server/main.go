package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes(r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Mattrax Started Listening At ") //TODO: Finish This Message
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("App Failed to start with the Error:", err.Error())
	}
}
