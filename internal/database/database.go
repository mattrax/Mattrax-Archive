package database

import (
  "fmt"
  "log" // Chnage TO Main Logger

  "github.com/go-pg/pg" // Database (Postgres)
)

var pgdb *pg.DB

//TODO: Go Doc
func init() { //Should Only Run Once
  fmt.Println("Ran Database Load")



  if options, err := pg.ParseURL("postgres://oscar.beaumont:@localhost/mattrax?sslmode=disable"); err != nil { log.Fatal(err) } else { //config.Database
    pgdb = pg.Connect(options)
  }
  if _, err := pgdb.Exec("SELECT 1"); err != nil { log.Fatal("Error Communicating With The Database: ", err) } //logrus.Fatal
  //if !correctSchema() { initDatabaseSchema() }
}

// TODO: Go Doc
func GetDatabase() *pg.DB { return pgdb }


//Clenaup Handling
