package mdm

type Setting struct {
	Item       string            `json:"item"`
	Enabled    *bool             `plist:",omitempty" json:"enabled,omitempty"`
	DeviceName *string           `plist:",omitempty" json:"device_name,omitempty"`
	HostName   *string           `plist:",omitempty" json:"hostname,omitempty"`
	Identifier *string           `plist:",omitempty" json:"identifier"`
	Attributes map[string]string `plist:",omitempty" json:"attributes,omitempty"`
}
