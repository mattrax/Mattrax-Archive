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
  "github.com/mattrax/Mattrax/internal" // Mattrax Internal (Logging, Database and Config)

  //Modules
  //"github.com/mattrax/Mattrax/demoMDM"
)

var (
  config, log, pgdb = internal.GetInternalState()
  srv *http.Server
)

func main() {
  out := config.Get2("domain", "")
  fmt.Println(out)









  //Load The Modules
	//appleMDM.Init()
	//windowsMDM.Init()

	//Webserver Routes



  var listenDomain string
  var err error //Cleanup This
  if listenDomain, err = config.GetString("domain", ""); err != nil {
    log.Fatal("Error Getting Config Parameter 'domain'")
  }

  var listenAddress string
  if listenAddress, err = config.GetString("listen", "0.0.0.0:8000"); err != nil {
    log.Fatal("Error Getting Config Parameter 'listen'")
  }



	router := mux.NewRouter()
	r := router.Host(listenDomain).Subrouter()

  //TODO: Add The Domain To The Startup Webserver Message











	//Webroutes -> Management Interface
	//r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Mattrax MDM Server!") }).Methods("GET")
	//r.HandleFunc("/enroll", enrollmentHandler).Methods("GET")

  r.HandleFunc("/enroll", func (w http.ResponseWriter, r *http.Request) {
  	//TODO: This Will Show Interface To Guide User Through Enrollment
  	http.Redirect(w, r, "/apple/enroll", 301)
  }).Methods("GET")


	//Webroutes -> Modules
	//appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())
	//windowsMDM.Mount(r.PathPrefix("/windows/").Subrouter(), router.Host(config.EEDomain).Subrouter())





	//React Interface
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../MattraxUI/build"))))

	r.HandleFunc("/api/testing", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "{ data: 'Hello World' }") }).Methods("GET")






	//Start The Webserver (In The Background)
	go func() { startWebserver(router, listenAddress) }()
	log.Info("The Mattrax Webserver Is Listening At " + listenAddress) //fmt.Sprintf("%v:%v", "0.0.0.0", config.Port))

	//Wait Until Shutting Down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//Cleanup
	log.Info("Mattrax is Shutting Down...")





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


		var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
		//log.Info(keys)


		handler.ServeHTTP(w, r)
	})
}



// TODO: Comit The Chnages To The Config Library To Github
// TODO Capitalise The Organisations Name
// TODO Go Doc On Functions

//TODO: /apple doesn;t work only /apple/ in browser
