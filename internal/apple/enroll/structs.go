package enroll

// TODO
type AppleMDMProfile struct {
	PayloadContent           []interface{} `plist:"PlayloadContent"`
	PayloadRemovalDisallowed bool          `plist:"PayloadRemovalDisallowed,omitempty"`
	PayloadDescription       string        `plist:"PayloadDescription"`
	PayloadDisplayName       string        `plist:"PayloadDisplayName"`
	PayloadIdentifier        string        `plist:"PayloadIdentifier"`
	PayloadOrganization      string        `plist:"PayloadOrganization"`
	PayloadType              string        `plist:"PayloadType"`
	PayloadUUID              string        `plist:"PayloadUUID"`
	PayloadVersion           uint32        `plist:"PayloadVersion"`
}

// TODO
type AppleMDMProfilePayload struct {
	PayloadContent      interface{} `plist:"PlayloadContent"`
	PayloadDescription  string      `plist:"PayloadDescription"`
	PayloadDisplayName  string      `plist:"PayloadDisplayName"`
	PayloadIdentifier   string      `plist:"PayloadIdentifier"`
	PayloadOrganization string      `plist:"PayloadOrganization"`
	PayloadType         string      `plist:"PayloadType"`
	PayloadUUID         string      `plist:"PayloadUUID"`
	PayloadVersion      uint32      `plist:"PayloadVersion"`
}

// TODO
type AppleMDMEnrollmentSCEPPayload struct {
	CAFingerprint []byte `plist:"CAFingerprint"`
	KeyType       string `plist:"Key Type"`
	KeyUsage      int    `plist:"Key Usage"`
	Keysize       int    `plist:"Keysize"`
	Name          string `plist:"Name"`
	//Subject       []interface{} `plist:"Subject"`
	URL string `plist:"URL"`
}

// TODO
type AppleMDMEnrollmentProfile struct {
}
