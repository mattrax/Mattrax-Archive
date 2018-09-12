package main

import ( //TODO: Full Paths
	"../../pkg/authentication"
	"../../pkg/enroll"
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

	//s.router.Methods("GET").HandleFunc("/", s.handleIndex())
	//s.router.Methods("GET").HandleFunc("/admin", s.adminOnly(s.handleIndex()))
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
