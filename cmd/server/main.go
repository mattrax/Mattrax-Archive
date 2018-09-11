package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router

	/*
		  config someConfig
			db     *someDatabase
			email  EmailSender
			logger *someLogger
	*/
}

func main() {
	server := Server{
		router: mux.NewRouter(),
	}

	server.routes()

	log.Fatal(http.ListenAndServe(":8000", server.router))

	//TODO: Better Failure Handline, HTTPS, Etc
	//TODO Auto HTTPS: https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
	//TODO Subdomains: https://translate.google.com/translate?hl=en&sl=am&tl=en&u=http%3A%2F%2Fcodepodu.com%2Fsubdomains-with-golang%2F&anno=2
	//TODO: Centeral Error Handling For All The Routes
}
