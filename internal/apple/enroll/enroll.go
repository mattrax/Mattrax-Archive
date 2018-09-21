package enroll

import (
	"bytes"
	"compress/gzip"
	"log"
	"net/http"
	"strings"

	"github.com/groob/plist"
	"github.com/iancoleman/strcase"
)

var enrollCacheFile = "enrollmentCache.mobileconfig" //TODO: Add To Config/Env Var

// The Web Handler
func Handler(config map[string]string) func(w http.ResponseWriter, r *http.Request) error {
	var enrollmentProfile []byte
	var enrollmentProfileGzip []byte

	EnrollmentProfilePayload := AppleMDMProfile{ //TODO: Load Values From Config Or Generate Them
		PayloadContent: []interface{}{
			AppleMDMProfilePayload{
				PayloadContent: AppleMDMEnrollmentSCEPPayload{
					CAFingerprint: []byte("58L5tHezPdS1z9oiIDiXWYG1KqTQGxT5svzdrvv2kuA"),
					KeyType:       "RSA",
					KeyUsage:      0,
					Keysize:       2048,
					Name:          "Profile Manager Device Identity CA",
					//Subject       []interface{} `plist:"Subject"`
					URL: "https://scep.otbeaumont.me/scep",
				},
				PayloadDescription:  "Configures Your Devices Identity To The Mattrax Server.",
				PayloadDisplayName:  "Device Identity Certificate",
				PayloadIdentifier:   "com.apple.security.scep.9AD147E7-A9CB-45BD-96FE-672A2F422216", //strcase.ToSnake(config["OrganisationShortName"]) + "." + "TODO UDID HERE", //TEMP Comment    com.apple.config.Admins-Mac.local.mdm
				PayloadOrganization: "Mattrax MDM Server",
				PayloadType:         "com.apple.security.scep",
				PayloadUUID:         "9AD147E7-A9CB-45BD-96FE-672A2F422216", //TODO
				PayloadVersion:      1,
			},
			/*
				AppleMDMProfilePayload{
					//PayloadContent
					PayloadDescription:  "Configures Your Devices Identity To The Mattrax Server.",
					PayloadDisplayName:  "Device Identity Certificate",
					PayloadIdentifier:   "com.apple.security.scep.9AD147E7-A9CB-45BD-96FE-672A2F422216", //strcase.ToSnake(config["OrganisationShortName"]) + "." + "TODO UDID HERE", //TEMP Comment    com.apple.config.Admins-Mac.local.mdm
					PayloadOrganization: "Mattrax MDM Server",
					PayloadType:         "com.apple.security.scep",
					PayloadUUID:         "9AD147E7-A9CB-45BD-96FE-672A2F422216", //TODO
					PayloadVersion:      1,
				},
			*/
		},
		PayloadRemovalDisallowed: true,
		PayloadDescription:       "Allow Your Organisation To Maintain and Secure Your Device.",
		PayloadDisplayName:       config["OrganisationName"] + "'s MDM Server",
		PayloadIdentifier:        strcase.ToSnake(config["OrganisationShortName"]) + "." + "TODO UDID HERE", //TEMP Comment    com.apple.config.Admins-Mac.local.mdm
		PayloadOrganization:      config["OrganisationName"],
		PayloadType:              "Configuration",
		PayloadUUID:              "B6F27D01-2D4B-4B08-A927-7A9C6021AB9D", //TODO
		PayloadVersion:           1,
	}

	/*
		<key>ConsentText</key>
		<dict>
			<key>default</key>
			<string>Tetsing</string>
		</dict>
	*/

	var err error
	enrollmentProfile, err = plist.Marshal(EnrollmentProfilePayload)
	if err != nil {
		log.Fatal(err) //TODO: Better Handling
	}

	enrollmentProfileGzip, err = gZipData(enrollmentProfile)
	if err != nil {
		log.Fatal(err) //TODO: Better Handling
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/x-apple-aspen-config")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			w.Write(enrollmentProfileGzip)
		} else {
			w.Write(enrollmentProfile)
		}
		return nil
	}
}

func gZipData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err = gz.Write(data)
	if err != nil {
		return
	}

	if err = gz.Flush(); err != nil {
		return
	}

	if err = gz.Close(); err != nil {
		return
	}

	compressedData = b.Bytes()

	return
}

//TODO: Cache This To Disk -> So It Is The Same After Reboots Unless Info Changes
// TODO: If There Is An Error Creating The Profile Report Back And Handle: http.Error(w, err.Error(), http.StatusInternalServerError)
