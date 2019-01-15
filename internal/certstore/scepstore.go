package certstore

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"math/big"
)

// This file implements: https://github.com/micromdm/scep/blob/master/depot/depot.go

func (cs *CertStore) CA(pass []byte) ([]*x509.Certificate, *rsa.PrivateKey, error) {
	fmt.Println("CA")
	return []*x509.Certificate{cs.ScepCert}, cs.ScepKey, nil
	//return nil, nil, nil
}

func (cs *CertStore) Put(name string, crt *x509.Certificate) error {
	fmt.Println("PUT: ", crt)
	return nil
}

func (cs *CertStore) Serial() (*big.Int, error) {
	fmt.Println("SERIAL")
	return big.NewInt(1), nil
}

func (cs *CertStore) HasCN(cn string, allowTime int, cert *x509.Certificate, revokeOldCertificate bool) error {
	fmt.Println("HAS CN: ", cn, allowTime, cert, revokeOldCertificate)
	return nil
}
