package main

import (
  "fmt"

  "github.com/achiku/xml"
  soapc "github.com/mattrax/Mattrax/windowsMDM/soapc"
  //soapc "github.com/achiku/soapc"
)

type name struct {
	XMLName xml.Name `xml:"name"`
	Testing   string   `xml:"first,omitempty"`
}

type header struct {
	XMLName xml.Name `xml:"name"`
	Testing   string   `xml:"headerItem,omitempty"`
}

func main() {
  v := soapc.Envelope{
    Header: &soapc.Header{
      Content: &header{
        Testing: "Hello World123",
      },
    },
		Body: soapc.Body{
      Content: &name{
        Testing: "Hello World",
      },
		},
	}

  out, _ := xml.MarshalIndent(v, "", "  ")
  fmt.Println(string(out))

  fmt.Println("")

  envelope := soapc.Envelope{
    Header: &soapc.Header{
      Content: &header{
        //Testing: "Hello World123",
      },
    },
		Body: soapc.Body{
      Content: &name{
        //Testing: "Hello World",
      },
		},
	}

  if err := xml.Unmarshal(out, &envelope); err != nil { fmt.Println(err); return } //TODO Replace The Input With The Direct Stream For Preformance

  //xml.MarshalIndent(envelope, "", "  ")
  fmt.Println(envelope)
}
