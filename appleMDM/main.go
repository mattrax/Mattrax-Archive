/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Apple MDM Core. It Manages The Webserver Routes and APNS for Apples MDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"fmt" //// TODO: Eliminate This (Only Used Once)
  //"io/ioutil"
  "net/http"
  "os"

  "github.com/op/go-logging" // Logging
	"github.com/gorilla/mux" // HTTP Router
  "github.com/go-pg/pg" // Database (Postgres)

	"github.com/groob/plist"
)

var (
  log *logging.Logger // TODO: This is A TEMP Sub Logging Solution
  pgdb *pg.DB
)

//TODO:
//	Auto Generate Profiles From Config Details Parsed In (Cache Sirectory)

// TODO:
//  Func Based (struct, err) Error Handling
//	Logging Log Line Number It Was Caled From
//	Redo Logging Messages/Levels

// TODO:
//  Setup Logger For These Sub Packages
//  Capitialise UDID in Database And Find Out What is Causing That Not To Work
//	Alert Admin And Prune Device That Are Do Not Deployed After Set Amount Of Time
//  Better, More Informative Error Messages

//  Add Logging To File Or Something For Any Errors Occurred (Debugging For The Me)
//  See What Checkin Does If None Of The Core Values (4 Of Them) Are Not Given By The Client Does It Plist Parsing Error?
//  Detect Device That Have Disconnected From Management
//  Prevent APNS Module Form Causing "DDOS" To Apples Servers
//  Use Verify Stuff To Stop People Forging The Enrollment Profile Even If They Know The URL's
//  Switch The Order Of All Routers So HandleFunc is After Attributes

func Init(_pgdb *pg.DB, _log *logging.Logger) {
  pgdb = _pgdb
  log = _log
	fmt.Println("Started The Apple MDM Module!")
}

func Mount(r *mux.Router) {
	r.HandleFunc("/", genericResponse).Methods("GET")

  r.HandleFunc("/ping-apns", pingApnsHandler).Methods("GET")
	r.HandleFunc("/testing", testingHandler).Methods("GET")
	r.HandleFunc("/WWDC_App.plist", wwdcHandler).Methods("GET")

	r.HandleFunc("/enroll", enrollHandler).Methods("GET")
	r.HandleFunc("/checkin", checkinHandler).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.HandleFunc("/server", serverHandler).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
	//r.HandleFunc("/scep", scepHandler).Methods("GET")
}

func genericResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Apple Mobile Device Management Server!")
}

func enrollHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-apple-aspen-config")
	http.ServeFile(w, r, "enroll.mobileconfig")
}








// The "/checkin" route is used to check if the device can join the mdm and update its push token to the server
func checkinHandler(w http.ResponseWriter, r *http.Request) {
  var cmd CheckinCommand
  if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
    log.Debug("Error Parsing Checkin Request: ", err)
    w.WriteHeader(http.StatusBadRequest)
    return
  }
	device := getDevice(cmd.UDID)

  if cmd.MessageType == "Authenticate" {
    if cmd.auth.OSVersion != "" && cmd.auth.BuildVersion != "" && cmd.auth.ProductName != "" && cmd.auth.SerialNumber != "" && cmd.auth.IMEI != "" && cmd.auth.MEID != "" {
			if device == nil {
				device = newDevice(cmd)

				if status := editDevice(device, false); status == false {
	        log.Debug("Failure To Add New Device To The Database")
	        w.WriteHeader(http.StatusUnauthorized) //TODO: Check This Kills The Client Joining
	      } else {
	        w.WriteHeader(http.StatusOK)
	      }
			} else {
				log.Warning("An Existing Device Has Requested To Enroll. -> 403 (Unauthorized)")
				w.WriteHeader(http.StatusUnauthorized)
			}
    }
  } else if cmd.MessageType == "TokenUpdate" && device != nil {
    if cmd.update.Token != nil && cmd.update.PushMagic != "string" && cmd.update.UnlockToken != nil && (cmd.update.AwaitingConfiguration == true || cmd.update.AwaitingConfiguration == false) {
			if device.Deployed == false {
				device.Deployed = true

				if status := editDevice(device, true); status == false {
	        log.Debug("Failure To Update The Devices Deployment Status")
	        w.WriteHeader(http.StatusUnauthorized)
					return
	      } else {
					log.Info("A New Device Joined The MDM: " + device.UDID)
	      }
			}

			if cmd.update.AwaitingConfiguration == true {
        // TODO: Do DEP Prestage Enrollment By Pushing The Profiles Now Then Run The Fully Setup Thing Then Push The Finished Command
        log.Error("Unsupported DEP Features")
				w.WriteHeader(http.StatusUnauthorized)
				return
      }

      device.Token = cmd.update.Token
      device.PushMagic = cmd.update.PushMagic
      if cmd.update.UnlockToken != nil {
        device.UnlockToken = cmd.update.UnlockToken
      }

			if status := editDevice(device, true); status == false {
				log.Debug("Failure To Update The Devices APNS Tokens")
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				log.Debug("Device Updated Its APNS Keys: " + device.UDID)
				w.WriteHeader(http.StatusOK)
			}
    } else {
      log.Warning("A Device Requested To Join With An Invalid Setup (Pre IOS 9 or Doesn't Have Perms)")
      w.WriteHeader(http.StatusUnauthorized)
    }
	} else if cmd.MessageType == "CheckOut" && device != nil {
		err := pgdb.Delete(device)
	  if err != nil {
	      log.Fatal("Failure To CheckOut Device. The Device Will Not Try Again", err)
				w.WriteHeader(http.StatusBadRequest)
				return
	  }
		log.Debug("A Device Has Been Removed From The MDM: " + cmd.UDID)
		w.WriteHeader(http.StatusOK)
  } else {
		if device != nil {
			log.Warning("Unkown Checkin MessageType of: " + cmd.MessageType)
		} else {
			log.Warning("A Device Not In The Database Attempted An Action: " + cmd.MessageType)
		}
		w.WriteHeader(http.StatusBadRequest)
  }
}





func pingApnsHandler(w http.ResponseWriter, r *http.Request) {
	devices := getDevices()

	if devices == nil { //[]Device{} {
		log.Debug("Error Getting Devices Or There Are None")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error Getting Devices Or There Are None")
		return
	}

  for _, device := range devices {
    log.Debug("APNS Update Sent To Device " + device.UDID)
		status := deviceUpdate(device)

		if !status {
	    w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error Sending APNS Update To The Device: " + device.UDID)
	    return
	  }
  }

  fmt.Fprintf(w, "All Devices Have Been Told To Update")
}



func testingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing Route!")
}















var run_commands = 1

func serverHandler(w http.ResponseWriter, r *http.Request) {
	/*buf, err := ioutil.ReadAll(r.Body)
	//r.Body.Close()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(string(buf))*/


	/* Break */

	var cmd DeviceStatus
  if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
    log.Debug("Error Parsing Checkin Request: ", err)
    w.WriteHeader(http.StatusBadRequest)
    return
	}
	device := getDevice(cmd.UDID)

	if device != nil && device.Deployed {
		if cmd.Status == "Idle" {
			log.Debug("Idle Device: " + device.UDID)
		} else if cmd.Status == "Acknowledged" { //TODO: Storing Feedback
			log.Warning("Device Successed At: " + cmd.CommandUUID)
		} else if cmd.Status == "Error" { //TODO: Storing Feedback
			log.Warning("Device Failed At: " + cmd.CommandUUID)
			log.Warning(cmd.ErrorChain)



		} else {
			log.Debug("Unkown Device Status of: " + cmd.Status)
		}






		if run_commands == 1 {
			run_commands = run_commands+1
			log.Info("Sending Lock Payload")

			LockPayload := ServerCommand{
				CommandUUID: "BBA5879E-2649-43B1-9934-D0D26BBC0E5D", //TODO: Build Generator For These
				Command: ServerCommandBody{
					RequestType: "DeviceLock",
				},
			}

			out, err := plist.MarshalIndent(LockPayload, "   ") //TODO: Clean This Plist Parsing And Error Handling (And Other Ones Using The Same Code)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(w, string(out))
		} else if run_commands == 2 {
			run_commands = run_commands+1
			log.Info("Sending WWDC App Payload")

			AppPayload := ServerCommand{
				CommandUUID: "4424F929-BDD2-4D44-B518-393C0DABD56A", //TODO: Build Generator For These
				Command: ServerCommandBody{
					RequestType: "InstallApplication",
					PayloadInstallApplication: PayloadInstallApplication{
						ITunesStoreID: 640199958, //WWDC App
						//ManifestURL: "https://mdm.otbeaumont.me/apple/WWDC_App.plist",
						ManagementFlags: 4, //Understand This
					},
				},
			}

			out, err := plist.MarshalIndent(AppPayload, "     ") //TODO: Clean This Plist Parsing And Error Handling (And Other Ones Using The Same Code)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintf(w, string(out))
			/*} else if run_commands == 3 {
				run_commands = run_commands+1
				log.Info("Sending Block Facetime Profile Payload")

				profileContent, _ := ioutil.ReadFile("BlockFacetime.mobileconfig") //TODO: Handle Errors
				AppPayload := ServerCommand{
					CommandUUID: "5428B959-BDD2-4H45-Q558-397I0DABD56B", //TODO: Build Generator For These
					Command: ServerCommandBody{
						RequestType: "InstallProfile",
						PayloadInstallProfile: PayloadInstallProfile{
							Payload: []byte(profileContent),
						},
					},
				}

				out, err := plist.MarshalIndent(AppPayload, "     ") //TODO: Clean This Plist Parsing And Error Handling (And Other Ones Using The Same Code)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Fprintf(w, string(out))*/
		} else {
			log.Debug("Device Deployed")
			fmt.Fprintf(w, "")
		}





	} else {
		log.Warning("A Device Attempted To Get Actions From Server Without Fully Enrolling")
	}



	/*buf, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		//Do something
		fmt.Println(err)
	}

	fmt.Println(string(buf))
	w.WriteHeader(http.StatusOK)
	return*/

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

func wwdcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	http.ServeFile(w, r, "WWDC_App.plist")
}




func exitOnError(err error) { //TODO: Maybe Remove This Or Merge To New Version
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
