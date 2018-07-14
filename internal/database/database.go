package database

import (
  // External Deps
  "github.com/go-pg/pg" // Database (Postgres)

  // Internal Functions
  mlg "github.com/mattrax/mattrax/internal/logging" //Mattrax Logging
  mcf "github.com/mattrax/mattrax/internal/configuration" //Mattrax Configuration
)

var log = mlg.GetLogger(); var config = mcf.GetConfig() // Get The Internal State
var pgdb *pg.DB // The Database

//TODO: Go Doc
func init() {
  if options, err := pg.ParseURL(config.Database); err != nil { log.Fatal(err) } else {
    pgdb = pg.Connect(options)
  }
  if _, err := pgdb.Exec("SELECT 1"); err != nil { log.Fatal("Error Communicating With The Database: ", err) } //logrus.Fatal
  //if !correctSchema() { initDatabaseSchema() }
  log.Info("The Database Connected Successfully")
}

// TODO: Go Doc
func GetDatabase() *pg.DB { return pgdb }

// TODO: Go Doc
func Cleanup() { pgdb.Close() }
