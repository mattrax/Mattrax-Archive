package certificates

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Store struct {
	Cert *x509.Certificate
	Key  *rsa.PrivateKey
}

func NewStore(certFile string, keyFile string) (Store, error) { // TODO: Cleanup
	// TODO: Load Or Generate The Cert Here
	cf, err := ioutil.ReadFile(certFile)
	if err != nil {
		return Store{}, errors.Wrap(err, "cfload")
	}

	kf, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return Store{}, errors.Wrap(err, "kfload")
	}

	cpb, _ := pem.Decode(cf)
	kpb, _ := pem.Decode(kf)

	crt, err := x509.ParseCertificate(cpb.Bytes)
	if err != nil {
		return Store{}, errors.Wrap(err, "parsex509")
	}

	key, err := x509.ParsePKCS1PrivateKey(kpb.Bytes)
	if err != nil {
		return Store{}, errors.Wrap(err, "parsekey")
	}

	return Store{crt, key}, nil
}
