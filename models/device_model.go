package models

import (
	"io"
	"log"

	"github.com/groob/plist"
	"github.com/jmoiron/sqlx"
)

// A Device (Maps Between Checkin Device Request And The Database)
//	WARNING: Any Changes To The DB Schema Need To Be Updated On This Struct And The Structs Method "UpdateDB"
type Device struct {
	Status int `db:"status"`
	AppleAuthenticateDetails

	//DeviceRequest `db:"-"`
	/*
		  DeviceName   string `plist:"DeviceName,omitempty"` //TODO: Do I Need These/What Devices Send It
			Challenge    []byte `plist:"Challenge,omitempty"`  //TODO: Do I Need These/What Devices Send It
			Model        string `plist:"Model,omitpempty"`     //TODO: Do I Need These/What Devices Send It
			ModelName    string `plist:"ModelName,omitempty"`  //TODO: Do I Need These/What Devices Send It
	*/
}

type AppleAuthenticateDetails struct {
	UDID                  string `db:"udid" plist:"UDID"`
	Topic                 string `db:"topic" plist:"Topic"`
	OSVersion             string `db:"os_version" plist:"OSVersion"`
	BuildVersion          string `db:"build_version" plist:"BuildVersion"`
	ProductName           string `db:"product_name" plist:"ProductName"`
	SerialNumber          string `db:"serial_number" plist:"SerialNumber"`
	IMEI                  string `db:"imei" plist:"IMEI"`
	MEID                  string `db:"meid" plist:"MEID"`
	Token                 []byte `db:"token" plist:"Token"`
	PushMagic             string `db:"push_magic" plist:"PushMagic"`
	UnlockToken           []byte `db:"unlock_token" plist:"UnlockToken"`
	AwaitingConfiguration bool   `db:"-" plist:"AwaitingConfiguration"`

	DeviceRequest
}

//TODO
type DeviceRequest struct { //TODO: Add Plist Mapping Tags
	MessageType string `db:"-"` // Could Be Authenticate or TokenUpdate or CheckOut
	//Topic       string `db:"-"`
	//UDID        string `db:"-"`
}

// TODO REMOVE THIS TEMP  STRUCT
/*
type Devicey struct {
	UDID                  string `db:"udid" plist:"UDID"`
	Topic                 string `db:"topic" plist:"Topic"`
	OSVersion             string `db:"os_version" plist:"OSVersion"`
	BuildVersion          string `db:"build_version" plist:"BuildVersion"`
	ProductName           string `db:"product_name" plist:"ProductName"`
	SerialNumber          string `db:"serial_number" plist:"SerialNumber"`
	IMEI                  string `db:"imei" plist:"IMEI"`
	MEID                  string `db:"meid" plist:"MEID"`
	Token                 []byte `db:"token" plist:"Token"`
	PushMagic             string `db:"push_magic" plist:"PushMagic"`
	UnlockToken           []byte `db:"unlock_token" plist:"UnlockToken"`
	AwaitingConfiguration bool   `db:"awaiting_configuration" plist:"AwaitingConfiguration"`

	DeviceRequest `db:"-"`
}*/

//TODO
func (req *AppleAuthenticateDetails) PopulateRequestData(body io.ReadCloser) error {
	if err := plist.NewXMLDecoder(body).Decode(&req); err != nil {
		return err
	}
	return nil
}

//TODO
func (d *Device) LoadFromDB(db *sqlx.DB) error {
	err := db.Get(d, "SELECT * FROM devices WHERE udid=$1 LIMIT 1", d.UDID)
	if err != nil {
		return err
	}
	return nil
}

//TODO
func (d *Device) UpdateDB(db *sqlx.DB, sql string) error {
	_, err := db.NamedExec(sql, d)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
