package appleMDM

// Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html

import (
	"fmt"

	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"

	"os"
)

// TODO:
//  Add Logging To File Or Something For Any Errors Occurred (Debugging For The Me)
//  See What Checkin Does If None Of The Core Values (4 Of Them) Are Not Given By The Client Does It Plist Parsing Error?
//  Detect Device That Have Disconnected From Management
//  Prevent APNS Module Form Causing "DDOS" To Apples Servers
//  Use Verify Stuff To Stop People Forging The Enrollment Profile Even If They Know The URL's
//  Switch The Order Of All Routers So HandleFunc is After Attributes

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Init() {
	fmt.Println("Started Apple MDM Module!")

	//Load/Create SCEP CA
}

func Mount(r *mux.Router) {
	r.HandleFunc("/", genericResponse).Methods("GET")
	r.HandleFunc("/enroll", enrollHandler).Methods("GET")
	r.HandleFunc("/checkin", checkinHandler).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.HandleFunc("/server", serverHandler).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
	r.HandleFunc("/scep", scepHandler).Methods("GET")
}

func genericResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Apple Mobile Device Management Server!")
}

func enrollHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-apple-aspen-config")
	http.ServeFile(w, r, "enroll.mobileconfig")
}

var lockedDevice = false //TEMP

func serverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MDM Server Request")

	buf, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		//Do something
		fmt.Println(err)
	}

	fmt.Println(string(buf))
	w.WriteHeader(http.StatusOK)
	return

	/*var cmd ServerCommand
	  if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
	    fmt.Println("Failed To Parse Checkin Request")
	    fmt.Println(err)

	    // TODO: Debug Event To Error Logs
	    w.WriteHeader(http.StatusBadRequest)
	    return
		}

	  if cmd.Status == "Idle" {
	    fmt.Println("The Device Is Idle")

	    if !lockedDevice {
	      lockedDevice = true
	      fmt.Println("Sending A Lock Command")

	      DeviceLock := struct {
	        RequestType string
	      }{
	        RequestType: "RestartDevice",
	      }

	      out, err := plist.MarshalIndent(DeviceLock, "   ")
	      if err != nil {
	        fmt.Println(err)
	      }

	      fmt.Println(string(out))
	      fmt.Fprintf(w, string(out))
	    } else {
	      fmt.Fprintf(w, "")
	    }
	  } else {
	    fmt.Fprintf(w, "")
	  }
	*/
}
