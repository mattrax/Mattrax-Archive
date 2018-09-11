package main

import (  //TODO: Full Paths
  "../../pkg/vue"
  "../../pkg/enroll"
)

func (s *Server) routes() { //TODO: Parsing Server To All These Endpoints
  // Vue Routes
  s.router.HandleFunc("/", vue.IndexHandler())

  // User Endpoints
  s.router.HandleFunc("/enroll", enroll.EnrollHandler())

  // MDM Endpoints





	//s.router.HandleFunc("/", s.handleIndex())
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleIndex()))
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
