/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Apple MDM Stricts. These Are The Go Structs For Database Communication and Plist Generation.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

// Devices Database Table
type Device struct {
  TableName struct{} `sql:"devices"`
  UDID          string `sql:"udid,pk"`
  // Device Details
  OSVersion    string `sql:"OSVersion"`
	BuildVersion string `sql:"BuildVersion"`
	ProductName  string `sql:"ProductName"`
	SerialNumber string `sql:"SerialNumber"`
	IMEI         string `sql:"IMEI"`
	MEID         string `sql:"MEID"`
  // APNS
  Token         []byte `sql:"Token"`
  PushMagic     string `sql:"PushMagic"`
  UnlockToken   []byte `sql:"UnlockToken"`
  //Status
  Deployed      bool `sql:"Deployed,notnull"`

  //Registered Time (To Detect Deployed Errors)
}










// CheckinRequest represents an MDM checkin command struct.
type CheckinCommand struct {
	MessageType string // Could Be Authenticate or TokenUpdate or CheckOut
	Topic       string
	UDID        string
	auth
	update
}

// Authenticate Message Type
type auth struct {
	OSVersion    string
	BuildVersion string
	ProductName  string
	SerialNumber string
	IMEI         string
	MEID         string
	DeviceName   string `plist:"DeviceName,omitempty"` //TODO: Do I Need These/What Devices Send It
	Challenge    []byte `plist:"Challenge,omitempty"`  //TODO: Do I Need These/What Devices Send It
	Model        string `plist:"Model,omitpempty"`     //TODO: Do I Need These/What Devices Send It
	ModelName    string `plist:"ModelName,omitempty"`  //TODO: Do I Need These/What Devices Send It
}

// TokenUpdate Mesage Type
type update struct {
	Token                 []byte
	PushMagic             string
	UnlockToken           []byte
	AwaitingConfiguration bool
	userTokenUpdate       //TODO: Do I Need These/What Devices Send It
}

// TokenUpdate with user keys
type userTokenUpdate struct { //TODO: Do I Need These/What Devices Send It
	UserID        string `plist:",omitempty"`
	UserLongName  string `plist:",omitempty"`
	UserShortName string `plist:",omitempty"`
	NotOnConsole  bool   `plist:",omitempty"`
}





type ServerCommand struct { //TODO: Is This Used
	UDID   string
	//Status string
}
