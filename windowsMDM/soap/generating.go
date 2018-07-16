package soap

import (
	"encoding/xml"
)

//XmlnsA	string   `xml:"xmlns:a,attr"` //EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
//XmlnsS    string   `xml:"xmlns:s,attr"`
 //http://www.w3.org/2003/05/soap-envelope

type GEnvelope struct {
    XMLName xml.Name   `xml:"s:Envelope"`
    XmlnsA	string   `xml:"xmlns:a,attr"`
		XmlnsS    string   `xml:"xmlns:s,attr"`

		Header        GHeader `xml:"s:Header"`
		Body  struct {
			Payload []byte `xml:",innerxml"`

      Xsi string `xml:"xmlns:xsi,attr,omitempty"`
      Xsd string `xml:"xmlns:xsd,attr,omitempty"`
		} `xml:"s:Body"`
}


type GHeader struct {
	Action GMustUnderstand `xml:"a:Action,omitempty"`
	MessageID string `xml:"a:MessageID,omitempty"`
	ReplyTo GHeaderReplyTo `xml:"a:ReplyTo,omitempty"`
	To GMustUnderstand `xml:"a:To,omitempty"`

  ActivityId string `xml:"ActivityId,omitempty"`
  RelatesTo string `xml:"a:RelatesTo,omitempty"`
}

type GHeaderReplyTo struct {
	Address string `xml:"a:Address,omitempty"`
}

//Payloads

type GDiscoverPayload struct {
    XMLName xml.Name `xml:"Discover"`
		Xmlns string `xml:"xmlns,attr"`
		Request Grequest `xml:"request"`
}

type Grequest struct {
	I string `xml:"xmlns:i,attr"`
	EmailAddress string
	OSEdition string
	RequestVersion string
	DeviceType string
	ApplicationVersion string
	AuthPolicies      []string `xml:"AuthPolicies>AuthPolicy"`

}

type GMustUnderstand struct {
	MustUnderstand  int      `xml:"s:mustUnderstand,attr"`
	Payload  string        `xml:",chardata"`
}



type GTesting struct {
  Payload  string        `xml:",chardata"`
}
