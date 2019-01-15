package apple

import (
	"crypto/x509"

	"github.com/mastahyeti/cms"
	"github.com/mattrax/Mattrax/internal/certstore"

	"github.com/groob/plist"
	uuid "github.com/satori/go.uuid"
)

/* TODO: Create A Profile Service Which Handles Generating Profile and move this to it */
/* TODO: Create a service which handles pushing profiles to the device - Just like a queue */

type Profile struct {
	//TEMP: Identifier string
	Body []byte
}

/* TEMP:
func (p *Profile) GetDetails() (Details, error) {
	// TODO:
}


func (p *Profile) Validate() error {
	if len(p.Body) < 1 {
		return errors.New("blank mobileconfig body")
	}

	return nil
}
*/

func (p *Profile) Sign(certStore *certstore.CertStore) ([]byte, error) {
	profileSigned, err := cms.Sign(p.Body, []*x509.Certificate{certStore.SigningCert}, certStore.SigningKey)
	if err != nil {
		return []byte{}, err
	}

	return profileSigned, nil
}

type PlistProfile struct {
	PayloadIdentifier   string
	PayloadUUID         string
	PayloadDisplayName  string
	PayloadDescription  string `plist:",omitempty"`
	PayloadOrganization string `plist:",omitempty"`
	PayloadType         string
	PayloadVersion      int
	PayloadContent      interface{} `plist:",omitempty"`
}

func NewProfile(plistBody interface{}) (Profile, error) {
	body, err := plist.Marshal(plistBody)
	if err != nil {
		return Profile{}, err
	}

	return Profile{
		Body: body,
	}, nil
}

func NewUUID() (string, error) {
	return uuid.NewV4().String(), nil
}
