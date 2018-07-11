/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Apple MDM Core. It Manages The Webserver Routes and APNS for Apples MDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"fmt" //// TODO: Eliminate This (Only Used Once)
  "io/ioutil"
  "net/http"
  "os"

  "github.com/op/go-logging" // Logging
	"github.com/gorilla/mux" // HTTP Router
  "github.com/go-pg/pg" // Database (Postgres)


	"github.com/groob/plist"
	"encoding/hex"
	"encoding/json"
	"github.com/RobotsAndPencils/buford/certificate"
	"github.com/RobotsAndPencils/buford/payload"
	"github.com/RobotsAndPencils/buford/push"
)

var (
  log *logging.Logger // TODO: This is A TEMP Sub Logging Solution
  pgdb *pg.DB
)

// TODO:
//  Setup Logger For These Sub Packages
//  Capitialise UDID in Database And Find Out What is Causing That Not To Work


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

	r.HandleFunc("/enroll", enrollHandler).Methods("GET")
	r.HandleFunc("/checkin", checkinHandler).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	//r.HandleFunc("/server", serverHandler).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
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
    log.Error("Failed To Parse Checkin Request")
    log.Error(err)
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if cmd.MessageType == "Authenticate" {
    if cmd.auth.OSVersion != "" && cmd.auth.BuildVersion != "" && cmd.auth.ProductName != "" && cmd.auth.SerialNumber != "" && cmd.auth.IMEI != "" && cmd.auth.MEID != "" {
      log.Info("New Device Trying To Join The MDM")

      device := Device{
        SerialNumber: cmd.auth.SerialNumber,
        ProductName: cmd.auth.ProductName,
        OSVersion: cmd.auth.OSVersion,
        Topic: cmd.Topic,
        UDID: cmd.UDID,
        Token: []byte{},
        PushMagic: "",
        UnlockToken: []byte{},
      }
      err := pgdb.Insert(&device) //.OnConflict("DO NOTHING") //TODO: Check The Do Nothing Works Then make It Error
      if err != nil {
        log.Fatal(err)
        w.WriteHeader(http.StatusUnauthorized) //TODO: Check This Kills The Client Joining
      } else {
        w.WriteHeader(http.StatusOK)
      }
    }
  } else if cmd.MessageType == "TokenUpdate" {
    if cmd.update.Token != nil && cmd.update.PushMagic != "string" && cmd.update.UnlockToken != nil && (cmd.update.AwaitingConfiguration == true || cmd.update.AwaitingConfiguration == false) {
      if cmd.update.AwaitingConfiguration == true {
        //Do DEP Prestage Enrollment By Pushing The Profiles Now Then Run The Fully Setup Thing
        //TODO: Debug Event (Until Feature is Built)
        fmt.Println("Unsupported DEP Features")
        //TODO: Future Feature: If set to true, the device is awaiting a DeviceConfigured MDM command before proceeding through Setup Assistant.
      }

      log.Info(cmd.UDID)

      var device Device
      err := pgdb.Model(&device).Where("udid = ?", cmd.UDID).Select()
      if err != nil {
        log.Fatal(err)
        return
      }
      // TODO If Not Found Handle That Separatly To Other Errors

      fmt.Println(device)

      device.Token = cmd.update.Token
      device.PushMagic = cmd.update.PushMagic
      if cmd.update.UnlockToken != nil {
        device.UnlockToken = cmd.update.UnlockToken
      }

      err2 := pgdb.Update(&device)
      if err2 != nil {
        log.Fatal(err2)
        w.WriteHeader(http.StatusUnauthorized)
      } else {
        w.WriteHeader(http.StatusOK)
        fmt.Println("Device Updated Its Tokens")
      }
    } else {
      log.Warning("A Device Requested To Join With An Invalid Setup (Pre IOS 9 or Doesn't Have Perms)")
      w.WriteHeader(http.StatusUnauthorized)
    }
  } else {
    log.Warning("Unkown Checkin MessageType of: " + cmd.MessageType)
    w.WriteHeader(http.StatusBadRequest)
  }
}





func pingApnsHandler(w http.ResponseWriter, r *http.Request) {
  var devices []Device
  err := pgdb.Model(&devices).Select()
  if err != nil {
		log.Error(err)
		return
	}

  for _, device := range devices {
    log.Debug("APNS Update Sent To Device " + device.UDID)

    cert, err := certificate.Load("PushCert.p12", "password")
    if err != nil {
      log.Error(err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    client, err := push.NewClient(cert)
    if err != nil {
      log.Error(err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    service := push.NewService(client, push.Production)

    // construct a payload to send to the device:
    p := payload.MDM{
      Token: device.PushMagic,
    }
    b, err := json.Marshal(p)
    if err != nil {
      log.Error(err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    // push the notification:
    deviceToken := hex.EncodeToString(device.Token)

    if !push.IsDeviceTokenValid(deviceToken) {
      fmt.Println("The Device Token Is Incorrect")
      return
    }

    id, err := service.Push(deviceToken, nil, b)
    if err != nil {
      log.Error(err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    log.Debug("APNS ID: " + id)
  }

  fmt.Fprintf(w, "Sent APNS Update For All Devices")
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

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
