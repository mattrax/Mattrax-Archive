package internal

import (
  "os"
  "io/ioutil"
  "encoding/json"

  "github.com/go-pg/pg"
  "github.com/sirupsen/logrus"
  "github.com/rifflock/lfshook"
  "github.com/oscartbeaumont/gonfig"
)

var (
  pgdb *pg.DB
  config gonfig.Gonfig
  log = logrus.New()
)

func init() {
  LoadConfig()
  LoadLogging()
  LoadDatabase()
}

func GetInternalState() (gonfig.Gonfig, *logrus.Logger, *pg.DB) {
  return config, log, pgdb
}

func CleanInternalState() {
  pgdb.Close()
}

/* Configuration */

type DefaultConfig struct {
	Name string `json:"name"`
	Domain string `json:"domain"`
  Database string `json:"database"`
}

func LoadConfig() gonfig.Gonfig {
  if configFile, err := os.Open("config.json"); err == nil {
    if _config, err := gonfig.FromJson(configFile); err == nil {
      config = _config
    } else {
      log.Fatal("Error in The Config: ", err)
    }
    //Check Default Params
    if _, err = config.GetString("name", nil); err != nil {
      log.Fatal("Missing The Required Config Parameter 'name'")
    }
    if _, err = config.GetString("domain", nil); err != nil {
      log.Fatal("Missing The Required Config Parameter 'domain'")
    }
    if _, err = config.GetString("database", nil); err != nil {
      log.Fatal("Missing The Required Config Parameter 'database'")
    }

    return config
  } else if os.IsNotExist(err) {
    defaultConfig := DefaultConfig{
  		Name: "Acme Inc",
  		Domain: "mdm.acme.com",
  		Database: "postgres://postgres:@postgres/postgres",
  	}
		if json, err := json.MarshalIndent(defaultConfig, "", "  "); err == nil {
      if err := ioutil.WriteFile("config.json", json, 0644); err != nil {
  			log.Fatal("Error Saving The New Config File To './config.json'")
  		}
  		log.Println("A New Config Was Created. Please Populate The Correct Information Before Starting Mattrax Again.")
  		os.Exit(0)
		} else {
      log.Fatal("Error Generating The Config File:", err)
    }
  } else {
    log.Fatal("Error Loading The Config File:", err)
  }
  return nil
}

/* Logging */

type Fields = logrus.Fields  // Export Logrus Fields (So It Does Have To Be Imported By Another Package)
func LoadLogging() *logrus.Logger {
  log.Formatter = &logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "02/01/06 15:04:05",
		FullTimestamp:   true,
	}

  if logFile, err := config.GetString("logFile", "log.json"); err != nil {
    log.Fatal("Missing The Required Config Parameter 'name'")
  } else {
    if logFile != "" {
      log.Hooks.Add(lfshook.NewHook(
    		logFile, //TODO: Append Data/Data+Number For Rolling Log Files Between Restarts
    		&logrus.JSONFormatter{},
    	))
    }
  }

  if verboseMode, err := config.GetBool("verbose", false); err == nil {
    if verboseMode {
      log.Debug("Enabled Verbose Logging")
      log.Level = logrus.DebugLevel
    }
  }

  return log
}

/* Database */

func LoadDatabase() *pg.DB {
  databaseURL := config.JustGetString("database", "")

  if options, err := pg.ParseURL(databaseURL); err != nil { //FIXME
    log.Fatal("Failed To Parse The Database Connection URL: '", databaseURL, "'")
    return nil
  } else {
    pgdb = pg.Connect(options)
  }
  if _, err := pgdb.Exec("SELECT 1"); err != nil {
    log.Fatal("Error ", err, " Communicating With The Database: '", databaseURL, "'")
    return nil
  }
  //if !correctSchema() { initDatabaseSchema() } //TODO Add Schema Generation
  log.Debug("Connected To: '", databaseURL, "'")
  return pgdb
}
