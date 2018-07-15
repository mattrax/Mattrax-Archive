package database

import (
	// External Deps
	"github.com/go-pg/pg" // Database (Postgres)

	// Internal Functions
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mlg "github.com/mattrax/Mattrax/internal/logging"       //Mattrax Logging
)

var log = mlg.GetLogger()
var config = mcf.GetConfig() // Get The Internal State
var pgdb *pg.DB              // The Database

//TODO: Go Doc
func init() {
	if options, err := pg.ParseURL(config.Database); err != nil {
		log.Fatal(err)
	} else {
		pgdb = pg.Connect(options)
	}
	if _, err := pgdb.Exec("SELECT 1"); err != nil {
		log.Fatal("Error Communicating With The Database: ", err)
	} //logrus.Fatal
	//if !correctSchema() { initDatabaseSchema() }
	log.Info("The Database Connected Successfuly")
}

// TODO: Go Doc
func GetDatabase() *pg.DB { return pgdb }

// TODO: Go Doc
func Cleanup() { pgdb.Close() }
