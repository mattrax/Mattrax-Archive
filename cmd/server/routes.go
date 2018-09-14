package main

import ( //TODO: Full Paths
	"log"
	"net/http"

	"../../pkg/apple/handlers/checkin"
	"../../pkg/apple/handlers/enroll"
	"../../pkg/apple/handlers/server"
	"../../pkg/authentication"
	"../../pkg/vue"
)

func (s *Server) routes() { //TODO: Parsing Server To All These Endpoints
	// Vue Interface
	s.router.HandleFunc("/", vue.IndexHandler()).Methods("GET")
	// /enroll

	// API Endpoints
	s.router.HandleFunc("/api/login", authentication.LoginHandler()).Methods("GET")

	// MDM Endpoints
	s.router.HandleFunc("/apple/enroll", enroll.Handler()).Methods("GET")
	s.router.HandleFunc("/MDMServiceConfig", enroll.MDMServiceConfigHandler()).Methods("GET")
	s.router.HandleFunc("/apple/checkin", checkin.Handler(s)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	s.router.HandleFunc("/apple/server", server.Handler()).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")

	// TEMP
	s.router.NotFoundHandler = checkin.Handler(s)

	s.router.HandleFunc("/enroll", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-apple-aspen-config")
		http.ServeFile(w, r, "enroll.mobileconfig")
	}).Methods("GET")
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
