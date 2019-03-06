package soap

import "github.com/Zauberstuhl/go-xml"

type Envelope struct { // TODO: Check All Header Are Being Used By Multiple Things
	XMLName xml.Name `xml:"s:Envelope"`
	XmlnsA  string   `xml:"xmlns:a,attr"`
	XmlnsS  string   `xml:"xmlns:s,attr"`

	HeaderAction string `xml:"s:Header>a:Action,omitempty"`

	// Decoded Header Tags
	HeaderMessageID string `xml:"s:Header>a:MessageID,omitempty"`
	HeaderAddress   string `xml:"s:Header>a:ReplyTo,omitempty>a:Address,omitempty"`
	HeaderTo        string `xml:"s:Header>a:To,omitempty"`

	// Encoder Header Tags
	HeaderActivityId string `xml:"s:Header>ActivityId,omitempty"`
	HeaderRelatesTo  string `xml:"s:Header>a:RelatesTo,omitempty"`
}

func (h *Envelope) FillEnvelopeAttrs() {
	h.XmlnsA = "http://www.w3.org/2005/08/addressing"
	h.XmlnsS = "http://www.w3.org/2003/05/soap-envelope"
}
