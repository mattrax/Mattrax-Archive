package appleMDM

import (
	"fmt"
	"net/http"

	//"github.com/micromdm/scep/scep"
	"bytes"
	"crypto/x509"
	"github.com/fullsailor/pkcs7"
	//"crypto/rsa"
	"encoding/pem"
	"io/ioutil"
	"os"
)

const (
	scep_cert        = "scepCA.crt"
	scep_cert_binary = "scepCA.der" // openssl x509 -in rootCA.pem -outform der -out rootCA.der
	scep_key         = "scepCA.key"
)

func DegenerateCertificates(cert *x509.Certificate) ([]byte, error) {
	var buf bytes.Buffer
	/*for _, cert := range certs {
		buf.Write(cert.Raw)
	}*/
	buf.Write(cert.Raw)
	degenerate, err := pkcs7.DegenerateCertificate(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return degenerate, nil
}

func scepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SCEP Request")

	operation := r.URL.Query().Get("operation")

	if operation == "GetCACert" { //TODO: Make This Encode To Binary On Strtup For Preformance Savings
		fmt.Println("SCEP GetCACert")
		w.Header().Set("Content-Type", "application/x-x509-ca-cert")

		/*raw_cert, err := ioutil.ReadFile("ca.pem")
		    if err != nil {
		      fmt.Println(err)
		  		return
		  	}*/

		cert := LoadX509KeyPair("ca.pem", "ca.key")

		out, err := DegenerateCertificates(cert)
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println(out)

		fmt.Fprintf(w, string(out))

		//http.ServeFile(w, r, scep_cert_binary)
		return
	} else if operation == "GetCACaps" {
		fmt.Println("SCEP GetCACaps")
		fmt.Fprintf(w, "") //TODO: Implementate This In The Future (For Now It Works)
	} else if operation == "PKIOperation" {

		fmt.Println("SCEP PKIOperation")
		//body := r.URL.Query().Get("message")

		/*data := body

		    if len(data) == 0 {
		      fmt.Println("pkcs7: input data is empty")
		      return
		  		//return nil, errors.New("pkcs7: input data is empty")
		  	}
		  	var info contentInfo
		  	der, err := ber2der(data)
		  	if err != nil {
		      fmt.Println(err)
		  		return
		  	}
		  	rest, err := asn1.Unmarshal(der, &info)
		  	if len(rest) > 0 {
		  		err = asn1.SyntaxError{Msg: "trailing data"}
		  		return
		  	}
		  	if err != nil {
		  		return
		  	}*/

		// fmt.Printf("--> Content Type: %s", info.ContentType)
		/*switch {
		case info.ContentType.Equal(oidSignedData):
			return parseSignedData(info.Content.Bytes)
		case info.ContentType.Equal(oidEnvelopedData):
			return parseEnvelopedData(info.Content.Bytes)
		}*/
		//return nil, ErrUnsupportedContentType

		/*p7, err := pkcs7.Parse([]byte(body))
		  	if err != nil {
		      fmt.Println("---------------------------------")
		      fmt.Println("Error:")
		      fmt.Println(err)
		  		return
		  	}

		    fmt.Println(p7.Content)*/
		/*decoded_body, err := url.QueryUnescape(body) //base64.StdEncoding.DecodeString(body)
		    if err != nil {
		      fmt.Println(err)
		  		return
		  	}

		    fmt.Println("")
		    fmt.Println("")
		    fmt.Println(decoded_body)
		    fmt.Println("")
		    fmt.Println("")

		    //var StdEncoding = base64.NewEncoding("ASCII")
		    decoded_body2, err2 := base64.StdEncoding.DecodeString(decoded_body) //base64.URLEncoding.EncodeToString([]byte(body)) //base64.StdEncoding.DecodeString(decoded_body)
		    if err2 != nil {
		      fmt.Println(err2)
		  		return
		  	}

		    fmt.Println("")
		    fmt.Println("")
		    fmt.Println(decoded_body2)
		    fmt.Println("")
		    fmt.Println("")*/

		/*
		   if(!req.query.message) {
		     console.log("The Client Sent A Blank Payload. SCEP Failed");
		     res.status(400).end();
		   }
		*/

		/*p7, err3 := pkcs7.Parse([]byte(body))  //pkcs7.Parse([]byte(body))
		  	if err3 != nil {
		      fmt.Println(err3)
		  		return
		  	}

		    fmt.Println(p7)*/

	} else {
		fmt.Println("Unkown Operation: " + operation)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func LoadX509KeyPair(certFile, keyFile string) *x509.Certificate { //, *rsa.PrivateKey) {
	cf, e := ioutil.ReadFile(certFile)
	if e != nil {
		fmt.Println("cfload:", e.Error())
		os.Exit(1)
	}

	cpb, _ := pem.Decode(cf) // _ = cr
	//fmt.Println(string(cr))
	crt, e := x509.ParseCertificate(cpb.Bytes)

	/*kf, e := ioutil.ReadFile(keyFile)
	  if e != nil {
	      fmt.Println("kfload:", e.Error())
	      os.Exit(1)
	  }

	  fmt.Println(string(cr))
	  kpb, kr := pem.Decode(kf)
	  fmt.Println(string(kr))


	  if e != nil {
	      fmt.Println("parsex509:", e.Error())
	      os.Exit(1)
	  }
	  key, e := x509.ParsePKCS1PrivateKey(kpb.Bytes)
	  if e != nil {
	      fmt.Println("parsekey:", e.Error())
	      os.Exit(1)
	  }*/
	return crt //, key
}
