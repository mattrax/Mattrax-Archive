package certstore

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	mathrand "math/rand"
	"path"
	"time"
)

// TODO: Deal With Renewal

var organisation = "Acme School Inc" // TODO: Load From Config
var certName = "Acme School Inc Signing Identity"

func generateSigningCert(certDir string) ([]byte, []byte, error) { // TODO: Log This Happening - Wrap Errors and Clean Output
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err // failed to generate private key
	}

	ExtraExtensions, err := asn1.Marshal([]asn1.ObjectIdentifier{asn1.ObjectIdentifier{1, 2, 840, 113635, 100, 4, 13}})
	if err != nil {
		return nil, nil, err // failed to create extension
	}

	NotBefore := time.Now().Add(time.Duration(mathrand.Int31n(120)) * -time.Minute) // This takes a random amount time from now within the last hour (The Negative time.Minute)
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:         certName,
			Organization:       []string{organisation},
			OrganizationalUnit: []string{"Mattrax Server"},
		},
		NotBefore:   NotBefore,
		NotAfter:    NotBefore.Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageCodeSigning}, // TODO: Fix This Line ???? What is wrong with it
		ExtraExtensions: []pkix.Extension{pkix.Extension{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 37},
			Critical: true,
			Value:    ExtraExtensions,
		}, pkix.Extension{
			Id:       asn1.ObjectIdentifier{1, 2, 840, 113635, 100, 6, 1, 14},
			Critical: true,
			Value:    []byte{0x05, 0x00},
		}},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	certDer, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, nil, err // failed to create the certificate
	}

	certPemBlock := &pem.Block{Type: "CERTIFICATE", Bytes: certDer}
	certPem := pem.EncodeToMemory(certPemBlock)
	if certPem == nil {
		return nil, nil, pem.Encode(ioutil.Discard, certPemBlock) // failed to encode the certificate as a PEM
	}

	keyPemBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	keyPem := pem.EncodeToMemory(keyPemBlock)
	if keyPem == nil {
		return nil, nil, pem.Encode(ioutil.Discard, keyPemBlock) // failed to encode the private key as a PEM
	}

	if err := ioutil.WriteFile(path.Join(certDir, "signing-cert.pem"), certPem, 0644); err != nil {
		return nil, nil, err // error saving the certificate
	}
	if err := ioutil.WriteFile(path.Join(certDir, "signing-cert.key"), keyPem, 0644); err != nil {
		return nil, nil, err // error saving the private key
	}

	return certPem, keyPem, nil
}
