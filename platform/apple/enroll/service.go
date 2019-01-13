package enroll

import (
	"github.com/mattrax/Mattrax/internal/certificates"
	"github.com/mattrax/Mattrax/internal/config"
)

const (
	// EnrollmentProfileID is the 'PayloadIdentifier' (Identifier For The Profile) for the Enrollment profile
	EnrollmentProfileID string = "com.github.mattrax.Mattrax"
	// EnrollmentPayloadID is the 'PayloadIdentifier' (Identifier For The Payload) for the MDM Enrollment Payload in the Enrollment profile
	EnrollmentPayloadID string = "com.github.mattrax.Mattrax.enroll"
	// IdentityPayloadID is the 'PayloadIdentifier' (Identifier For The Payload) for the SCEP Enrollment Payload in the Enrollment profile
	IdentityPayloadID string = "com.github.mattrax.Mattrax.identity"
)

// Service contains the dependencies and functions for the Service
type Service struct {
	CertStore          certificates.Store
	PublicURL          string // Protocol + Domain + Port
	TenantName         string
	ProfileDescription string // Optional
	SCEPChallenge      string
	Topic              string
}

// New returns a new Service
func New(certStore certificates.Store, config config.Config) Service {
	return Service{
		CertStore:     certStore,
		PublicURL:     config.PublicURL,
		TenantName:    config.TenantName,
		SCEPChallenge: "secret",                                                      // TODO: Deal With This
		Topic:         "com.apple.mgmt.XServer.232a74b5-7a81-4d6c-82fa-00f351e2c4f9", // TODO: Deal With This
	}
}
