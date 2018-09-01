package pgsql

import (
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log" //TODO: Use This One Everywhere
)

var pgdb *pg.DB

func init() { //TODO: Is It Good This Creates A New Connection For Each File It Is Loaded With
	connURL := "postgres://oscar.beaumont:@localhost/mattrax_new?sslmode=disable" //TODO: Load From Config

	//TODO: Verify the Connection
	if options, err := pg.ParseURL(connURL); err != nil {
		log.Fatal(err)
	} else {
		pgdb = pg.Connect(options)
	}
	if _, err := pgdb.Exec("SELECT 1"); err != nil {
		log.Fatal("Error Communicating With The Database: ", err)
	} //logrus.Fatal
	//if !correctSchema() { initDatabaseSchema() }
	log.Info("The Database Connected Successfully")
}

func GetDB() *pg.DB {
	return pgdb
}

func NotFound(err error) bool {
	if err == pg.ErrNoRows || err == pg.ErrMultiRows {
		return false
	} else {
		return true
	}
}

func Disconnect() {
	//TODO: Test This Is Called
	pgdb.Close()
}
