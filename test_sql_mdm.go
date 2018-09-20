package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

type Device struct {
	Id   string `db:"id"`
	UDID string `db:"UDID"`

	Topic                 string `db:"Topic"`
	OSVersion             string `db:"OSVersion"`
	BuildVersion          string `db:"BuildVersion"`
	ProductName           string `db:"ProductName"`
	SerialNumber          string `db:"SerialNumber"`
	IMEI                  string `db:"IMEI"`
	MEID                  string `db:"MEID"`
	PushMagic             string `db:"PushMagic"`   //TODO: Byte Array
	UnlockToken           string `db:"UnlockToken"` //TODO: Byte Array
	AwaitingConfiguration string `db:"AwaitingConfiguration"`
}

func main() {
	start := time.Now()

	connStr := "postgres://oscar.beaumont:@localhost/oscar.beaumont?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//db.Ping()

	log.Printf("Connected To Database In %s", time.Since(start))
	queryStart := time.Now()

	var devices []Device
	rows, err := db.Query("SELECT * FROM devices") // WHERE id = $1", 1)
	if err != nil {
		log.Fatal("Query: ", err)
	}
	_, err = dbr.Load(rows, &devices)

	log.Println(devices)

	log.Printf("Queried DB In %s", time.Since(queryStart))
	singleQuery := time.Now()

	var row Device
	err = db.QueryRow("SELECT * FROM devices WHERE id = $1", 5).Scan(&row.Id, &row.UDID, &row.Topic, &row.OSVersion, &row.BuildVersion, &row.ProductName, &row.SerialNumber, &row.IMEI, &row.MEID, &row.PushMagic, &row.UnlockToken, &row.AwaitingConfiguration) // SELECT|devices|id,UDID|age=?
	if err != nil {
		log.Fatal("Query: ", err)
	}

	log.Println(row)

	log.Printf("Queried DB In %s", time.Since(singleQuery))

	/*
	 */

	/*rows, err := db.Query("SELECT|devices|id,UDID,Topic|")
	if err != nil {
		t.Fatalf("Query: %v", err)
	}

	log.Println(rows)*/
}
