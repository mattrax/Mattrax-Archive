package soap

import (
	"encoding/xml"
)

//XmlnsA	string   `xml:"xmlns:a,attr"` //EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
//XmlnsS    string   `xml:"xmlns:s,attr"`
 //http://www.w3.org/2003/05/soap-envelope

type Envelope struct {
    XMLName xml.Name   `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`

		Header        Header `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
		Body  struct {
			Payload []byte `xml:",innerxml"`
		} `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}


type Header struct {
	Action MustUnderstand `xml:"http://www.w3.org/2005/08/addressing Action"`
	MessageID string `xml:"http://www.w3.org/2005/08/addressing MessageID"`
	ReplyTo HeaderReplyTo `xml:"http://www.w3.org/2005/08/addressing ReplyTo"`
	To MustUnderstand `xml:"http://www.w3.org/2005/08/addressing To"`
	Security HeaderSecurity `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`

}

type HeaderSecurity struct {
	BinarySecurityToken string `"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd BinarySecurityToken"`
}


type HeaderReplyTo struct {
	Address string `xml:"http://www.w3.org/2005/08/addressing Address"`
}

//Payloads

type DiscoverPayload struct {
  XMLName xml.Name `xml:"Discover"`
	Xmlns string `xml:"xmlns,attr"`
	Request request `xml:"request"`
}

type request struct {
	I string `xml:"xmlns:i,attr"`
	EmailAddress string
	OSEdition string
	RequestVersion string
	DeviceType string
	ApplicationVersion string
	AuthPolicies      []string `xml:"AuthPolicies>AuthPolicy"`

}

type MustUnderstand struct {
	MustUnderstand  int      `xml:"http://www.w3.org/2003/05/soap-envelope mustUnderstand,attr"`
	Payload  string        `xml:",chardata"`
}





type SecurityBody struct {
	RequestSecurityToken Security2Body `xml:"http://docs.oasis-open.org/ws-sx/ws-trust/200512 wst:RequestSecurityToken"`
}

type Security2Body struct {
	BinarySecurityToken string `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd wsse:BinarySecurityToken"`
}


//RequestSecurityToken

// wsse:BinarySecurityToken
