package soap

import (
	"encoding/xml"
	"io"
	"log"
)

func Parse(body io.Reader, out interface{}) error { //TODO: Change From Interface To Internal Wrapper
	decoder := xml.NewDecoder(body)

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Space == "http://www.w3.org/2003/05/soap-envelope" && se.Name.Local == "Envelope" { //TODO: Determain the XML Name Space Automatticly
				log.Println("YEP")

				decoder.DecodeElement(&out.Head, &se)

			}

			//log.Println(se.Name)
		}
	}

	/*scanner := bufio.NewScanner(body)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text())  // token in unicode-char
		fmt.Println(scanner.Bytes()) // token in bytes

	}*/

	return nil
}
