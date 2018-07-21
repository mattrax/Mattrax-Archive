package main

import (
  "fmt"
  "encoding/xml"
)

func main() {
  type Envelope struct {
    XMLName  xml.Name  `xml:"s:Envelope"`
    XmlNSa    string    `xml:"xmlns:a,attr"`
    XmlNSs    string    `xml:"xmlns:s,attr"`
    Data string `xml:"data"`
  }

  root := Envelope {
    XmlNSa: "http://www.w3.org/2005/08/addressing",
    XmlNSs: "http://www.w3.org/2003/05/soap-envelope",
    Data: "Testing123",
  }

  b, _ := xml.MarshalIndent(root, "", "    ")

  fmt.Println(string(b))

  back := Envelope{}
  if err := xml.Unmarshal(b, &back); err != nil { fmt.Println(err); return } //TODO Replace The Input With The Direct Stream For Preformance

  fmt.Println(back)
}
