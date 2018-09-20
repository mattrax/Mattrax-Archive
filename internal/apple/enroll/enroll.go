package enroll

import (
	"bytes"
	"compress/gzip"
	"log"
	"net/http"
	"strings"

	"github.com/groob/plist"
)

var enrollCacheFile = "enrollmentCache.mobileconfig" //TODO: Add To Config/Env Var

// The Web Handler
func Handler(config map[string]string) func(w http.ResponseWriter, r *http.Request) error {
	var enrollmentProfile []byte
	var enrollmentProfileGzip []byte

	EnrollmentProfilePayload := AppleMDMProfile{ //TODO: Load Values From Config Or Generate Them
		PayloadContent: []interface{}{
			AppleMDMEnrollmentCertificateProfile{
				Password:                   "password",
				PayloadCertificateFileName: "PushCert.p12",
			},
			AppleMDMEnrollmentProfile{},
		},
		PayloadDescription:  "Allows remote management of your device by your administrator.",
		PayloadDisplayName:  config["OrganisationName"] + " MDM Server", //TODO: Deal With Extremly Long Org Names
		PayloadIdentifier:   "com.apple.config.Admins-Mac.local.mdm",
		PayloadOrganization: "Mattrax Academy",
		PayloadType:         "Configuration",
		PayloadUUID:         "B6F27D01-2D4B-4B08-A927-7A9C6021AB9D",
		PayloadVersion:      1,
	}

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
		//w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
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
