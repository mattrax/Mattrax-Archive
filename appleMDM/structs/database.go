/**
 * Mattrax: An Open Source Device Management System
 * File Description: This File Has All of The Structs Used For Communication With The Database.
 * Package Description: These Are The Structs and Helpers For Device Communication, The API and Database Communication.
 * A HUGE Thanks To MicroMDM. This Package Is A Modied Version Of The (github.com/micromdm/mdm) Package. It Is Used Under The MIT Licence and The Original Work Is Copyright Of MicroMDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package structs

//TODO: On All Structs Files Reodo notnull and omit if empty on everything. Remove All json serialising (For Now Redo For API)

type Device struct {
	TableName      struct{}       `sql:"devices"`
	UDID           string         `sql:"uuid,pk"`
	DeviceState    int            `sql:"DeviceState,notnull"`
	DeviceDetails  DeviceDetails  `sql:"DeviceDetails,notnull"`
	DeviceTokens   DeviceTokens   `sql:"DeviceTokens,notnull"`
	DevicePolicies DevicePolicies `sql:"DevicePolicies,notnull"`
}

type DeviceDetails struct {
	OSVersion    string            `sql:"OSVersion,notnull"`
	BuildVersion string            `sql:"BuildVersion,notnull"`
	ProductName  string            `sql:"ProductName,notnull"`
	SerialNumber string            `sql:"SerialNumber,notnull"`
	IMEI         string            `sql:"IMEI,notnull"`
	MEID         string            `sql:"MEID,notnull"`
	Profiles     []ProfileListItem `sql:"Profiles,notnull"`
}

type DeviceTokens struct {
	Token       []byte `sql:"Token,notnull"`
	PushMagic   string `sql:"PushMagic,notnull"`
	UnlockToken []byte `sql:"UnlockToken,notnull"`
}

type DevicePolicies struct {
	Installed  map[string]DevicePolicy `sql:"Installed,notnull"`
	LastUpdate int64                   `sql:"LastUpdate,notnull"`
	Queued     map[string]DevicePolicy `sql:"Queued,notnull"`
	Inventory  DevicePoliciesInventory `sql:"Inventory,notnull"`
}

type DevicePoliciesInventory struct {
	State      int               `sql:"State,notnull"`
	Commands   map[string]string `sql:"CommandUUIDs,notnull"` // [2]bool `sql:"CommandUUIDs,notnull"`
	LastUpdate int64             `sql:"LastUpdate,notnull"`
}

/*
type DevicePoliciesInventoryCommand struct {
	RequestType string `sql:"RequestType,notnull"`
	Status bool `sql:"Status,notnull"`
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
	//Actions []ServerCommand `sql:"Actions,notnull"` ///////// FIXME TEMP
}
*/

//TODO Doc
type DevicePolicy struct {
	Status int
	//Result
	//Policy Policy
}

//TODO Doc
type Policy struct {
	TableName struct{} `sql:"policies"`
	UDID      string   `sql:"uuid,pk"`

	Config  PolicyConfig `sql:"config,notnull"`
	Command Command      `sql:"command,notnull"` // PolicyOptions `sql:"options,notnull"`
	Content string       `sql:"content"`         //////////////////////////////////////////////// Does This Break if Parsed Null

	//Config    PolicyConfig  `sql:"config,notnull"`
	//Options   PolicyOptions `sql:"options,notnull"`
}

type PolicyConfig struct {
	Name       string   `sql:"name"`
	Targets    []string `sql:"targets"` //ie. Mac, IOS and Software Versions
	PolicyType string   `sql:"name"`
	//Type string
}

type PolicyOptions struct { // Lots of Optional Values
	InstallApplication
}
