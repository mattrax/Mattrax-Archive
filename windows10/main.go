package windows10

// Documentation: https://docs.microsoft.com/en-us/windows/client-management/mdm/federated-authentication-device-enrollment

import (
	"fmt"
	_ "os"

	"encoding/xml"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

/* Parse Enrollment Request */

type Envelope struct {
	Header HeaderStruct `xml:"Header"`
	Body   BodyStruct   `xml:"Body"`
}

type HeaderStruct struct {
	Action    string        `xml:"Action"`
	MessageID string        `xml:"MessageID"`
	ReplyTo   ReplyToStruct `xml:"ReplyTo"`
	To        string        `xml:"To"`
}

type ReplyToStruct struct {
	Address string `xml:"Address"`
}

type BodyStruct struct {
	Discover DiscoverStruct `xml:"Discover"`
}

type DiscoverStruct struct {
	Request RequestStruct `xml:"request"`
}

type RequestStruct struct {
	EmailAddress       string             `xml:"EmailAddress"`
	RequestVersion     string             `xml:"RequestVersion"`
	DeviceType         string             `xml:"DeviceType"`
	ApplicationVersion string             `xml:"ApplicationVersion"`
	OSEdition          string             `xml:"OSEdition"`
	AuthPolicies       AuthPoliciesStruct `xml:"AuthPolicies"` ////////// TODO: Make This Work
}

type AuthPoliciesStruct struct {
	AuthPolicy []string `xml:"AuthPolicy"`
}

/* Generate Enrollment Request Response */
/*
type ResponseEnvelope struct {
  Header ResponseHeaderStruct `xml:"Header"`
  //Body ResponseBodyStruct `xml:"Body"`
}

type ResponseHeaderStruct struct {
  Action string `xml:"Action"`
  ActivityId string `xml:"ActivityId"`
  RelatesTo string `xml:"RelatesTo"`
}
*/

func Init() {
	fmt.Println("Started Windows MDM Module!")

	/*
	  sd := &ResponseEnvelope{
	    Header: ResponseHeaderStruct{
	      Action: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/DiscoverResponse",
	      ActivityId: "d9eb2fdd-e38a-46ee-bd93-aea9dc86a3b8",
	      RelatesTo: "urn:uuid: 748132ec-a575-4329-b01b-6171a9cf8478",
	    },

	  }

	  data, err := xml.MarshalIndent(sd, "", "  ")
	  if err != nil {
	      panic(err)
	  }

	  fmt.Println(string(data))*/

}

func Mount(r *mux.Router, r2 *mux.Router) {
	r.HandleFunc("/", genericResponse).Methods("GET")
	r.HandleFunc("/authenticate", authenticationService).Methods("GET")

	r2.HandleFunc("/", genericResponse).Methods("GET")
	r2.HandleFunc("/EnrollmentServer/Discovery.svc", genericResponse).Methods("GET")
	r2.HandleFunc("/EnrollmentServer/Discovery.svc", enrollmentServicePOST).Methods("POST")
}

func genericResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Windows 10 Mobile Device Management Server!")
}

func enrollmentServicePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Discovery Service Hit")

	buf, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		//Do something
		fmt.Println(err)
	}

	res := &Envelope{}
	err2 := xml.Unmarshal([]byte(string(buf)), res)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Enrollment Request From Email: " + res.Body.Discover.Request.EmailAddress)
		if res.Body.Discover.Request.EmailAddress == "oscar@otbeaumont.me" {
			/*fmt.Println(res.Header.Action)
			  fmt.Println(res.Header.MessageID)
			  fmt.Println(res.Header.ReplyTo.Address)
			  fmt.Println(res.Header.To)

			  fmt.Println(res.Body.Discover.Request.EmailAddress)
			  fmt.Println(res.Body.Discover.Request.RequestVersion)
			  fmt.Println(res.Body.Discover.Request.DeviceType)
			  fmt.Println(res.Body.Discover.Request.ApplicationVersion)
			  fmt.Println(res.Body.Discover.Request.OSEdition)
			  fmt.Println(res.Body.Discover.Request.AuthPolicies)*/

			fmt.Fprintf(w, `<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope"
       xmlns:a="http://www.w3.org/2005/08/addressing">
      <s:Header>
        <a:Action s:mustUnderstand="1">
          http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/DiscoverResponse
        </a:Action>
        <ActivityId>
          d9eb2fdd-e38a-46ee-bd93-aea9dc86a3b8
        </ActivityId>
        <a:RelatesTo>urn:uuid: 748132ec-a575-4329-b01b-6171a9cf8478</a:RelatesTo>
      </s:Header>
      <s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xmlns:xsd="http://www.w3.org/2001/XMLSchema">
        <DiscoverResponse
           xmlns="http://schemas.microsoft.com/windows/management/2012/01/enrollment">
          <DiscoverResult>
            <AuthPolicy>Federated</AuthPolicy>
            <EnrollmentVersion>3.0</EnrollmentVersion>
            <EnrollmentPolicyServiceUrl>
              https://mdm.otbeaumont.me/windows10/DEVICEENROLLMENTWEBSERVICE.SVC
            </EnrollmentPolicyServiceUrl>
            <EnrollmentServiceUrl>
              https://mdm.otbeaumont.me/windows10/DEVICEENROLLMENTWEBSERVICE.SVC
            </EnrollmentServiceUrl>
            <AuthenticationServiceUrl>
              https://auth.otbeaumont.me/
            </AuthenticationServiceUrl>
          </DiscoverResult>
        </DiscoverResponse>
      </s:Body>
    </s:Envelope>`)

		} else {
			fmt.Println("Rejected Enrollment Request")
			//Reject (Using HTTP Error Response)

			fmt.Fprintf(w, "Error Occured!")
		}
	}

}

func authenticationService(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Windows Enrollment Authentication Services!")
}
