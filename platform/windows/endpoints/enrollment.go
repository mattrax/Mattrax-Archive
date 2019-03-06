package endpoints

import (
	"log"
	"net/http"

	"github.com/Zauberstuhl/go-xml"
	"github.com/mattrax/Mattrax/pkg/soap"
)

func EnrollmentPolicyEndpoint() http.HandlerFunc {
	type req struct {
		soap.Envelope
	}

	type res struct {
		soap.Envelope

		// TODO: Future: The s:Body doen't have The Attrs xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"
		// TODO: Future: The GetPoliciesResponse Doesn't Have The Attr xmlns="http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy"
		// AuthPolicy                 string `xml:"s:Body>GetPoliciesResponse>response>AuthPolicy"`
		// EnrollmentVersion          string `xml:"s:Body>DiscoverResponse>DiscoverResult>EnrollmentVersion"`
		// EnrollmentPolicyServiceUrl string `xml:"s:Body>DiscoverResponse>DiscoverResult>EnrollmentPolicyServiceUrl"`
		// EnrollmentServiceUrl       string `xml:"s:Body>DiscoverResponse>DiscoverResult>EnrollmentServiceUrl"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var cmd req
		if err := xml.NewDecoder(r.Body).Decode(&cmd); err != nil {
			panic(err) // TODO: Error Handling
			return
		}

		if cmd.HeaderAction != "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/Discover" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if cmd.HeaderMessageID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("Enrollment Policy POST") // TODO: Log This Event (And All The Devices Version/Details) And Failures If They Happen
		// TODO: Send Telemetry Back To Me About Device Version/etc

		response := res{
			Envelope: soap.Envelope{
				HeaderAction:     "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/DiscoverResponse",
				HeaderActivityId: "d9eb2fdd-e38a-46ee-bd93-aea9dc86a3b8", // TODO: Generate It
				HeaderRelatesTo:  cmd.HeaderMessageID,
			},
			AuthPolicy:                 "OnPremise", // TODO: Future: Support Web Based Auth As Well
			EnrollmentVersion:          "4.0",
			EnrollmentPolicyServiceUrl: "https://mdm.otbeaumont.me/EnrollmentServer/PolicyService.svc",     // TODO: Config
			EnrollmentServiceUrl:       "https://mdm.otbeaumont.me/EnrollmentServer/EnrollmentService.svc", // TODO: Config
		}
		response.Envelope.FillEnvelopeAttrs()
		if err := xml.NewEncoder(w).Encode(response); err != nil {
			panic(err) // TODO: Error Handling
			return
		}
	}
}
