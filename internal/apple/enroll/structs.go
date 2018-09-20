package enroll

// TODO
type AppleMDMProfile struct {
	PayloadContent      []interface{} `plist:"PlayloadContent"`
	PayloadDescription  string        `plist:"PayloadDescription"`
	PayloadDisplayName  string        `plist:"PayloadDisplayName"`
	PayloadIdentifier   string        `plist:"PayloadIdentifier"`
	PayloadOrganization string        `plist:"PayloadOrganization"`
	PayloadType         string        `plist:"PayloadType"`
	PayloadUUID         string        `plist:"PayloadUUID"`
	PayloadVersion      uint32        `plist:"PayloadVersion"`
}

// TODO
type AppleMDMEnrollmentSCEPPayload struct {
	Password                   string `plist:"Password"`
	PayloadCertificateFileName string `plist:"PayloadCertificateFileName"`
	PayloadContent             []byte `plist:"PayloadContent"`
}

// TODO
type AppleMDMEnrollmentProfile struct {
}
