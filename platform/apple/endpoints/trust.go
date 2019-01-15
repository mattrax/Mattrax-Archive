package endpoints

import (
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/mattrax/Mattrax/platform/apple"
	"github.com/rs/zerolog/log"
)

// TODO: This returns the Mattrax trust profile - For signing or maybe put it in the enrollment profile???

func trustHandler(svc *Service) http.HandlerFunc {
	profile, err := svc.SignedTrustProfile()
	if err != nil {
		log.Fatal().Err(err).Msg("Error Creating The Trust Profile")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-apple-aspen-config")
		w.Write(profile)
	}
}

/* Below is TEMP: Should be Moved */

const (
	// TrustProfileID is the 'PayloadIdentifier' for the Trust profile
	TrustProfileID string = "com.github.mattrax.Mattrax.trust"
	// SigningCertPayloadID is the 'PayloadIdentifier' for the Signing Certificate Payload in the Trust certificate
	SigningCertPayloadID string = "com.github.mattrax.Mattrax.trust-certificate"
)

func (svc *Service) SignedTrustProfile() ([]byte, error) { // TODO: Cleanup + Caching Result
	payloadContent := []interface{}{}

	signingCertPayloadUUID, err := apple.NewUUID()
	if err != nil {
		return nil, err
	}

	certFile, _ := ioutil.ReadFile(path.Join("./certs", "signing-cert.pem")) // TODO: This is TEMP and should not need to touch the disk - and know the file name
	certPem, _ := pem.Decode(certFile)
	if certPem.Bytes == nil {
		log.Fatal().Msg("Error Decoding The Signing Certificate for the Trust Profile")
	}

	payloadContent = append(payloadContent, apple.PlistProfile{
		PayloadIdentifier:  SigningCertPayloadID,
		PayloadUUID:        signingCertPayloadUUID,
		PayloadDisplayName: "Mattrax Signing Identity",
		PayloadType:        "com.apple.security.pem",
		PayloadVersion:     1,
		PayloadContent:     certPem.Bytes,
	})

	// TODO: If Using Self Signed HTTPS Cert Add It

	profileUUID, err := apple.NewUUID()
	if err != nil {
		return nil, err
	}

	rawProfile := apple.PlistProfile{
		PayloadIdentifier:   TrustProfileID,
		PayloadUUID:         profileUUID,
		PayloadDisplayName:  svc.TenantName + " Trust Profile",
		PayloadDescription:  "Configures your device to trust the Mattrax server",
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
