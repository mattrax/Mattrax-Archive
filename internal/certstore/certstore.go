package certstore

import (
	"crypto/rsa"
	"crypto/x509"
	"os"
)

type CertStore struct {
	SigningCert *x509.Certificate
	SigningKey  *rsa.PrivateKey
	ScepCert    *x509.Certificate
	ScepKey     *rsa.PrivateKey
}

func New(certDir string) (*CertStore, error) { // TODO: Wrap Errors
	err := os.MkdirAll(certDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	signCert, signKey, err := loadSigningCert(certDir)
	if err != nil {
		return nil, err
	}

	scepCert, scepKey, err := loadScepCert(certDir)
	if err != nil {
		return nil, err
	}

	return &CertStore{signCert, signKey, scepCert, scepKey}, nil
}
