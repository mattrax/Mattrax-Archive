package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Device struct {
	UDID                  string `db:"udid"`
	Topic                 string `db:"topic"`
	OSVersion             string `db:"os_version"`
	BuildVersion          string `db:"build_version"`
	ProductName           string `db:"product_name"`
	SerialNumber          string `db:"serial_number"`
	IMEI                  string `db:"imei"`
	MEID                  string `db:"meid"`
	PushMagic             string `db:"push_magic"`   //TODO: Byte Array
	UnlockToken           string `db:"unlock_token"` //TODO: Byte Array
	AwaitingConfiguration string `db:"awaiting_configuration"`
}

func main() {
	start := time.Now()

	db, err := sqlx.Connect("postgres", "user=oscar.beaumont dbname=oscar.beaumont sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Connected To Database In %s", time.Since(start))
	queryStart := time.Now()

	devices := []Device{}
	db.Select(&devices, "SELECT * FROM devices") // WARNING THIS LOAD ENTIRE DATABASE SET INTO MEMORY WHICH COULD BE DANGEROUS UNLESS PAGINATED

	log.Println(devices)

	log.Printf("Queried DB In %s", time.Since(queryStart))
	singleQuery := time.Now()

	device := Device{}
	err = db.Get(&device, "SELECT * FROM devices WHERE udid=$1", 1)
	if err != nil {
		panic(err)
	}
	log.Println(device)

	log.Printf("Queried DB In %s", time.Since(singleQuery))
}
