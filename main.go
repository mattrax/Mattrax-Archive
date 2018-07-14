/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Core. It Loads The Database, Logging and The Webserver Components.
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package main

import (
  "fmt"

  "context"
	"net/http"
	"os/signal"
	"time"

  "os"

  "encoding/json"
  "io"

  "github.com/Sirupsen/logrus" // Logging
  "github.com/gorilla/handlers" // HTTP Handlers
	"github.com/gorilla/mux" // HTTP Router
  "github.com/go-pg/pg" // Database (Postgres)

  //TODO: Change All Of These Imports From Dyanmic Paths
  "github.com/mattrax/mattrax/appleMDM" //TODO: Should These Be Put In Main() After Logging & Database Are Ready???
	//"github.com/mattrax/mattrax/windowsMDM"
)

var (
  config = Config{}
  log = logrus.New()
  pgdb *pg.DB // The Database
  srv *http.Server // The Webserver
)

type Config struct {
  Name string `json:"name"`
  Verbose bool `json:"verbose"`
  LogFile string `json:"logFile"`
  Port int `json:"port"`
  Database struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Host string `json:"host"`
    Port string `json:"port"`
  } `json:"database"`
}

func main() {
  //flag.Parse() // Parse The Command Line Arguments //TODO: Remove It

  //TODO: Handle Error By Quitting During This Func -> INcluding config Errors

  // Load The Configuration
  if configFile, err := os.Open("config.json"); err != nil { logrus.Fatal("Error Loading The Config File:", err) } else {
    if err := json.NewDecoder(configFile).Decode(&config); err != nil { logrus.Fatal("Error Parsing The Config File:", err) }
  }


  // TODO: Create Default Config If Missing

  /*defer file.Close()
  decoder := json.NewDecoder(file)
  configuration := Configuration{}
  err := decoder.Decode(&configuration)
  if err != nil {
    fmt.Println("error:", err)
  }
  fmt.Println(configuration.Users)*/






  //Logging (File and Console Output)
  //log = logrus.New()
  var logFile io.Writer
  var logError error //TODO: Attempt To Remove This
  if logFile, logError = os.OpenFile(config.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); logError != nil { log.Fatal("Error Loading Log File", logError) }

  //logFile, _ := os.OpenFile(*logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)   //TODO: Close The Logging File On Shutdown
  //TODO: Optional Values In The Config




  log.SetOutput(io.MultiWriter(logFile, os.Stdout))
  if config.Verbose { log.Level = logrus.DebugLevel }
  log.Info("Started The Mattrax Daemon...")

  //Database
  pgdb = pg.Connect(&pg.Options{
    User:     "oscar.beaumont", //TODO: Load These Settings From Somewhere
    //Password: "",
    Database: "mattrax",
    //Host:     postgres.Host,
    //Port:     postgres.Port,
  })

  pgdb = pgdb.WithTimeout(30 * time.Second) //TODO: Check This Works
  // TODO: Initialise Schema Here For New DB's
  // TODO: if database scheme doesn't exist run initDatabaseSchema()

  //Load The Modules
  appleMDM.Init(pgdb, log)
  //windowsMDM.Init()

  //Webserver Routes
  router := mux.NewRouter()


  /* Redo Rouets */
	r := router.Host("mdm.otbeaumont.me").Subrouter() //.Schemes("http") //TODO: Get Hostnames From Config
	r.HandleFunc("/", indexHandler).Methods("GET")
	appleMDM.Mount(r.PathPrefix("/apple/").Subrouter())
	//windowsMDM.Mount(r.PathPrefix("/windows10/").Subrouter(), router.Host("enterpriseenrollment.otbeaumont.me").Subrouter())
	authServices := router.Host("auth.otbeaumont.me").Subrouter()
	authServices.HandleFunc("/", authHandler).Methods("GET")
  /* End Redo Routes */

  //Start The Webserver
  go func() { startWebserver(router) }() //log.Info("The Mattrax Webserver Is Listening At " + address)

  //Wait For Shutdown Interupt
  //logFile.Close()



	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Wait Until Shutdown Signal Received.

  log.Warning("Shutting Down...")
  pgdb.Close()
  shutdownWebserver()
  log.Warning("Shutdown Complete")
	os.Exit(0)
}



func initDatabaseSchema() {
  panic("This will create the database schema")

}














func startWebserver(router *mux.Router) {
  var handler http.Handler
  if config.Verbose {
    handler = logRequest(handlers.CORS()(router))
  } else {
    handler = handlers.CORS()(router)
  }

	srv = &http.Server{
		Addr:         fmt.Sprintf("%v:%v", "0.0.0.0", config.Port), //TODO Configurable Listen IP
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	// Run in a Goroutine So That It Doesn't Block.
	//go func() {
	if err := srv.ListenAndServe(); err != nil {
		//fatalLog(err) //// TODO: Fix This Error Parsing The Error In (Cause Its Not A String)
		fmt.Println(err) //// ^
	}
	//}()
}

func shutdownWebserver() {
  //log.Warning("Wait For Webserver Request To All Be Complete...")
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Tried To Shutdown The Server If Fails To Do It Before The Deadline Kill It
	srv.Shutdown(ctx)
  //log.Warning("All Webserver Request Complete. The Webserver Has Been Shutdown.")
}


func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Debug("Request: " + r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		handler.ServeHTTP(w, r)
	})
}







/* Placeholder Routes */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Authentication Service!")
}











/*
func shutdownDatabase() {
  pgdb.Close()
  //log.Warning("Database Connection Closed")
}
*/

/*
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
*/

/*
func setupLogger() {
  var format logging.Formatter
  if *verbose {
    format = logging.MustStringFormatter(`%{time:15:04:05} %{color}[%{level}]%{color:reset} [%{shortfile}] %{message}`)
  } else {
    format = logging.MustStringFormatter(`%{time:15:04:05} %{color}[%{level}]%{color:reset} %{message}`)
  }

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
*/

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

/*
func init() {
  logging2.Testing()
  os.Exit(1)
}
*/

// TODO: Get The Database Config From Somewhere
// TODO: Redo Start and Shutdown Logging Messages
// TODO: Contant Pining Database To Stop HTTP Soon As It Stops Connecting
// TODO: Retrys For The Database If it Fails The First Time
// TODO: Better Lookin Error Handling (The If Statements Everywhere)
// TODO: File Roation So The Log Files Doesn't Get To Bit
// TODO: Support For External Logging Server
// TODO: HTTPS And HTTP Support With Automatic Redirection Between
// TODO: Firewalling IP Allowed To Access Admin Area
// TODO: Postgress Database Workers Queue For Better Preformance
// TODO: Initilise Database Schema

//var format = logging.MustStringFormatter(`%{color}[%{level}]%{color:reset} %{time:15:04:05} ▶ %{message}`) // :-7s //%{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}
