package endpoints

import (
	"net/http"

	"github.com/mattrax/Mattrax/platform/apple"
	"github.com/rs/zerolog/log"
)

func enrollHandler(svc *Service) http.HandlerFunc {
	profile, err := svc.SignedEnrollmentProfile()
	if err != nil {
		log.Fatal().Err(err).Msg("Error Creating The Enrollment Profile")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-apple-aspen-config")
		w.Write(profile)
	}
}

// TODO: Cache Profile Because The Profile ID will keep changing

/* TODO: BELOW IS TEMP */

const (
	// EnrollmentProfileID is the 'PayloadIdentifier' (Identifier For The Profile) for the Enrollment profile
	EnrollmentProfileID string = "com.github.mattrax.Mattrax"
	// EnrollmentPayloadID is the 'PayloadIdentifier' (Identifier For The Payload) for the MDM Enrollment Payload in the Enrollment profile
	EnrollmentPayloadID string = "com.github.mattrax.Mattrax.enroll"
	// IdentityPayloadID is the 'PayloadIdentifier' (Identifier For The Payload) for the SCEP Enrollment Payload in the Enrollment profile
	IdentityPayloadID string = "com.github.mattrax.Mattrax.identity"
)

// SCEPPayload is the payload returned in the Enrollment profile to configure SCEP on the device
type SCEPPayload struct {
	Challenge  string
	KeyType    string `plist:"Key Type"`
	KeyUsage   int    `plist:"Key Usage"`
	Keysize    int
	Retries    int
	RetryDelay int
	URL        string
	Subject    [][][]string
}

// EnrollmentPayload is the payload returned in the Enrollment profile to configure MDM on the device
type EnrollmentPayload struct {
	AccessRights            int
	CheckInURL              string
	ServerURL               string
	SignMessage             bool
	Topic                   string
	UseDevelopmentAPNS      bool
	CheckOutWhenRemoved     bool
	IdentityCertificateUUID string
	ServerCapabilities      []string
	apple.PlistProfile
}

func (svc *Service) SignedEnrollmentProfile() ([]byte, error) { // TODO: Cleanup + Caching Result
	payloadContent := []interface{}{}

	identityUUID, err := apple.NewUUID()
	if err != nil {
		return nil, err
	}

	payloadContent = append(payloadContent, apple.PlistProfile{
		PayloadIdentifier:  IdentityPayloadID,
		PayloadUUID:        identityUUID,
		PayloadDisplayName: "Mattrax Identity",
		PayloadType:        "com.apple.security.scep",
		PayloadVersion:     1,
		PayloadContent: SCEPPayload{
			Challenge:  svc.SCEPChallenge,
			KeyType:    "RSA",
			KeyUsage:   0,
			Keysize:    1024, // TODO: MicroMDM is 2048 but The Test Profile is 1024
			Retries:    3,
			RetryDelay: 10,
			URL:        svc.PublicURL + "/apple/scep",
			Subject: [][][]string{
				[][]string{
					[]string{
						"O",
						"Acme",
					},
				},
				[][]string{
					[]string{
						"CN",
						"Mattrax Identity (%ComputerName%)", /* TODO: This ComputerName is not accepted by IOS Devices */
					},
				},
			},
		},
	})

	enrollmentUUID, err := apple.NewUUID()
	if err != nil {
		return nil, err
	}

	payloadContent = append(payloadContent, EnrollmentPayload{
		AccessRights:            8191,
		CheckInURL:              svc.PublicURL + "/apple/checkin",
		ServerURL:               svc.PublicURL + "/apple/server",
		SignMessage:             true,
		Topic:                   svc.Topic,
		UseDevelopmentAPNS:      false,
		CheckOutWhenRemoved:     true,
		IdentityCertificateUUID: identityUUID,
		ServerCapabilities:      []string{"com.apple.mdm.per-user-connections"},
		PlistProfile: apple.PlistProfile{
			PayloadIdentifier:  EnrollmentPayloadID,
			PayloadUUID:        enrollmentUUID,
			PayloadDisplayName: "Mattrax Management",
			PayloadType:        "com.apple.mdm",
			PayloadVersion:     1,
		},
	})

	profileUUID, err := apple.NewUUID()
	if err != nil {
		return nil, err
	}

	if svc.ProfileDescription == "" {
		svc.ProfileDescription = "Allows for your IT admins to manage and secure your device."
	}

	rawProfile := apple.PlistProfile{
		PayloadIdentifier:   EnrollmentProfileID,
		PayloadUUID:         profileUUID,
		PayloadDisplayName:  svc.TenantName,
		PayloadDescription:  svc.ProfileDescription,
		PayloadOrganization: "Mattrax MDM Server",
		PayloadType:         "Configuration",
		PayloadVersion:      1,
		PayloadContent:      payloadContent,
	}

	profile, err := apple.NewProfile(rawProfile)
	if err != nil {
		return nil, err
	}

	return profile.Sign(svc.CertStore)
}
