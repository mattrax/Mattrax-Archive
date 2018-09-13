package management_payloads

//TODO: XML Mappings and Omit Empty -> Get From Apple Configuration Profiles Documentation

type Profile struct {
	PayloadDisplayName       string
	PayloadIdentifier        string
	PayloadRemovalDisallowed bool
	PayloadType              string
	PayloadUUID              string
	PayloadVersion           int32
	PayloadContent           []interface{}
}

type WifiConfiguration struct {
	AutoJoin           bool
	CaptiveBypass      bool
	EncryptionType     string
	HIDDEN_NETWORK     bool
	IsHotspot          bool
	Password           string
	PayloadDescription string
	PayloadDisplayName string
	PayloadIdentifier  string
	PayloadType        string
	PayloadUUID        string
	PayloadVersion     int32
	ProxyType          string
	SSID_STR           string
}
