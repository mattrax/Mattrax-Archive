package structs

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
	Token                 []byte
	PushMagic             string
	UnlockToken           []byte
	AwaitingConfiguration bool // TODO: Handle This

	// User Specific Details
	UserID        string `plist:",omitempty"`
	UserLongName  string `plist:",omitempty"`
	UserShortName string `plist:",omitempty"`
	NotOnConsole  bool   `plist:",omitempty"`
}
