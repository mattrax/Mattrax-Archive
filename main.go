/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Core. It Loads The Database, Logging and The Webserver Components.
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package main

import (
  "fmt" //// TODO: Eliminate This (Only Used Once)

  "context"
	"net/http"
	"os/signal"
	"time"

  "os"
  "flag"

  "github.com/op/go-logging" // Logging
  "github.com/gorilla/handlers" // HTTP Handlers
	"github.com/gorilla/mux" // HTTP Router
  "github.com/go-pg/pg" // Database (Postgres)

  "./appleMDM" //TODO: Should These Be Put In Main() After Logging & Database Are Ready???
	//"./windowsMDM"
)

/* Internal Global Varibles */
var (
  log = logging.MustGetLogger("main") // Logger
  pgdb *pg.DB // The Daatabse
  srv *http.Server // The Webserver
)

/* Command Line Arguments */
var (
  logFile     = flag.String("log-file", "log.txt", "The File For Log To Be Stored To. Blank To Disable")
  verbose     = flag.Bool("verbose", false, "Log With More Detail. To Be Used For Debugging Only")

  address     = "0.0.0.0:8000" //TODO: Make This Alot Less Hardcoded
  //TODO: Database Creds From Maybe .env file or command line in
)

func main() {
  flag.Parse() // Parse The Command Line Arguments
  setupLogger() // Setup The Logger
  log.Info("Started The Mattrax Daemon...")
  loadDatabase()

  appleMDM.Init(pgdb, log)
  //windowsMDM.Init()

  router := mux.NewRouter()
	r := router.Host("mdm.otbeaumont.me").Subrouter() //.Schemes("http")
	r.HandleFunc("/", indexHandler).Methods("GET")
	appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())
	//windowsMDM.Mount(r.PathPrefix("/windows10/").Subrouter(), router.Host("enterpriseenrollment.otbeaumont.me").Subrouter())
	authServices := router.Host("auth.otbeaumont.me").Subrouter()
	authServices.HandleFunc("/", authHandler).Methods("GET")

  startWebserver(router, address)

  log.Info("The Mattrax Webserver Is Listening At " + address)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Wait Until Shutdown Signal Received.

  log.Warning("Shutting Down...")
  shutdownDatabase()
  shutdownWebserver()
  log.Warning("Shutdown Complete")
	os.Exit(0)
}





func setupLogger() {
  format := logging.MustStringFormatter(`%{time:15:04:05} %{color}[%{level}]%{color:reset} %{message}`)
  consoleBackend := logging.AddModuleLevel(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))
  var fileBackend logging.LeveledBackend

  if *logFile != "" {
    logFile, loadLogErr := os.OpenFile(*logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
    if loadLogErr != nil {
      fmt.Println(loadLogErr)
      fmt.Println("Failed To Load or Create The Log File. A New Path Can Be Parsed In By Using --log-file")
      os.Exit(2)
    }

    _, extraLinesError1 := logFile.Write([]byte("----------------------------------\n"))
    _, extraLinesError2 := logFile.Write([]byte("-------- Starting Mattrax --------\n")) //TODO Make These Lines Longer
    _, extraLinesError3 := logFile.Write([]byte("----------------------------------\n"))
    if extraLinesError1 != nil || extraLinesError2 != nil || extraLinesError3 != nil { //TODO Better Error Handling INstead Of So Many If Statements
      fmt.Println(extraLinesError1)
      fmt.Println(extraLinesError2)
      fmt.Println(extraLinesError3)
      fmt.Println("Failed To Add Separator To The Logging File") //TODO: Reword This
      os.Exit(2)
    }


    fileBackend = logging.AddModuleLevel(logging.NewBackendFormatter(logging.NewLogBackend(logFile, "", 0), format))
    logging.SetBackend(consoleBackend, fileBackend)
  } else {
    logging.SetBackend(consoleBackend)
  }

  // Set The Logging Level Based on The "verbose" Command Line Flag
  if *verbose {
    consoleBackend.SetLevel(logging.DEBUG, "")
    if *logFile != "" { fileBackend.SetLevel(logging.DEBUG, "") }
    log.Notice("Verbose Logging Is Enabled")
  } else {
    consoleBackend.SetLevel(logging.INFO, "")
    if *logFile != "" { fileBackend.SetLevel(logging.INFO, "") }
  }
}

func loadDatabase() {
  //Initialise Schema Here For New DB's
  // if database scheme doesn't exist run initDatabaseSchema()

  pgdb = pg.Connect(&pg.Options{ //https://github.com/go-pg/pg/issues/188
    User: "oscar.beaumont",
    Database: "mattrax",
    //Host:     postgres.Host,
    //Port:     postgres.Port,
  })
}
func shutdownDatabase() {
  pgdb.Close()
  log.Warning("Database Connection Closed")
}

func initDatabaseSchema() {
  panic("This will create the database schema")

}














func startWebserver(router *mux.Router, address string) {
  var handler http.Handler
  if *verbose {
    handler = logRequest(handlers.CORS()(router))
  } else {
    handler = handlers.CORS()(router)
  }

	srv = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	// Run in a Goroutine So That It Doesn't Block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			//fatalLog(err) //// TODO: Fix This Error Parsing The Error In (Cause Its Not A String)
			fmt.Println(err) //// ^
		}
	}()
}

func shutdownWebserver() {
  log.Warning("Wait For Webserver Request To All Be Complete...")
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Tried To Shutdown The Server If Fails To Do It Before The Deadline Kill It
	srv.Shutdown(ctx)
  log.Warning("All Webserver Request Complete. The Webserver Has Been Shutdown.")
}


func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Debug("Request: " + r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		handler.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Authentication Service!")
}


/*log.Debug("debug")
log.Info("info")
log.Notice("notice")
log.Warning("warning")
log.Error("err")*/
//log.Critical("crit")
/*
r.HandleFunc("/test", testingHandler).Methods("GET")

type Device struct {
  TableName struct{} `sql:"devices"`
  SerialNumber  string `sql:"SerialNumber,pk"`
  ProductName   string `sql:"ProductName"`
  OSVersion     string `sql:"OSVersion"`
  Topic         string `sql:"Topic"`
  UDID          string `sql:"UDID"`
  Token         string `sql:"Token"`
  PushMagic     string `sql:"PushMagic"`
  UnlockToken   string `sql:"UnlockToken"`
}

func testingHandler(w http.ResponseWriter, r *http.Request) {
  var devices []Device
  err := pgdb.Model(&devices).Select()
  if err != nil {
		log.Error(err)
		return
	}

  fmt.Println(devices)
  fmt.Println(devices[0].SerialNumber)

	fmt.Fprintf(w, "Hello World!")
}
*/

// TODO: Contant Pining Database To Stop HTTP Soon As It Stops Connecting
// TODO: Retrys For The Database If it Fails The First Time
// TODO: Better Lookin Error Handling (The If Statements Everywhere)
// TODO: File Roation So The Log Files Doesn't Get To Bit
// TODO: Support For External Logging Server
// TODO: HTTPS And HTTP Support With Automatic Redirection Between
// TODO: Firewalling IP Allowed To Access Admin Area
// TODO: Postgress Database Workers Queue For Better Preformance

//var format = logging.MustStringFormatter(`%{color}[%{level}]%{color:reset} %{time:15:04:05} ▶ %{message}`) // :-7s //%{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}
