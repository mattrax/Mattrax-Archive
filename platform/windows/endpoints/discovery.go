package endpoints

import (
	"log"
	"net/http"
	"strings"

	mattrax "github.com/mattrax/Mattrax/internal"

	"github.com/Zauberstuhl/go-xml"
	"github.com/mattrax/Mattrax/pkg/soap"
)

func DiscoveryGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// TODO: See How The Device Handles The Rejections - Try And Make Error The Errors Sorta Make Sense To The User On The Client
func DiscoveryPost(as mattrax.AuthService) http.HandlerFunc {
	type req struct {
		soap.Envelope
		Body struct {
			EmailAddress       string   `xml:"Discover>request>EmailAddress"`
			RequestVersion     string   `xml:"Discover>request>RequestVersion"`
			DeviceType         string   `xml:"Discover>request>DeviceType"`
			ApplicationVersion string   `xml:"Discover>request>ApplicationVersion"`
			OSEdition          string   `xml:"Discover>request>OSEdition"`
			AuthPolicies       []string `xml:"Discover>request>AuthPolicies>AuthPolicy"`
		} `xml:"s:Body"`
	}

	type res struct {
		soap.Envelope

		// TODO: Future: The s:Body doen't have The Attrs xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"
		// TODO: Future: The DiscoverResponse Doesn't Have The Attr xmlns="http://schemas.microsoft.com/windows/management/2012/01/enrollment"
		AuthPolicy                 string `xml:"s:Body>DiscoverResponse>DiscoverResult>AuthPolicy"`
		EnrollmentVersion          string `xml:"s:Body>DiscoverResponse>DiscoverResult>EnrollmentVersion"`
		EnrollmentPolicyServiceUrl string `xml:"s:Body>DiscoverResponse>DiscoverResult>EnrollmentPolicyServiceUrl"`
		EnrollmentServiceUrl       string `xml:"s:Body>DiscoverResponse>DiscoverResult>EnrollmentServiceUrl"`
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

		if cmd.Body.DeviceType != "CIMClient_Windows" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !strings.Contains(strings.Join(cmd.Body.AuthPolicies, " "), "OnPremise") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := as.VerifyEmail(cmd.Body.EmailAddress); err == mattrax.ErrInvalidEmail {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			panic(err) // TODO: Error Handling
		}

		log.Println("Discovery POST") // TODO: Log This Event (And All The Devices Version/Details) And Failures If They Happen
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
