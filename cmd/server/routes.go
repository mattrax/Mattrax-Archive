package main

import ( //TODO: Full Paths
	"log"
	"net/http"

	"../../pkg/apple/authentication"
	"../../pkg/apple/enroll"
	"../../pkg/vue"
)

func (s *Server) routes() { //TODO: Parsing Server To All These Endpoints
	// Vue Interface
	s.router.HandleFunc("/", vue.IndexHandler()).Methods("GET")

	// Special Interfaces
	s.router.HandleFunc("/enroll", enroll.EnrollHandler()).Methods("GET")

	// API Endpoints
	s.router.HandleFunc("/api/login", authentication.LoginHandler()).Methods("GET")

	// MDM Endpoints
	//  /MDMServiceConfig
	//  /apple/checkin
	//  /apple/server

	// TEMP
	//s.router.Methods("GET").HandleFunc("/", s.handleIndex())
	//s.router.Methods("GET").HandleFunc("/admin", s.adminOnly(s.handleIndex()))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// TEMP //
/*
func (s *server) handleIndex() http.HandlerFunc {
	//Defines Structs/Static Stuff Here
	msg := "Hello World"
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, msg)
	}
}

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if false { //!currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}*/
