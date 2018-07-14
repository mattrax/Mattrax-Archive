/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Core. It Loads The Database, Logging and The Webserver Components.
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package main

import (
  "fmt"
  "os"
  "io/ioutil"
  //"io"
  "time"
  "context"
	"net/http"
	"os/signal"
  "encoding/json"

  "github.com/sirupsen/logrus" // Logging
  //rotatelogs "github.com/lestrrat-go/file-rotatelogs" // Logging -> Log Rotation
  "github.com/rifflock/lfshook" // Logging -> Console and File Output With Different Formattings
  "github.com/gorilla/handlers" // HTTP Handlers
	"github.com/gorilla/mux" // HTTP Router
  //"github.com/go-pg/pg" // Database (Postgres)






  mdb "github.com/mattrax/mattrax/internal/database" //Mattrax Database



  "github.com/mattrax/mattrax/appleMDM" // The Apple MDM Module
	//"github.com/mattrax/mattrax/windowsMDM" // The Windows MDM Module
)

var pgdb = mdb.Database()


var (
  config = Config{} // The Configuration ('config.json')
  log = logrus.New() // The Logger //TODO: Is this Still needed. not Just have it Blank
  //pgdb *pg.DB // The Database
  srv *http.Server // The Webserver
)

// This Function Loads/Creates The Config, Connects To The Database, Mounts The Webserver Routes And Starts The Webserver. It Also Handles Clean Of These Things On Exit
func main() {
  // Load/Create The Configuration
  if configFile, err := os.Open("config.json"); os.IsNotExist(err) {
    json, err := json.MarshalIndent(newConfig(), "", "  ")
    if err != nil { logrus.Fatal("Error Generating The Config File:", err) }
    if err := ioutil.WriteFile("config.json", json, 0644); err != nil { logrus.Fatal("Error Saving The New Config File To './config.json'") }
    logrus.Warning("A New Config Was Created. Please Populate The Correct Information Before Starting Mattrax Again.")
    return
  } else if err != nil {
    logrus.Fatal("Error Loading The Config File:", err)
  } else {
    if err := json.NewDecoder(configFile).Decode(&config); err != nil { logrus.Fatal("Error Parsing The Config File:", err) }
  }
  //FIXME: Optional Values In The Config

  //Logging (File and Console Output)
  log.Formatter = &logrus.TextFormatter{
      DisableColors: false,
      TimestampFormat : "02/01/06 15:04:05",
      FullTimestamp:true,
  }
	log.Hooks.Add(lfshook.NewHook(
		config.LogFile, //TODO: Append Data/Data+Number For Rolling Log Files Between Restarts
		&logrus.JSONFormatter{},
	))

  //Database
  /*if options, err := pg.ParseURL(config.Database); err != nil { log.Fatal(err) } else {
    pgdb = pg.Connect(options)
  }
  if _, err := pgdb.Exec("SELECT 1"); err != nil { log.Fatal("Error Communicating With The Database: ", err) }
  if !correctSchema() { initDatabaseSchema() }*/

  //Load The Modules
  appleMDM.Init(log) //pgdb
  //windowsMDM.Init()

  //Webserver Routes
  router := mux.NewRouter()
  r := router.Host(config.Domain).Subrouter()
  r.HandleFunc("/", indexHandler).Methods("GET")
  r.HandleFunc("/enroll", enrollmentHandler).Methods("GET")
  appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())

  //Start The Webserver (In The Background)
  go func() { startWebserver(router) }()
  log.Info("The Mattrax Webserver Is Listening At " + fmt.Sprintf("%v:%v", "0.0.0.0", config.Port))

  //Wait Untill Shutting Down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

  //Cleanup
  log.Info("Mattrax is Shutting Down...")
  pgdb.Close() //Shutdown The Database
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*15); defer cancel(); srv.Shutdown(ctx) // Shutdown The Webserver
	os.Exit(0) //Exit Successfully
}

/* Local Configuration */
func newConfig() Config {
  return Config{
    Name: "Acme Inc",
    Domain: "mdm.acme.com",
    Verbose: false,
    LogFile: "data/log.txt",
    Port: 8000,
    Database: "postgres://postgres:@postgres/postgres",
  }
}

type Config struct {
  Name string `json:"name"`
  Domain string `json:"domain"`
  Verbose bool `json:"verbose"`
  LogFile string `json:"logFile"`
  Port int `json:"port"`
  Database string `json:"database"`
}

/* Database Initialisation */ //FIXME: Make This Entire Section Work
func correctSchema() bool {
  if _, err := pgdb.Exec("SELECT * FROM devices"); err != nil {
    //Find Out If Error Was Database Table If Not Log Fatal
    return false
  }
  return true
}

func initDatabaseSchema() {
  panic("This will create the database schema")
}

/* The Webserver */
func startWebserver(router *mux.Router) {
  var handler http.Handler
  if config.Verbose {
    handler = verboseRequestLogger(handlers.CORS()(router))
  } else {
    handler = handlers.CORS()(router)
  }

	srv = &http.Server{
		Addr:         fmt.Sprintf("%v:%v", "0.0.0.0", config.Port), //FIXME Configurable Listen IP (Optional)
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	if err := srv.ListenAndServe(); err != nil { log.Fatal(err) }
}

func verboseRequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Debug("Request: " + r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		handler.ServeHTTP(w, r)
	})
}

/* HTTP Handling Routes */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mattrax MDM Server!")
}

func enrollmentHandler(w http.ResponseWriter, r *http.Request) {
  //TODO: This Will Show Interface To Guide User Through Enrollment
  http.Redirect(w, r, "/apple/enroll", 301)
}


/* The End */
//Now
// FUTURE FEATURE: Redo Separator Between Blocks Of Function -> They Don't Stand Out Enought
// TODO: Contant Pinging Database To Stop HTTP Soon As It Stops Connecting

// Log Wiping After Restart (Fix That)
//      Redo Logging For Subfiles To Use The Features Of The New System
// TODO: Log File Roation So The Log Files Doesn't Get To Big
//      Config Options For External Logging Server
// TODO: Godoc Documentation Throughout Code
// TODO: GoDEP Package Management
// TODO: Check Line Ending for ; and Remove (Maybe Add Test To Check For Them)

//Far In The Future
// IDEA: HTTPS And HTTP Support With Automatic Redirection Between
//      Certbot ACME Built In For Automaticly Issuing And Renewing Cert
// IDEA: Built In IP Whitelist For Access To The Admin Area
