package certstore

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path"
)

func loadSigningCert(certDir string) (*x509.Certificate, *rsa.PrivateKey, error) { // TODO: Wrap Errors
	certFile, certErr := ioutil.ReadFile(path.Join(certDir, "signing-cert.pem")) // TOOD: Set the file names to const vars
	keyFile, keyErr := ioutil.ReadFile(path.Join(certDir, "signing-cert.key"))

	if os.IsNotExist(certErr) || os.IsNotExist(certErr) {
		var err error
		certFile, keyFile, err = generateSigningCert(certDir)
		if err != nil {
			return nil, nil, err
		}
	} else if certErr != nil {
		return nil, nil, certErr
	} else if keyErr != nil {
		return nil, nil, keyErr
	}

	certPem, _ := pem.Decode(certFile)
	keyPem, _ := pem.Decode(keyFile)

	cert, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		return nil, nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(keyPem.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}

func loadScepCert(certDir string) (*x509.Certificate, *rsa.PrivateKey, error) {
	return loadSigningCert(certDir) // TEMP
	//return nil, nil, nil
}
