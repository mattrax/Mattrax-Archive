package checkin

import "encoding/hex"

// CheckinCommand represents an MDM checkin request.
type CheckinCommand struct {
	MessageType string // Authenticate, TokenUpdate or CheckOut
	Topic       string
	UDID        string
	authenticate
	tokenUpdate
	// TODO: checkout
}

// authenticate represents the specific fields for the 'Authenticate' message type of the MDM checkin request.
type authenticate struct {
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

// tokenUpdate represents the specific fields for the 'TokenUpdate' message type of the MDM checkin request.
type tokenUpdate struct {
	Token                 hexData
	PushMagic             string
	UnlockToken           hexData
	AwaitingConfiguration bool

	// User Specific Details
	UserID        string `plist:",omitempty"`
	UserLongName  string `plist:",omitempty"`
	UserShortName string `plist:",omitempty"`
	NotOnConsole  bool   `plist:",omitempty"`
}

// hexData is a special type for storing hex data that has an attached 'String()' func.
// Thanks to the MicroMDM code base for this idea.
type hexData []byte

func (d hexData) String() string {
	return hex.EncodeToString(d)
}
