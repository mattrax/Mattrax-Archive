/* A HUGE Thanks To MicroMDM. This Package Is A Modied Version Of The (github.com/micromdm/mdm) Package. It Is Used Under The MIT Licence and The Original Work Is Copyright Of MicroMDM. */

package structs

// CheckinCommand represents an MDM checkin command struct
type CheckinCommand struct {
	MessageType string // Either Authenticate, TokenUpdate or CheckOut
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
	DeviceName   string `plist:"DeviceName,omitempty"`
	Challenge    []byte `plist:"Challenge,omitempty"`
	Model        string `plist:"Model,omitpempty"`
	ModelName    string `plist:"ModelName,omitempty"`
}

// TokenUpdate Mesage Type
type Update struct {
	Token                 []byte
	PushMagic             string
	UnlockToken           []byte
	AwaitingConfiguration bool
	userTokenUpdate       //TODO: Do I Need These/What Devices Send It and Handle or Remove it
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
