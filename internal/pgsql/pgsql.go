package pgsql

import (
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log" //TODO: Use This One Everywhere
)

var DB *pg.DB

func Connect(connURL string) {
	//TODO: Verify the Connection
	if options, err := pg.ParseURL(connURL); err != nil {
		log.Fatal(err)
	} else {
		DB = pg.Connect(options)
	}
	if _, err := DB.Exec("SELECT 1"); err != nil {
		log.Fatal("Error Communicating With The Database: ", err)
	} //logrus.Fatal
	//if !correctSchema() { initDatabaseSchema() }
	log.Info("The Database Connected Successfuly")
}

func Disconnect() {
	//TODO: Test This Is Called
	DB.Close()
}
