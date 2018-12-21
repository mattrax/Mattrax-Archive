// Content in this file is used under the MIT Licence (Copyright (c) 2016 Victor Vrantchan) from MicroMDM (https://github.com/micromdm/micromdm)
package appleHttp

import "encoding/hex"

// CheckinRequest represents an MDM checkin command struct.
type CheckinCommand struct {
	// MessageType can be either Authenticate,
	// TokenUpdate or CheckOut
	MessageType string
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

// data decodes to []byte,
// we can then attach a string method to the type
// Tokens are encoded as Hex Strings
type hexData []byte

func (d hexData) String() string {
	return hex.EncodeToString(d)
}
