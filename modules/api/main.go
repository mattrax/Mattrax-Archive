//TODO Header

package api

import (
	"fmt"
	"net/http"
  //"io/ioutil"
  //"encoding/xml"
  //"github.com/juju/xml"
  //"regexp"

	//"strings"
	//"crypto/x509"
	//"encoding/base64"
	//"encoding/hex"

	//External Deps
	"github.com/gorilla/mux" //HTTP Router

	// Internal Functions
	mlg "github.com/mattrax/Mattrax/internal/logging" //Mattrax Logging
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database" //Mattrax Database
	//errors "github.com/mattrax/Mattrax/internal/errors" // Mattrax Error Handling
  //auth "github.com/mattrax/Mattrax/internal/authentication"

	// Internal Modules
	//restAPI "github.com/mattrax/Mattrax/windowsMDM/api" //The Windows MDM REST API
  //structs "github.com/mattrax/Mattrax/windowsMDM/structs" //The Windows MDM Structs
  //soap "github.com/mattrax/Mattrax/windowsMDM/soap" //SOAP Data Handling
)

var ( // Get The Internal State
	pgdb = mdb.GetDatabase()
	log = mlg.GetLogger()
	config = mcf.GetConfig()
)

//TODO
func Init() { log.Info("Loaded The API Module") }

//TODO Docs
func Mount(r *mux.Router) {
  //Main MDM Domain
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Windows Mobile Device Management Server!") }).Methods("GET")
  //r.HandleFunc("/enroll", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "ms-device-enrollment:?mode=mdm", 301) }).Methods("GET") //ms-device-enrollment:?mode=mdm ms-device-enrollment:?mode=mdm&username=oscar@otbeaumont.me&servername=https://mdm.otbeaumont.me", 301)
  //r.HandleFunc("/auth", authHandler).Methods("GET")
  //r.Handle("/enrollmentPolicyService", errors.Handler(enrollmentPolicyService)).Methods("POST")
	//r.Handle("/enrollmentService", errors.Handler(enrollmentService)).Methods("POST")
}
