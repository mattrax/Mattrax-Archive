package main

import (
	"encoding/xml"
	"fmt"

	"github.com/alexrsagen/go-libxml"
)

type Envelope struct {
	XMLName   xml.Name `xml:"s:Envelope"`
  Attrs   []xml.Attr `xml:"-"`



  //XmlNSa    string    `xml:"xmlns:a,attr"` //Try Putting Attributes In Dynamic Map
  //XmlNSs    string    `xml:"xmlns:s,attr"`

  Header  *Header  `xml:",omitempty"`
  Body
  //Body string
	//Body    Body
  //Body interface{} `xml:"s:Body"`
}

type Body struct {
  XMLName xml.Name    `xml:"s:Body"`
	Testing string `xml:"s:Testing"`
}



// Header header
type Header struct {
	XMLName xml.Name    `xml:"s:Header"`
	Content interface{} `xml:",omitempty"`
}

// Body body
/*
type Body struct {
	XMLName xml.Name    `xml:"s:Body"`
	//Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}
*/


type TestingBody struct {
  Hello string //`xml:",omitempty"`
}

func main() {
  atrributes := make([]xml.Attr, 1)
  atrributes[0] = xml.Attr{
    Name: xml.Name{
      //Space
      Local: "test",
    },
    Value: "What is up",
  }

	v := &Envelope{
    Attrs: atrributes, /*[]xml.Attr{
        xml.Attr{
          Name: "test",
          Value: "What is up",
        },
      },*/



    //XmlNSa: "http://www.w3.org/2005/08/addressing",
    //XmlNSs: "http://www.w3.org/2003/05/soap-envelope",

    /*Body: "Hey", Body{
      Content: TestingBody{
        Hello: "Hello World",
      },
    },*/
  }
  v.Body = Body{
    Testing: "World",
  }

    //FirstName: "John", LastName: "Doe", Age: 42}
	//v.address = address{"Hanga Roa", "Easter Island"}

	output, err := libxml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("output: %s\n", output)



  //Body:  Body{ Content: TestingBody{} }
  v2 := &Envelope{  } //address: address{}

	err2 := libxml.Unmarshal(output, v2)
	if err2 != nil {
		fmt.Printf("error: %v\n", err2)
		return
	}

	fmt.Printf("v: %v\n", v2)

  fmt.Println(v2.Body.Testing)
}
