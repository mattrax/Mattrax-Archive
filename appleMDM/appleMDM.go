/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is The Apple MDM Core. It Manages The Webserver Routes and APNS for Apples MDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"fmt" //// TODO: Eliminate This (Only Used Once)
  "net/http"

	"github.com/sirupsen/logrus" // Logging
	//"github.com/go-pg/pg" // Database (Postgres)
	"github.com/gorilla/mux" // HTTP Router
	"github.com/groob/plist" //Plist Parsing

	errors "github.com/mattrax/mattrax/internal/errors" // Internal Error Handling (For HTTP)
	restAPI "github.com/mattrax/mattrax/appleMDM/api"
	structs "github.com/mattrax/mattrax/appleMDM/structs"

	mdb "github.com/mattrax/mattrax/internal/database" //Mattrax Database
)

var pgdb = mdb.GetDatabase()
var log *logrus.Logger //; pgdb *pg.DB )

func Init(_log *logrus.Logger) {  log = _log //pgdb = _pgdb; //_pgdb *pg.DB,
	log.Info("Started The Apple MDM Module!")
	restAPI.Init(pgdb, _log)
}

func Mount(r *mux.Router) {
	r.HandleFunc("/", genericResponse).Methods("GET")
	r.HandleFunc("/enroll", enrollHandler).Methods("GET")
	//REST API
	restAPI.Mount(r.PathPrefix("/api/").Subrouter())
	// MDM Device Endpoints
	r.Handle("/checkin", errors.Handler(checkinHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.Handle("/server", errors.Handler(serverHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
}

func genericResponse(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Apple Mobile Device Management Server!") }
func enrollHandler(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "application/x-apple-aspen-config"); http.ServeFile(w, r, "data/enroll.mobileconfig") }











// The "/checkin" route is used to check if the device can join the mdm and update its push token to the server
func checkinHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	//Parse Request
	var cmd structs.CheckinCommand
  if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { return 403, err }
	//Attempt To Get The Device From the Database
	var device structs.Device
  if err := pgdb.Model(&device).Where("uuid = ?", cmd.UDID).Select(); err != nil && errors.PgError(err) { return 403, err }
	//Handle The Request
  if cmd.MessageType == "Authenticate" { //New Device Requesting To Enrolling
    if cmd.Auth.OSVersion == "" && cmd.Auth.BuildVersion == "" && cmd.Auth.ProductName == "" && cmd.Auth.SerialNumber == "" && cmd.Auth.IMEI == "" && cmd.Auth.MEID == "" {
			return 403, errors.New("Internal: The Request To 'Authenticate' From The Device Is Malformed")
		} else { //Create A New Device In The Database And Return Success To The Device
			enrollingDevice := structs.NewDevice(cmd)
			enrollingDevice.DeviceState = 1
			if device.DeviceState == 3 {
				return 403, errors.New("A Device That is Already Enrolled Attempted To Enroll")
			} else if device.DeviceState <= 4 && device.UDID != "" { // TODO: If The Device State Is 0, 1 or 2 Retain Certain Information Here
				if _, err := pgdb.Model(&enrollingDevice).Where("uuid = ?", &cmd.UDID).Update(); err != nil { return 403, err }
			} else {
				if _, err := pgdb.Model(&enrollingDevice).Set("uuid = ?", &cmd.UDID).Insert(); err != nil { return 403, err }
			}

			fmt.Fprintf(w, "")
			return 200, nil
    }
	} else if cmd.MessageType == "TokenUpdate" {
		if cmd.Update.Token == nil && cmd.Update.PushMagic == "" && cmd.Update.UnlockToken == nil && (cmd.Update.AwaitingConfiguration == true || cmd.Update.AwaitingConfiguration == false) {
			return 403, errors.New("The Request To 'TokenUpdate' From The Device Is Malformed Or Thier Device Is Pre IOS 9 or Is Missing The Device Information Permission In The Profile")
		} else if device.DeviceState == 0 {
			return 403, errors.New("A Device Tried To Do A TokenUpdate Without Having Enrolled Via A 'Authenticate' Request")
		} else if device.DeviceState == 1 {
			// TODO: Handle DEP (Currently Bypassed)
			/* TEMP Bypass */
			device.DeviceState = 3
			if err := pgdb.Update(&device); err != nil { return 403, err }
			log.Info("A New Device Joined The MDM: " + device.UDID)
			/* End Bypass */

		} else if device.DeviceState == 2 {
			device.DeviceState = 3
			if err := pgdb.Update(&device); err != nil { return 403, err }
			log.Info("A New Device Joined The MDM: " + device.UDID)
		} else if device.DeviceState == 4 {
			return 403, errors.New("A Not Enrolled Device Tried To Do A 'TokenUpdate'")
		} else if cmd.Update.AwaitingConfiguration { //Device Enrollment Program
			// TODO: Do DEP Prestage Enrollment By Pushing The Profiles Now Then Run The Fully Setup Thing Then Push The Finished Command
			return 403, errors.New("DEP Is Currently Not Supported But Is Coming Soon")
		}

		device.DeviceTokens.Token = cmd.Update.Token
		device.DeviceTokens.PushMagic = cmd.Update.PushMagic
		if cmd.Update.UnlockToken != nil { device.DeviceTokens.UnlockToken = cmd.Update.UnlockToken }

		if err := pgdb.Update(&device); err != nil { return 403, err }
		log.Debug("Device Updated Its APNS Keys: " + device.UDID)
		return 200, nil
	} else if cmd.MessageType == "CheckOut" {
		device.DeviceState = 4
		if err := pgdb.Update(&device); err != nil { return 403, err }
		log.Info("A Device Has Been Removed From The MDM: " + cmd.UDID)
		return 200, nil
  } else {
		return 403, errors.New("A Device Not In The Database Attempted An Action: " + cmd.MessageType)
	}
} //TODO: Double Check If It Tried To TokenUpdate Without DB Entry It Still Fails



// TOMORROW:
//	New Logging
//	Clean Chrome
//	Clean The errorHandling Package + Chnage Import Name (utils)
//	Clean/Redo The main.go File
//	The /server Route




var run_commands = 1

func serverHandler(w http.ResponseWriter, r *http.Request) (int, error) {






	return 200, nil //TEMP (Disable Everything Below) //TODO: Next Few Comit This Will Be Redone For Pushing Policies

	//The Code That Goes Here Is In A File Called backup_server_handler Before The Redo (Error Handling Broke It)
}













//Handle Devices DOSing The Server When it Keeps Failing -> Prevent It Fast And Alert The Admin

// Try To:
//	 Simplier Logging Library. Less Configuration and More Opionated. -> Subfile Logging (Separated) Support
//	 Better/Neater Error Handling
// 				-  Handle PG Errors From (Maybe Add More In PR If Needed To Add More): https://github.com/go-pg/pg/blob/master/error.go
//   Func Based (struct, err) Error Handling Insead Of *struct (Remove All Of *struct)
//	 Neater Error Messages/Logging Output (Decide What Each Log Level Is For)

// Features:
//	 Auto Generate Profiles From Config Details Parsed In (Cache In Subdirectory)
//	 After Enrolling Do Inventory -> Get All The Devices Profiles, Apps, Details, Configuration, etc
//	 Detect Via APNS (And Logging) If Device Was Removed Without CheckOut Working -> Alert Admin
//	 Auto Prune Devices That Did Not Complete Enrollment -> Alert Admin
//	 Postgres Handle Database Lossing Connection and Stress Testing -> Handle Errors
//	 Clean APNS. Combine Multiple APNS Into One Request (For Bulk Without DOSing Apple)
//	 Prevent Forging Enrollment Certificates
//	 Prestage Enrollment (Template For Devices) -> DEP Support

// Future Features:
//   Build Tests -> For All Function And Routes (Fake Device Requests/Response Verifying)
//	 Optimisng Preformance


//  See What Checkin Does If None Of The Core Values (4 Of Them) Are Not Given By The Client Does It Plist Parsing Error?
