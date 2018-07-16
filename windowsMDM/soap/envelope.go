package soap

import (
	"log"
	"encoding/xml"
	//"bytes"
)

type Envelope struct {
    XMLName       xml.Name       `xml:"Envelope"`

    //EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
		A	string   `xml:"xmlns:a,attr"`
		S    string   `xml:"xmlns:s,attr"`


    Header        EnvelopeHeader `xml:"s:Header"`
    Body          EnvelopeBody   `xml:"s:Body"`
}

// This is fixed independently of whatever port we call
/*
type EnvelopeHeader struct {
    Credentials string `xml:"http://esb/definitions ESBCredentials"`
}*/
type EnvelopeHeader struct {
	XMLName      xml.Name `xml:"s:Header"`
	Action interface{} `xml:"a:Action"`
	MessageID string `xml:"a:MessageID"`
	ReplyTo HeaderReplyTo `xml:"a:ReplyTo"`
	To interface{} `xml:"a:To"`
}

type MustUnderstand struct {
	MustUnderstand  int      `xml:"s:mustUnderstand,attr"`
	Payload  interface{}        `xml:",chardata"`
}

type HeaderReplyTo struct {
	Address string `xml:"a:Address"`
}

// This has to change depending on which service we invoke
type EnvelopeBody struct {
    Payload interface{} // the empty interface lets us assign any other type
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
	/*AuthPolicies []struct{
		AuthPolicy []string `xml:"AuthPolicy"`
	}*/

}





func CreatePayload(Action string, MessageID string, ReplyTo string, To string) *Envelope {
    // Build the envelope (this could be farmed out to another func)
    env := &Envelope{
			A: "http://www.w3.org/2005/08/addressing",
			S: "http://www.w3.org/2003/05/soap-envelope",
		}

		env.Header = EnvelopeHeader{
			Action: MustUnderstand{
				MustUnderstand: 1,
				Payload: Action,
			},
			MessageID: MessageID,
			ReplyTo: HeaderReplyTo{
				Address: ReplyTo,
			},
			To: MustUnderstand{
				MustUnderstand: 1,
				Payload: To,
			},
		}

    return env
}


func init() {
	//buffer := &bytes.Buffer{}
	//encoder := xml.NewEncoder(buffer)
	//Decalare Envelope
	//_ = encoder.Encode(envelope) //TODO Error Handling


	payload := DiscoverPayload{ Xmlns: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/" }
	payload.Request.I = "http://www.w3.org/2001/XMLSchema-instance"
	payload.Request.EmailAddress = "Hello World"
	payload.Request.OSEdition = "Hello World"
	payload.Request.RequestVersion = "Hello World"
	payload.Request.DeviceType = "Hello World"
	payload.Request.ApplicationVersion = "Hello World"
	//payload.Request.AuthPolicies.AuthPolicy = []string{ "OnPremise" }

	envelope := CreatePayload("http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/Discover",
														"Testing Message ID", "http://www.w3.org/2005/08/addressing/anonymous",
														"https://mdm.otbeaumont.me/EnrollmentServer/Discovery.svc")
	envelope.Body = EnvelopeBody{ Payload: payload }



	out, _ := xml.MarshalIndent(envelope, "", "   ") //TEMP Pretty Print
	log.Println(string(out))
	//log.Println(buffer)
}










/*

type Envelope struct {
	XMLName      xml.Name `xml:"s:Envelope"`
	XmlnsA	string   `xml:"xmlns:a,attr"`
	XmlnsS    string   `xml:"xmlns:s,attr"`

	Header *Header
	Body *Body
}

type Header struct {
	XMLName      xml.Name `xml:"s:Header"`
	Action interface{} `xml:"a:Action"`
	MessageID string `xml:"a:MessageID"`
	ReplyTo ReplyTo `xml:"a:ReplyTo"`
	To interface{} `xml:"a:To"`
}

type MustUnderstand struct {
	Value  string        `xml:",chardata"`
	MustUnderstand  int      `xml:"s:mustUnderstand,attr"`
}

// TODO: Docs For This File
type ReplyTo struct {
	Address string `xml:"a:Address"`
}

type Body struct {
	XMLName xml.Name `xml:"s:Body"`
	Payload interface{} `xml:",chardata"`
}


*/
/* Different Bodys */

/*
<Discover xmlns="http://schemas.microsoft.com/windows/management/2012/01/enrollment/">
	<request xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
		<EmailAddress>user@contoso.com</EmailAddress>
		<OSEdition>3</OSEdition> <!--New -->
		<RequestVersion>3.0</RequestVersion> <!-- Updated -->
		<DeviceType>WindowsPhone</DeviceType> <!--Updated -->
		<ApplicationVersion>10.0.0.0</ApplicationVersion>
		<AuthPolicies>
			 <AuthPolicy>OnPremise</AuthPolicy>
		</AuthPolicies>
	</request>
</Discover>
*/
/*
type Discover struct {
	XMLName      xml.Name `xml:"Discover"`
	Xmlns string `xml:"xmlns,attr"`
	MessageID string `xml:"a:MessageID"`
}
*/

/*
type Discover struct {
	XMLName xml.Name `xml:"Discover"`
	Xmlns string `xml:"xmlns,attr"`
	Value string `xml:",chardata"`
	//Request DiscoverRequest `xml:",chardata"`
}

type DiscoverRequest struct {
	XMLName xml.Name `xml:"Request"`
	Address string `xml:"a:Address"`
}
*/









/*
type Header struct {
	XMLName      xml.Name `xml:"soapenv:Header"`
	WsseSecurity *WsseSecurity
}
*/

/*
type WsseSecurity struct {
	MustUnderstand string   `xml:"soapenv:mustUnderstand,attr"`
	XMLName        xml.Name `xml:"wsse:Security"`
	XmlnsWsse      string   `xml:"xmlns:wsse,attr"`
	XmlnsWsu       string   `xml:"xmlns:wsu,attr"`

	UsernameToken *UsernameToken
}

type UsernameToken struct {
	XMLName  xml.Name `xml:"wsse:UsernameToken"`
	WsuId    string   `xml:"wsu:Id,attr,omitempty"`
	Username *Username
	Password *Password
	Nonce    *Nonce
	Created  *Created
}

type Username struct {
	XMLName xml.Name `xml:"wsse:Username"`
	Value   string   `xml:",chardata"`
}
type Password struct {
	XMLName xml.Name `xml:"wsse:Password"`
	Type    string   `xml:"Type,attr"`
	Value   string   `xml:",chardata"`
}
type Nonce struct {
	XMLName      xml.Name `xml:"wsse:Nonce,omitempty"`
	EncodingType string   `xml:"EncodingType,attr,omitempty"`
	Value        string   `xml:",chardata"`
}
type Created struct {
	XMLName xml.Name `xml:"wsu:Created,omitempty"`
	Value   string   `xml:",chardata"`
}
*/
