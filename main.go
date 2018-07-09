package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"log"

	_ "github.com/boltdb/bolt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"./appleMDM"
	"./windows10"
)

//func log(input string) { fmt.Println("" + input) }
/*
type TestingDB struct {
  title string
  author string
}

func init() {
  log.Println("Init")

  db, err := bolt.Open("my.db", 0644, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

  os.Exit(0)
}*/

func main() {
	log.Println("Starting Mattrax Daemon Version: " + "V0.0.1") //TODO: Load Version From Somewhere

	appleMDM.Init()
	windows10.Init()
	webserver()
}

func webserver() {
	router := mux.NewRouter()
	r := router.Host("mdm.otbeaumont.me").Subrouter() //.Schemes("http")

	r.HandleFunc("/", indexHandler).Methods("GET")

	appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())
	windows10.Mount(r.PathPrefix("/windows10/").Subrouter(), router.Host("enterpriseenrollment.otbeaumont.me").Subrouter())

	authServices := router.Host("auth.otbeaumont.me").Subrouter()
	authServices.HandleFunc("/", authHandler).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      logRequest(handlers.CORS()(router)),
	}

	// Run in a Goroutine So That It Doesn't Block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			//fatalLog(err) //// TODO: Fix This Error Parsing The Error In (Cause Its Not A String)
			fmt.Println(err) //// ^
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Wait Until Shutdown Signal Received.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Tried To Shutdown The Server If Fails To Do It Before The Deadline Kill It
	srv.Shutdown(ctx)
	log.Println("Shutting Down...")
	os.Exit(0)
}

func logRequest(handler http.Handler) http.Handler { ///////// This is TEMP Code That Will Be Removed
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// right not all this does is log like
		// "github.com/gorilla/handlers"
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		// However since this is middleware you can have it do other things
		// Examples, auth users, write to file, redirects, handle panics, ect
		// add code to log to statds, remove log.Printf if you want

		handler.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Authentication Service!")
}
