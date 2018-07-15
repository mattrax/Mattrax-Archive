/**
 * Mattrax: An Open Source Device Management System
 * File Description: This File Has All of The Structs For The Checkin Hanlder.
 * Package Description: These Are The Structs For The Plist (Device Communication) and JSON (API) Data.
 * A HUGE Thanks To MicroMDM. This Package Is A Modied Version Of The (github.com/micromdm/mdm) Package. It Is Used Under The MIT Licence and The Original Work Is Copyright Of MicroMDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package mdm

import "encoding/hex"

// CheckinCommand represents an MDM checkin command struct
type CheckinCommand struct {
	MessageType string // Either Authenticate, TokenUpdate or CheckOut
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
	DeviceName   string `plist:"DeviceName,omitempty"`
	Challenge    []byte `plist:"Challenge,omitempty"`
	Model        string `plist:"Model,omitpempty"`
	ModelName    string `plist:"ModelName,omitempty"`
}

// TokenUpdate Mesage Type
type update struct {
	Token                 hexData
	PushMagic             string
	UnlockToken           hexData
	AwaitingConfiguration bool
	userTokenUpdate
}

// TokenUpdate with user keys
type userTokenUpdate struct {
	UserID        string `plist:",omitempty"`
	UserLongName  string `plist:",omitempty"`
	UserShortName string `plist:",omitempty"`
	NotOnConsole  bool   `plist:",omitempty"`
}

// DEPEnrollmentRequest is a request sent
// by the device to an MDM server during
// DEP Enrollment
/*
type DEPEnrollmentRequest struct {
	Language string `plist:"LANGUAGE"`
	Product  string `plist:"PRODUCT"`
	Serial   string `plist:"SERIAL"`
	UDID     string `plist:"UDID"`
	Version  string `plist:"VERSION"`
	IMEI     string `plist:"IMEI,omitempty"`
	MEID     string `plist:"MEID,omitempty"`
}
*/

// Custom Format With .String() Function
// For Easy Conversion
type hexData []byte

func (d hexData) String() string {
	return hex.EncodeToString(d)
}
