package internal

import (
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

func GetInternalState() (gonfig.Gonfig, *logrus.Logger, *pg.DB) {
  return LoadConfig(), LoadLogging(), LoadDatabase()
}

func CleanInternalState() {
  pgdb.Close()
}

func LoadConfig() gonfig.Gonfig {
  var err error
	config, err = gonfig.FromJsonFile("config.json") //TODO FIXME Generate Config If Not Found
	if err != nil {
		log.Fatal(err)
	}
  //Check Default Params
  if _, err = config.GetString("name", nil); err != nil {
    log.Fatal("Missing The Required Config Parameter 'name'")
  }

  return config
}

type Fields = logrus.Fields  // Export Logrus Fields (So It Does Have To Be Imported By Another Package)
func LoadLogging() *logrus.Logger {
  if verboseMode, err := config.GetBool("verbose", false); err == nil {
    if verboseMode {
      log.Info("Enabled Verbose Logging")
      log.Level = logrus.DebugLevel
    }
  } //TODO Error Handling

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
  return log
}

func LoadDatabase() *pg.DB {
  if DatabaseURL, err := config.GetString("database", "postgresql://localhost/mattrax"); err != nil {
    log.Fatal(err)
  } else {
    if options, err := pg.ParseURL(DatabaseURL); err != nil { //FIXME
  		log.Fatal(err)
  	} else {
  		pgdb = pg.Connect(options)
  	}
  	if _, err := pgdb.Exec("SELECT 1"); err != nil {
      log.Warning("Connecting To The Server: " + DatabaseURL)
  		log.Fatal("Error Communicating With The Database: ", err)
  	}
  	//if !correctSchema() { initDatabaseSchema() } //TODO Add Schema Generation
  	log.Info("Connected To The Database Successfuly")
    return pgdb
  }
  return nil
}

// TODO: Go Doc Every Function
