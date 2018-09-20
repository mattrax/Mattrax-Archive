package checkin

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
