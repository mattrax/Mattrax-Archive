package main

import (
  "context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/mattrax/Mattrax/internal"

  //Modules
  "github.com/mattrax/Mattrax/appleMDM"
)

var (
  config, log, pgdb = internal.GetInternalState()
  srv *http.Server
)

func main() {
	//Webserver Routes
	router := mux.NewRouter()
	r := router.Host(config.JustGetString("domain", "")).Subrouter()

	//Webroutes -> Management Interface

	//r.HandleFunc("/enroll", enrollmentHandler).Methods("GET")

  r.HandleFunc("/enroll", func (w http.ResponseWriter, r *http.Request) {
  	//TODO: This Will Show Interface To Guide User Through Enrollment
  	http.Redirect(w, r, "/apple/enroll", 301)
  }).Methods("GET")


	//Webroutes -> Modules
	appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())
	//windowsMDM.Mount(r.PathPrefix("/windows/").Subrouter(), router.Host(config.EEDomain).Subrouter())



  //TODO: User Authentcation
  //TODO: Show Mattrax Admin Interface + API/Login If IP Range Is Good

	//React Interface
  /*MattraxUI := http.Dir("../MattraxUI/build")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(MattraxUI)))
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(MattraxUI))))
  r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(MattraxUI))))
  */
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Mattrax MDM Server!")
  }).Methods("GET")





  //r.HandleFunc("/api/testing", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "{ data: 'Hello World' }") }).Methods("GET")






	//Start The Webserver (In The Background)
	go func() { startWebserver(router, config.JustGetString("listen", "0.0.0.0:8000")) }()
	log.Info("Listening For: '", config.JustGetString("domain", ""), "' At '", config.JustGetString("listen", "0.0.0.0:8000"), "'")

	//Wait Until Shutting Down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//Cleanup
	log.Info("Mattrax is Shutting Down...")




  internal.CleanInternalState()
	//mdb.Cleanup() //Shutdown The Database






	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx) // Shutdown The Webserver
	os.Exit(0)        //Exit Successfuly


  //internal.CleanInternalState() //Run On Exit
}




//TODO Docs
func startWebserver(router *mux.Router, addr string) {
	var handler http.Handler



	/*if config.Verbose { //FIXME
		handler = verboseRequestLogger(handlers.CORS()(router))
	} else {
		handler = handlers.CORS()(router)
	}*/
  handler = verboseRequestLogger(handlers.CORS()(router))




	srv = &http.Server{
		Addr:         addr, //FIXME Configurable Listen IP (Optional)
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
		handler.ServeHTTP(w, r)
	})
}


// TODO: Go Doc Every Function & File
// TODO: Check log.Fatal Does Cleanup Funcs (Safly Kills DB And Webserver)
// TODO: Comit The Chnages To The Config Library To Github
// TODO Capitalise The Organisations Name
// TODO Go Doc On Functions

//TODO: /apple doesn;t work only /apple/ in browser
