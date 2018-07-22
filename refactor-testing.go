package main

import (
  //"fmt"
  "log"

  "github.com/go-pg/pg" // Database (Postgres)
  //"github.com/mattrax/Mattrax/devices"
)
var pgdb *pg.DB

func main() {
  _ = ConnectDB("postgres://oscar.beaumont:@localhost/mattrax?sslmode=disable")

  //var device structs.Device
	//if err := pgdb.Model(&device).Where("uuid = ?", "201abc84d0df045564da0597c7795152a21bd29c").Select(); err != nil { fmt.Println(err); return }
}

func ConnectDB(str string) *pg.DB {
  if options, err := pg.ParseURL(str); err != nil {
		log.Fatal(err)
	} else {
		pgdb = pg.Connect(options)
	}
	if _, err := pgdb.Exec("SELECT 1"); err != nil {
		log.Fatal("Error Communicating With The Database: ", err)
	} //logrus.Fatal
	log.Println("The Database Connected Successfuly")
  return pgdb
}
