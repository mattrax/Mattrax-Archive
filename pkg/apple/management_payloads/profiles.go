package management_payloads

import (
	"log"

	"github.com/fullsailor/pkcs7"
	"github.com/groob/plist"
)

func SignProfile(profile []byte) { //TODO: Get CA Parsed In Too
	return pkcs7.Encypt(profile)
}

////////////// TEMP //////////////

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
	signedProfile, _ := SignProfile(wifiProfile)
	log.Println(string(signedProfile))
}
