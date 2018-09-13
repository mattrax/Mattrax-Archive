package main

import (
	"log"
	"net/http"

	_ "../../pkg/apple/management_payloads" //TEMP
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
	// Load The Config File
	/*file, err := os.Open(filename) if err != nil {  return err }
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {  return err }*/
	//TODO: Create It If Not THere

	server := Server{
		router: mux.NewRouter(),
	}

	server.routes()

	log.Println("Listening Port 8000")
	log.Fatal(http.ListenAndServe(":8000", logRequest(server.router)))

	//TODO: Better Failure Handline, HTTPS, Etc
	//TODO Auto HTTPS: https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
	//TODO Subdomains: https://translate.google.com/translate?hl=en&sl=am&tl=en&u=http%3A%2F%2Fcodepodu.com%2Fsubdomains-with-golang%2F&anno=2
	//TODO: Centeral Error Handling For All The Routes
	//TODO: CORS, XSS Preventions, Etc
	//TODO: Try Using Defers To Handle Closing The Resources, And Check It Runs Even If There Are Panics
}

// TODO Features:
//		MacOS Update Caching
//		Support For Using Something Like Github Pages As A Cache Of Assets and Maybe Even Files/Software
//		Remote Terminal And Screen Viewer/Controller (ie. SSH & VNC) Using Extra Binary
