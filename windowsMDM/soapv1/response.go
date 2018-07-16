package soap

import (
	"encoding/xml"
)

type ResponseEnvelope struct {
	XMLName          xml.Name     `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	ResponseBodyBody ResponseBody `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

type ResponseBody struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	Fault   Fault    `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`
}

type Fault struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}
