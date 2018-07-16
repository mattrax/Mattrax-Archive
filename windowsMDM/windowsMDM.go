/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is The Windows MDM Core. It Manages The Webserver Routes // TODO: Windows APSN Thingo Here and APNS for Apples MDM.
 * Protcol Documentation: https://docs.microsoft.com/en-us/windows/client-management/mdm/
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package windowsMDM

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "encoding/xml"
  //"regexp"



	//External Deps
	"github.com/gorilla/mux" //HTTP Router

	// Internal Functions
	mlg "github.com/mattrax/Mattrax/internal/logging" //Mattrax Logging
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database" //Mattrax Database
	errors "github.com/mattrax/Mattrax/internal/errors" // Mattrax Error Handling

	// Internal Modules
	//restAPI "github.com/mattrax/Mattrax/windowsMDM/api" //The Windows MDM REST API
  //structs "github.com/mattrax/Mattrax/windowsMDM/structs" //The Windows MDM Structs
  soap "github.com/mattrax/Mattrax/windowsMDM/soap" //SOAP Data Handling
)

var ( // Get The Internal State
	pgdb = mdb.GetDatabase()
	log = mlg.GetLogger()
	config = mcf.GetConfig()
)

//TODO
func Init() { log.Info("Loaded The Windows MDM Module") }

//TODO Docs
func Mount(r *mux.Router, ee *mux.Router) {
  //Custom Discovery Domain
  ee.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Windows Mobile Device Management Server!") }).Methods("GET")
  ee.HandleFunc("/EnrollmentServer/Discovery.svc", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "") }).Methods("GET")
  ee.Handle("/EnrollmentServer/Discovery.svc", errors.Handler(enrollmentDiscover)).Methods("POST")

  //Main MDM Domain
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Windows Mobile Device Management Server!") }).Methods("GET")
  r.HandleFunc("/enroll", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "ms-device-enrollment:?mode=mdm", 301) }).Methods("GET") //ms-device-enrollment:?mode=mdm ms-device-enrollment:?mode=mdm&username=oscar@otbeaumont.me&servername=https://mdm.otbeaumont.me", 301)





	//REST API
	//restAPI.Mount(r.PathPrefix("/api/").Subrouter())

	// MDM Device Endpoints
	//r.Handle("/inform", errors.Handler(informHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	//r.Handle("/server", errors.Handler(serverHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
}





//TODO: Redo Error Returns
// TODO: Doc
func enrollmentDiscover(w http.ResponseWriter, r *http.Request) (int, error) {
  bodyBytes, _ := ioutil.ReadAll(r.Body)
  log.Info(string(bodyBytes))

  cmd := &soap.Envelope{ Body: soap.EnvelopeBody{ Payload: soap.DiscoverPayload{} } }
  //cmd.Body =


  if err := xml.Unmarshal([]byte(string(bodyBytes)), cmd); err != nil { return 403, err } //TODO Replace The Input With The Direct Stream For Preformance

  log.Warning(cmd)







  return 200, nil



  //Parse The Request
  //cmd := &structs.Envelope{}
  //cmd := &soap.ResponseEnvelope{}
  //if err := xml.Unmarshal([]byte(string(bodyBytes)), cmd); err != nil { return 403, err } //TODO Replace The Input With The Direct Stream For Preformance

  /*if cmd.Header.Action != "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/Discover" {
    return 403, errors.New("The Device Is Not Discovering ???")
  }

  if !regexp.MustCompile("^[A-Za-z0-9._%+-]+@otbeaumont.me$").MatchString(cmd.Body.Discover.Request.EmailAddress) { // TODO: Replace With Check Users In The Database/Active Directory
    return 403, errors.New("The Device's Email Is Invalid/Not Of The Correct Domain")
  }*/

  /*soapResponse, _ := soap.SoapFomMTOM(bodyBytes)
  cmd := soap.ResponseEnvelope{}
	_ = xml.Unmarshal(soapResponse, &cmd)
  */


  data := `<?xml version="1.0"?>
<s:Envelope xmlns:a="http://www.w3.org/2005/08/addressing" xmlns:s="http://www.w3.org/2003/05/soap-envelope">
<s:Header>
    <a:Action s:mustUnderstand="1">http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/Discover</a:Action>
    <a:MessageID>urn:uuid:748132ec-a575-4329-b01b-6171a9cf8478</a:MessageID>
    <a:ReplyTo>
        <a:Address>http://www.w3.org/2005/08/addressing/anonymous</a:Address>
    </a:ReplyTo>
    <a:To s:mustUnderstand="1">https://EnterpriseEnrollment.otbeaumont.me:443/EnrollmentServer/Discovery.svc</a:To>
</s:Header>


<s:Body>
    <Discover xmlns="http://schemas.microsoft.com/windows/management/2012/01/enrollment">
    <request xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
    <EmailAddress>oscar@otbeaumont.me</EmailAddress>
    <RequestVersion>4.0</RequestVersion>
    <DeviceType>CIMClient_Windows</DeviceType>
    <ApplicationVersion>10.0.17134.0</ApplicationVersion>
    <OSEdition>48</OSEdition>
    <AuthPolicies>
        <AuthPolicy>OnPremise</AuthPolicy>
        <AuthPolicy>Federated</AuthPolicy>
    </AuthPolicies>
</request>
</Discover>
</s:Body>
</s:Envelope>`


  //log.Info(soap.CheckFault([]byte(data)))

  s := &soap.ResponseEnvelope{}
  if err := xml.Unmarshal([]byte(data), s); err != nil { return 403, err } //TODO Replace The Input With The Direct Stream For Preformance




  log.Info("Done")
  log.Info(s)
  log.Printf("%+v\n", s)






















  //MessageID := "urn:uuid: 748132ec-a575-4329-b01b-6171a9cf8478"//Get From The Input

  /*body := &soap.Discover{
    Xmlns: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/",
    Request: soap.DiscoverRequest {
      Address: "Hello World",
    },
  }*/


  /*env := &soap.Envelope{
    XmlnsA: "http://www.w3.org/2005/08/addressing",
    XmlnsS: "http://www.w3.org/2003/05/soap-envelope",
    Header: &soap.Header{
      Action: soap.MustUnderstand{
        MustUnderstand: 1,
        Value: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/Discover",
      },
      MessageID: MessageID,
      ReplyTo: soap.ReplyTo{
        Address: "http://www.w3.org/2005/08/addressing/anonymous",
      },
      To: soap.MustUnderstand{
        MustUnderstand: 1,
        Value: "https://mdm.otbeaumont.me/EnrollmentServer/Discovery.svc",
      }, //TODO: This Or The Old Domain ????
    },
    Body: &soap.Body{
      Payload: &soap.Discover{
        Xmlns: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/",
        MessageID: "bo",
        //Request: soap.DiscoverRequest {
        //  Address: "Hello World",
        //
        //},
      },
    },
	}



	//env.Header.WsseSecurity.UsernameToken.Username.Value = "username"
	//env.Header.WsseSecurity.UsernameToken.Password.Value = "pass"
	//env.Body = &soap.Body{} // interface

	output, err := xml.MarshalIndent(env, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))*/






  /*
  fmt.Fprintf(w, ` <?xml version="1.0"?>
    <s:Envelope xmlns:a="http://www.w3.org/2005/08/addressing"
       xmlns:s="http://www.w3.org/2003/05/soap-envelope">
      <s:Header>
        <a:Action s:mustUnderstand="1">
          http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/Discover
        </a:Action>
        <a:MessageID>urn:uuid: 748132ec-a575-4329-b01b-6171a9cf8478</a:MessageID>
        <a:ReplyTo>
          <a:Address>http://www.w3.org/2005/08/addressing/anonymous</a:Address>
        </a:ReplyTo>
        <a:To s:mustUnderstand="1">
          https://mdm.otbeaumont.me/EnrollmentServer/Discovery.svc
        </a:To>
      </s:Header>
      <s:Body>
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
      </s:Body>
    </s:Envelope>`)
    */


  return 200, nil
}


//TODO:
//  Add Thanks For The SOAP Library In Every Windows MDM File
