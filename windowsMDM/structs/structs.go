/**
 * Mattrax: An Open Source Device Management System
 * File Description: TODO
 * Package Description: These Are The Structs and Helpers For Device Communication, The API and Database Communication.
 * Protcol Documentation: https://docs.microsoft.com/en-us/windows/client-management/mdm/
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package windowsMDM

//import "fmt"

/* Parse Enrollment Request */

/*
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
*/

type Envelope struct {
	Header struct {
  	Action    string        `xml:"Action"`
  	MessageID string        `xml:"MessageID"`
  	ReplyTo   ReplyToStruct `xml:"ReplyTo"`
  	To        string        `xml:"To"`
  } `xml:"Header"`
	Body   BodyStruct   `xml:"Body"`
}



type ReplyToStruct struct {
	Address string `xml:"Address"`
}

type BodyStruct struct {
	Discover DiscoverStruct `xml:"Discover"`
	//Security SecurityStruct `xml:"wsse:Security"`
}


/*
type SecurityStruct struct {
	BinarySecurityToken string `xml:"wsse:BinarySecurityToken"`
}
*/





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
/*
func init() {
	fmt.Println("Testing!")


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
	  fmt.Println(string(data))

}*/
