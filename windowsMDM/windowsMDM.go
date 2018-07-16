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
  //"encoding/xml"
  "github.com/juju/xml"
  //"regexp"



	//External Deps
	"github.com/gorilla/mux" //HTTP Router

	// Internal Functions
	mlg "github.com/mattrax/Mattrax/internal/logging" //Mattrax Logging
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database" //Mattrax Database
	errors "github.com/mattrax/Mattrax/internal/errors" // Mattrax Error Handling
  //auth "github.com/mattrax/Mattrax/internal/authentication"

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
  r.HandleFunc("/auth", authHandler).Methods("GET")
  r.Handle("/enrollmentService", errors.Handler(enrollmentService)).Methods("POST")






	//REST API
	//restAPI.Mount(r.PathPrefix("/api/").Subrouter())

	// MDM Device Endpoints
	//r.Handle("/inform", errors.Handler(informHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	//r.Handle("/server", errors.Handler(serverHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
}









func enrollmentService(w http.ResponseWriter, r *http.Request) (int, error) {
  bodyBytes, _ := ioutil.ReadAll(r.Body)
  log.Info(string(bodyBytes))

  envelope := soap.Envelope{
    //XmlnsA: "http://www.w3.org/2005/08/addressing",
    //XmlnsS: "http://www.w3.org/2003/05/soap-envelope",
  }

  if err := xml.Unmarshal([]byte(string(bodyBytes)), &envelope); err != nil { return 403, err } //TODO Replace The Input With The Direct Stream For Preformance








  testing123 := &soap.GEnvelope{
    XmlnsA: "http://www.w3.org/2005/08/addressing",
    XmlnsS: "http://www.w3.org/2003/05/soap-envelope",
    Header: soap.GHeader{
      Action: soap.GMustUnderstand{
        MustUnderstand: 1,
        Payload: "http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy/IPolicy/GetPoliciesResponse",
      },
      RelatesTo: envelope.Header.MessageID,
    },
  }

  testing123.Body.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
  testing123.Body.Xsd = "http://www.w3.org/2001/XMLSchema"
  testing123.Body.Payload = []byte(`<GetPoliciesResponse
             xmlns="http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy">
            <response>
            <policyID />
              <policyFriendlyName xsi:nil="true"
                 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
              <nextUpdateHours xsi:nil="true"
                 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
              <policiesNotChanged xsi:nil="true"
                 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
              <policies>
                <policy>
                  <policyOIDReference>0</policyOIDReference>
                  <cAs xsi:nil="true" />
                  <attributes>
                    <commonName>CEPUnitTest</commonName>
                    <policySchema>3</policySchema>
                    <certificateValidity>
                      <validityPeriodSeconds>1209600</validityPeriodSeconds>
                      <renewalPeriodSeconds>172800</renewalPeriodSeconds>
                    </certificateValidity>
                    <permission>
                      <enroll>true</enroll>
                      <autoEnroll>false</autoEnroll>
                    </permission>
                    <privateKeyAttributes>
                      <minimalKeyLength>2048</minimalKeyLength>
                      <keySpec xsi:nil="true" />
                      <keyUsageProperty xsi:nil="true" />
                      <permissions xsi:nil="true" />
                      <algorithmOIDReference xsi:nil="true" />
                      <cryptoProviders xsi:nil="true" />
                    </privateKeyAttributes>
                    <revision>
                      <majorRevision>101</majorRevision>
                      <minorRevision>0</minorRevision>
                    </revision>
                    <supersededPolicies xsi:nil="true" />
                    <privateKeyFlags xsi:nil="true" />
                    <subjectNameFlags xsi:nil="true" />
                    <enrollmentFlags xsi:nil="true" />
                    <generalFlags xsi:nil="true" />
                    <hashAlgorithmOIDReference>0</hashAlgorithmOIDReference>
                    <rARequirements xsi:nil="true" />
                    <keyArchivalAttributes xsi:nil="true" />
                    <extensions xsi:nil="true" />
                  </attributes>
                </policy>
              </policies>
            </response>
            <cAs xsi:nil="true" />
            <oIDs>
              <oID>
                <value>1.3.14.3.2.29</value>
                <group>1</group>
                <oIDReferenceID>0</oIDReferenceID>
                <defaultName>szOID_OIWSEC_sha1RSASign</defaultName>
              </oID>
            </oIDs>
          </GetPoliciesResponse>`)

  //testing123.Body =

  out, _ := xml.MarshalIndent(testing123, "", "   ") //TEMP Pretty Print
	log.Info(string(out))
  //fmt.Fprintf(w, string(out))

  fmt.Fprintf(w, `<s:Envelope
   xmlns:a="http://www.w3.org/2005/08/addressing"
   xmlns:s="http://www.w3.org/2003/05/soap-envelope">
   <s:Header>
     <a:Action s:mustUnderstand="1">
 http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy/IPolicy/GetPoliciesResponse
     </a:Action>
   </s:Header>
   <s:Body
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xmlns:xsd="http://www.w3.org/2001/XMLSchema">
     <GetPoliciesResponse   xmlns="http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy">
       <response>
         <policies>
           <policy>
             <attributes>
               <policySchema>3</policySchema>
               <privateKeyAttributes>
                 <minimalKeyLength>2048</minimalKeyLength>
                 <algorithmOIDReferencexsi:nil="true"/>
               </privateKeyAttributes>
               <hashAlgorithmOIDReference xsi:nil="true"></hashAlgorithmOIDReference>
             </attributes>
           </policy>
         </policies>
       </response>
       <oIDs>
         <oID>
           <value>1.3.6.1.4.1.311.20.2</value>
           <group>1</group>
           <oIDReferenceID>5</oIDReferenceID>
           <defaultName>Certificate Template Name</defaultName>
         </oID>
       </oIDs>
     </GetPoliciesResponse>
   </s:Body>
 </s:Envelope>`)

  return 200, nil
}












func authHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprintf(w, `<!DOCTYPE>
  <html>
     <head>
        <title>Working...</title>
        <script>
           function formSubmit() {
              document.forms[0].submit();
           }
           window.onload=formSubmit;
        </script>
     </head>
     <body>
       <h1>Mattrax authentication</h1>
      <!-- appid below in post command must be same as appid in previous client https request. -->
        <form method="post" action="` + r.URL.Query().Get("appru") + `">
           <p><input name="wresult" value="` + r.URL.Query().Get("appru") + `"/></p>
           <input type="submit"/>
        </form>
     </body>
  </html>
`)
}


//TODO: Redo Error Returns
// TODO: Doc
func enrollmentDiscover(w http.ResponseWriter, r *http.Request) (int, error) {
  bodyBytes, _ := ioutil.ReadAll(r.Body)
  //log.Info(string(bodyBytes))

  /*bodyBytes, _ := ioutil.ReadAll(r.Body)
  log.Info(string(bodyBytes))
  return 200, nil*/

  envelope := soap.Envelope{
    //XmlnsA: "http://www.w3.org/2005/08/addressing",
    //XmlnsS: "http://www.w3.org/2003/05/soap-envelope",
  }

  if err := xml.Unmarshal([]byte(string(bodyBytes)), &envelope); err != nil { return 403, err } //TODO Replace The Input With The Direct Stream For Preformance

  //if err := xml.NewDecoder(r.Body).Decode(&envelope); err != nil { return 403, err }
  //cmd := soap.DiscoverPayload{}
  //if err := xml.Unmarshal(envelope.Body.Payload, &cmd); err != nil { return 403, err }


  //log.Info(env)

  //log.Warning(envelope.Header.Action.Payload)
  log.Warning(envelope.Header.MessageID)
  //log.Warning(envelope.Header.ReplyTo.Address)
  //log.Warning(envelope.Header.To.Payload)
  //log.Warning(cmd.Request.EmailAddress)


  /*log.Warning(env.Header.Action.Payload)
  log.Warning(env.Header.MessageID)
  log.Warning(env.Header.ReplyTo.Address)
  log.Warning(env.Header.To.Payload)
  log.Warning(env.Header.To.AuthPolicies)
  */



  //log.Warning(string(env.Body.Payload))

  //log.Warning(cmd.Request.EmailAddress) //Could Not be Email But Domain\Username


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
        <AuthPolicy>Federated</AuthPolicy>
        <EnrollmentVersion>3.0</EnrollmentVersion>
        <EnrollmentPolicyServiceUrl>
          https://enrolltest.contoso.com/ENROLLMENTSERVER/DEVICEENROLLMENTWEBSERVICE.SVC
        </EnrollmentPolicyServiceUrl>
        <EnrollmentServiceUrl>
          https://enrolltest.contoso.com/ENROLLMENTSERVER/DEVICEENROLLMENTWEBSERVICE.SVC
        </EnrollmentServiceUrl>
        <AuthenticationServiceUrl>
          https://portal.manage.contoso.com/LoginRedirect.aspx
        </AuthenticationServiceUrl>
      </DiscoverResult>
    </DiscoverResponse>
  </s:Body>
</s:Envelope>
*/

  testing123 := &soap.GEnvelope{
    XmlnsA: "http://www.w3.org/2005/08/addressing",
    XmlnsS: "http://www.w3.org/2003/05/soap-envelope",
    Header: soap.GHeader{
      Action: soap.GMustUnderstand{
        MustUnderstand: 1,
        Payload: "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/DiscoverResponse",
      },
      ActivityId: "d9eb2fdd-e38a-46ee-bd93-aea9dc86a3b8",
      RelatesTo: envelope.Header.MessageID,
    },
    /*Body: &soap.GTesting{
Payload: `<DiscoverResponse
   xmlns="http://schemas.microsoft.com/windows/management/2012/01/enrollment">
  <DiscoverResult>
    <AuthPolicy>Federated</AuthPolicy>
    <EnrollmentVersion>3.0</EnrollmentVersion>
    <EnrollmentPolicyServiceUrl>
      https://enrolltest.contoso.com/ENROLLMENTSERVER/DEVICEENROLLMENTWEBSERVICE.SVC
    </EnrollmentPolicyServiceUrl>
    <EnrollmentServiceUrl>
      https://enrolltest.contoso.com/ENROLLMENTSERVER/DEVICEENROLLMENTWEBSERVICE.SVC
    </EnrollmentServiceUrl>
    <AuthenticationServiceUrl>
      https://portal.manage.contoso.com/LoginRedirect.aspx
    </AuthenticationServiceUrl>
  </DiscoverResult>
  </DiscoverResponse>`, },*/
  }

  testing123.Body.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
  testing123.Body.Xsd = "http://www.w3.org/2001/XMLSchema"
  testing123.Body.Payload = []byte(`<DiscoverResponse
           xmlns="http://schemas.microsoft.com/windows/management/2012/01/enrollment">
          <DiscoverResult>
            <AuthPolicy>Federated</AuthPolicy>
            <EnrollmentVersion>3.0</EnrollmentVersion>
            <EnrollmentPolicyServiceUrl>
              https://mdm.otbeaumont.me/windows/enrollmentService?policies
            </EnrollmentPolicyServiceUrl>
            <EnrollmentServiceUrl>
              https://mdm.otbeaumont.me/windows/enrollmentService
            </EnrollmentServiceUrl>
            <AuthenticationServiceUrl>
              https://` + config.Domain + `/windows/auth
            </AuthenticationServiceUrl>
          </DiscoverResult>
        </DiscoverResponse>`)

  //testing123.Body =

  out, _ := xml.MarshalIndent(testing123, "", "   ") //TEMP Pretty Print
	log.Info(string(out))
  fmt.Fprintf(w, string(out))
  //Send



  //Verify The To In The Headers Is Correct

  //TODO: Check User Maybe (Is There Response For It)
  //      Maybe Only Check The Email Is Of Correct Domains
  //TODO: Check AuthPolicies Are In The Thing Before Using Them

  //if auth.CheckUser
  //log.Println(cmd)





  return 200, nil
}































//TODO:
//  Add Thanks For The SOAP Library In Every Windows MDM File
// Note The HTTP server response must not be chunked; it must be sent as one message.
