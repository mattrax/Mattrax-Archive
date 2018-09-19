package main

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Pattern for endpoint on middleware chain, not takes a diff signature.
func httpHandler(h HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		if err := h(w, r); err != nil {
			//TODO: Log It And Return Request Failure ID For Post Debuggging
			log.Println("An Error In A HTTP Route Happened ", r.URL.String(), err)
			http.Error(w, "An Internal Error Occurred", 500) // TODO: Show Debugging Code For Backend
		}
	})
}

/*
func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if false { //!currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return
*/
