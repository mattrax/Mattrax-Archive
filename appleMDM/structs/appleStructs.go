/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is The Apple MDM Structs. These Are The Go Structs For Database Communication and Plist Generation.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleStruct

///// Device Model /////
type Device struct {
	TableName      struct{}       `sql:"devices"`
	UDID           string         `sql:"uuid,pk"`
	DeviceState    int            `sql:"DeviceState,notnull"`
	DeviceDetails  DeviceDetails  `sql:"DeviceDetails,notnull"`
	DeviceTokens   DeviceTokens   `sql:"DeviceTokens,notnull"`
	DevicePolicies DevicePolicies `sql:"DevicePolicies,notnull"`
}

type DeviceDetails struct {
	OSVersion    string `sql:"OSVersion,notnull"`
	BuildVersion string `sql:"BuildVersion,notnull"`
	ProductName  string `sql:"ProductName,notnull"`
	SerialNumber string `sql:"SerialNumber,notnull"`
	IMEI         string `sql:"IMEI,notnull"`
	MEID         string `sql:"MEID,notnull"`
}

type DeviceTokens struct {
	Token       []byte `sql:"Token,notnull"`
	PushMagic   string `sql:"PushMagic,notnull"`
	UnlockToken []byte `sql:"UnlockToken,notnull"`
}

type DevicePolicies struct {
	CurrentAction DeviceCurrentAction `sql:"CurrentAction,notnull"` // FIXME: Probs This Is Optional
	Queued    []string `sql:"Queued,notnull"`
	Installed []string `sql:"Installed,notnull"`
	LastUpdate int64 `sql:"LastUpdate,notnull"`
}

type DeviceCurrentAction struct {
	UDID string `sql:"UDID,notnull"`
	Name string `sql:"Name,notnull"`
	Actions []ServerCommand `sql:"Actions,notnull"`
}







///// Policies Model /////
type Policy struct {
	TableName struct{}      `sql:"policies"`
	UDID      string        `sql:"uuid,pk"`
	Config    PolicyConfig  `sql:"config,notnull"`
	Options   PolicyOptions `sql:"options,notnull"`
	Content   string        `sql:"content"` //////////////////////////////////////////////// Does This Break if Parsed Null
}

type PolicyConfig struct {
	Name       string   `sql:"name"`
	Targets    []string `sql:"targets"` //ie. Mac, IOS and Software Versions
	PolicyType string   `sql:"name"`
	//Type string
}

type PolicyOptions struct { // Lots of Optional Values
	PayloadInstallApplication
}

//TODO: Add "Not NULLS" To Everything

/* End */

// CheckinRequest represents an MDM checkin command struct.
type CheckinCommand struct {
	MessageType string // Could Be Authenticate or TokenUpdate or CheckOut
	Topic       string
	UDID        string
	Auth
	Update
}

// Authenticate Message Type
type Auth struct {
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
type Update struct {
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

/* Device Response */
type DeviceStatus struct {
	UDID        string
	CommandUUID string
	Status      string
	ErrorChain
}

type ErrorChain struct {
	LocalizedDescription string
	USEnglishDescription string ///// TODO: This Should be Optional
	ErrorDomain          string
	ErrorCode            int
}

/* Server Response */
type ServerCommand struct {
	CommandUUID string `plist:"CommandUUID,notnull"` //Check notnull work for plist
	Command     ServerCommandBody `plist:"ServerCommandBody,notnull"` //TODO: Replace With Any Stuct (interface)
}

type ServerCommandBody struct { //TODO: Is This Used
	RequestType string `plist:"RequestType,notnull"`
	PayloadInstallApplication
	PayloadInstallProfile
	PayloadDeviceInformation
}

type PayloadInstallApplication struct {
	ITunesStoreID   int `plist:"iTunesStoreID,omitempty"` //TODO: Should I Have ,omitempty
	ManagementFlags int `plist:",omitempty"`              //TODO: Should I Have ,omitempty
}

type PayloadInstallProfile struct {
	Payload []byte `plist:"Payload,omitempty"` //TODO Add This Here //TODO: Should I Have ,omitempty
}

type PayloadDeviceInformation struct {
	Queries []string `plist:"Queries,omitempty"` //TODO: Should I Have ,omitempty
}
