package management_payloads

import (
	"crypto/x509"

	"github.com/mastahyeti/cms"
)

var certPEM = []byte(`-----BEGIN CERTIFICATE-----
MIID2zCCAsOgAwIBAgIBATANBgkqhkiG9w0BAQsFADCBmjEbMBkGA1UEAwwSTWF0
dHJheCBUcnVzdCBDZXJ0MRwwGgYDVQQKDBNNYXR0cmF4IERldmVsb3BtZW50MQww
CgYDVQQLDANJQ1QxCzAJBgNVBAgMAldBMQswCQYDVQQGEwJBVTEQMA4GA1UEBwwH
QnVuYnVyeTEjMCEGCSqGSIb3DQEJARYUaGVscGRlc2tAbWF0dHJheC5jb20wHhcN
MTgwOTEzMTA0NDE4WhcNMTkwOTEzMTA0NDE4WjCBmjEbMBkGA1UEAwwSTWF0dHJh
eCBUcnVzdCBDZXJ0MRwwGgYDVQQKDBNNYXR0cmF4IERldmVsb3BtZW50MQwwCgYD
VQQLDANJQ1QxCzAJBgNVBAgMAldBMQswCQYDVQQGEwJBVTEQMA4GA1UEBwwHQnVu
YnVyeTEjMCEGCSqGSIb3DQEJARYUaGVscGRlc2tAbWF0dHJheC5jb20wggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC96A7fQ7CqbeJFtvUW4ws2xfWhpHVm
/SR9QZwiH7avYoJX6Ut2NV0tzxQIjX72B5WEsRTaRFhLkYDNWVbDFMsd7Djt+PH+
0gd70DAfPBdTkm99Sv2KkLOIiA9/jqF51RPqU4MbTZ74eFgE+kT/iBUdybCzBjQY
qtUhqSCXXTUTFDEDabr3EDXmj4r8gaEVzoGsq5iDZSVNeFme2M/yDtM//YZaO3Rn
aff4QnoQFblGbi5OhLsy1HdjSjTgeR63MUbkc22QJWbn+8dvyoFKEAmKDSaTst7k
IdRgjBB4oEhux7V+DX2eZIUR+AqsC1U/cOa97Zy5yTb7D6gCFveF/TU7AgMBAAGj
KjAoMA4GA1UdDwEB/wQEAwIHgDAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAzANBgkq
hkiG9w0BAQsFAAOCAQEAJDAwnu7IK8EzkADb1rUFzCW/2Rnfw3VREu0OgnBRSnNB
iSdA8HLP9FvQ8UErN3w6rV8HFJGuQFV0z+9N5vbRIPhF/LVKHlS8eeBY2gDvGffd
R6UHkSCN2/j1WoocSgT12lwcEvJiU2rGoLN9zr8Yhs8qRhJy0q7Puxlm2YRwHzD2
Yw/eTPSIFi9KLYD5QsOUMH4MP/s4rxgUedjZ2Dk2GB+H/EjB6/GmzvlY/0Hdk0Dd
F14j2fLYoQZ+RsOK8FQTiC+VEJLBtSBkcpR7m0Nr5ZOdF4gW4OR+1yjHmdLYMp0A
lYy8US91kqtu9dDh3e17eGuHxQUAgC+j0fLoTMmpxQ==
-----END CERTIFICATE-----`)

var certKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvegO30Owqm3iRbb1FuML
NsX1oaR1Zv0kfUGcIh+2r2KCV+lLdjVdLc8UCI1+9geVhLEU2kRYS5GAzVlWwxTL
Hew47fjx/tIHe9AwHzwXU5JvfUr9ipCziIgPf46hedUT6lODG02e+HhYBPpE/4gV
HcmwswY0GKrVIakgl101ExQxA2m69xA15o+K/IGhFc6BrKuYg2UlTXhZntjP8g7T
P/2GWjt0Z2n3+EJ6EBW5Rm4uToS7MtR3Y0o04HketzFG5HNtkCVm5/vHb8qBShAJ
ig0mk7Le5CHUYIwQeKBIbse1fg19nmSFEfgKrAtVP3Dmve2cuck2+w+oAhb3hf01
OwIDAQAB
-----END RSA PUBLIC KEY-----`)

func SignProfile(profile []byte) []byte { //TODO: Get CA Parsed In Too
	cert, _ := x509.ParseCertificate(certPEM)
	key, _ := x509.ParseECPrivateKey(certKey)

	der, _ := cms.Sign(profile, []*x509.Certificate{cert}, key)

	return der

	//roots := x509.NewCertPool()
	/*ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}*/

	/*block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}*/
	/*cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}*/

	//log.Println(pkcs7.Encypt(profile, []*x509.Certificate{cert.Certificate}))
}

////////////// TEMP //////////////
/*
func init() { // TODO: Test Parsing And Generating
	rawWifiProfile := Profile{
		PayloadDisplayName:       "Mattrax Wifi Configuration",
		PayloadIdentifier:        "oscar-beaumont.B7228FFE-ED2E-499A-96E1-475FB3F934E7",
		PayloadRemovalDisallowed: false,
		PayloadType:              "Configuration",
		PayloadUUID:              "C504AC1D-8FF5-47C0-A143-20E35ACBD204",
		PayloadVersion:           1,
		PayloadContent: []interface{}{
			WifiConfiguration{
				AutoJoin:           true,
				CaptiveBypass:      false,
				EncryptionType:     "WPA",
				HIDDEN_NETWORK:     false,
				IsHotspot:          false,
				Password:           "SecureWifiPassword",
				PayloadDescription: "Configures Wi-Fi settings",
				PayloadDisplayName: "Wi-Fi",
				PayloadIdentifier:  "com.apple.wifi.managed.E008CC37-A712-4CA3-8222-4F300CA3CCBE",
				PayloadType:        "com.apple.wifi.managed",
				PayloadUUID:        "E008CC37-A712-4CA3-8222-4F300CA3CCBE",
				PayloadVersion:     1,
				ProxyType:          "None",
				SSID_STR:           "TestingWifiSSID",
			},
		},
	}

	wifiProfile, _ := plist.MarshalIndent(rawWifiProfile, "   ")
	signedProfile := SignProfile(wifiProfile)
	log.Println(string(signedProfile))
}*/
