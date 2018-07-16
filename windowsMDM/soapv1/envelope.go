package soap

import (
	"log"
	//"encoding/xml"
	"github.com/juju/xml"
	//"bytes"
)


type Envelope1 struct {
    XMLName xml.Name   `xml:"s:Envelope"` //http://www.w3.org/2003/05/soap-envelope
		XmlnsA	string   `xml:"xmlns:a,attr"` //EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
		XmlnsS    string   `xml:"xmlns:s,attr"`

		Header        Header2 `xml:"s:Header"` //Header2
		Body  struct {
			Payload []byte `xml:",innerxml"`
		} `xml:"s:Body"`
}

type Envelope2 struct {
    XMLName xml.Name   `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"` //http://www.w3.org/2003/05/soap-envelope
		XmlnsA	string   `xml:"xmlns:a,attr"` //EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
		XmlnsS    string   `xml:"xmlns:s,attr"`

		Header        Header2 `xml:"http://www.w3.org/2003/05/soap-envelope Header"` //Header2
		Body  struct {
			Payload []byte `xml:",innerxml"`
		} `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

type Header2 struct {
	//XMLName      xml.Name `xml:"s Header"`
	Action MustUnderstand `xml:"http://www.w3.org/2005/08/addressing Action"`
	MessageID string `xml:"http://www.w3.org/2005/08/addressing MessageID"`
	ReplyTo HeaderReplyTo `xml:"http://www.w3.org/2005/08/addressing ReplyTo"`
	To MustUnderstand `xml:"http://www.w3.org/2005/08/addressing To"`
	/*
	//XMLName      xml.Name `xml:"s:Header"`
	Action MustUnderstand //`xml:"a:Action"`
	MessageID string //`xml:"http://www.w3.org/2005/08/addressing MessageID"`
	ReplyTo HeaderReplyTo2 //`xml:"a:ReplyTo"`
	To MustUnderstand //`xml:"a:To"`
	*/
}

type HeaderReplyTo2 struct {
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
	/*AuthPolicies []struct{
		AuthPolicy []string `xml:"AuthPolicy"`
	}*/

}
















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
	Action MustUnderstand `xml:"a:Action"`
	MessageID string `xml:"a:MessageID"`
	ReplyTo HeaderReplyTo `xml:"a:ReplyTo"`
	To MustUnderstand `xml:"a:To"`
}

type MustUnderstand struct {
	MustUnderstand  int      `xml:"s:mustUnderstand,attr"`
	Payload  string        `xml:",chardata"`
}

type HeaderReplyTo struct {
	Address string `xml:"a:Address"`
}

// This has to change depending on which service we invoke
type EnvelopeBody struct {
    Payload interface{} // the empty interface lets us assign any other type
}





/*
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope"
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
			<AuthPolicy>OnPremise</AuthPolicy>
			<EnrollmentVersion>3.0</EnrollmentVersion>
			<EnrollmentPolicyServiceUrl>
				https://enrolltest.contoso.com/ENROLLMENTSERVER/DEVICEENROLLMENTWEBSERVICE.SVC
			</EnrollmentPolicyServiceUrl>
			<EnrollmentServiceUrl>
				https://enrolltest.contoso.com/ENROLLMENTSERVER/DEVICEENROLLMENTWEBSERVICE.SVC
			</EnrollmentServiceUrl>
		</DiscoverResult>
	</DiscoverResponse>
</s:Body>
</s:Envelope>
*/


/*
func init() {
	envelope := Envelope2{
		S: "http://www.w3.org/2003/05/soap-envelope",
		A: "http://www.w3.org/2005/08/addressing",
		Header: Header2{
			Action: MustUnderstand{
				MustUnderstand: 1,
				Payload: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/DiscoverResponse",
			},


		},
	}


	out, _ := xml.MarshalIndent(envelope, "", "   ") //TEMP Pretty Print
	log.Println(string(out))
}
*/

func init() { log.Println() }


/*
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
*/









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
