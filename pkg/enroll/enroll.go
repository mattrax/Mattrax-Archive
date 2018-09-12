package enroll

import (
	"log"
	"net/http"
	"sync"

	"github.com/groob/plist"
)

func EnrollHandler() http.HandlerFunc {
	var (
		init                 sync.Once
		rawEnrollmentProfile []byte
		//err               error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		type AppleMDMProfile struct {
			PayloadContent      []interface{} `plist:"PlayloadContent"`
			PayloadDescription  string        `plist:"PayloadDescription"`
			PayloadDisplayName  string        `plist:"PayloadDisplayName"`
			PayloadIdentifier   string        `plist:"PayloadIdentifier"`
			PayloadOrganization string        `plist:"PayloadOrganization"`
			PayloadType         string        `plist:"PayloadType"`
			PayloadUUID         string        `plist:"PayloadUUID"`
			PayloadVersion      uint32        `plist:"PayloadVersion"`
		}

		type AppleMDMEnrollmentCertificateProfile struct {
			Password                   string `plist:"Password"`
			PayloadCertificateFileName string `plist:"PayloadCertificateFileName"`
			PayloadContent             []byte `plist:"PayloadContent"`
		}

		type AppleMDMEnrollmentProfile struct {
		}

		init.Do(func() {
			enrollmentProfile := AppleMDMProfile{ //TODO: Load Values From Config Or Generate Them
				PayloadContent: []interface{}{
					AppleMDMEnrollmentCertificateProfile{
						Password:                   "password",
						PayloadCertificateFileName: "PushCert.p12",
					},
					AppleMDMEnrollmentProfile{},
				},
				PayloadDescription:  "Allows remote management of your device by your administrator.",
				PayloadDisplayName:  "Mattrax MDM Server",
				PayloadIdentifier:   "com.apple.config.Admins-Mac.local.mdm",
				PayloadOrganization: "Mattrax Academy",
				PayloadType:         "Configuration",
				PayloadUUID:         "B6F27D01-2D4B-4B08-A927-7A9C6021AB9D",
				PayloadVersion:      1,
			}

			var err error
			rawEnrollmentProfile, err = plist.Marshal(enrollmentProfile)
			if err != nil {
				log.Fatal(err) //TODO: Better Handling
			}

			//TODO: Cache This To Disk -> So It Is The Same After Reboots Unless Info Changes
			// TODO: If There Is An Error Creating The Profile Report Back And Handle: http.Error(w, err.Error(), http.StatusInternalServerError)
		})

		//w.Header().Set("Content-Type", r.Header.Get("Content-Type")) //TODO: Correct This
		w.Write(rawEnrollmentProfile)
	}
}
