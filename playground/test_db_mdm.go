package main

import (
	"log"
	"time"

	"upper.io/db.v3/postgresql"
)

// Connection settings.
var settings = postgresql.ConnectionURL{
	Host:     "127.0.0.1",
	Database: "oscar.beaumont",
	User:     "oscar.beaumont",
	Password: "",
}

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

	sess, err := postgresql.Open(settings) // Open a connection.
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// Set this to true to enable the query logger which will print all SQL
	// statements to stdout.
	//sess.SetLogging(true)

	log.Printf("Connected To Database In %s", time.Since(start))
	queryStart := time.Now()

	devicesTable := sess.Collection("devices")

	var devices []Device
	err = devicesTable.Find().All(&devices)
	if err != nil {
		panic(err)
	}
	log.Println(devices)

	log.Printf("Queried DB In %s", time.Since(queryStart))
	singleQuery := time.Now()

	devicesTable = sess.Collection("devices")

	var device Device
	err = devicesTable.Find("id", 5).One(&device)
	if err != nil {
		panic(err)
	}
	log.Println(device)

	log.Printf("Queried DB In %s", time.Since(singleQuery))
}
