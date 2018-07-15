/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is The Apple MDM Core. It Manages The Webserver Routes and APNS for Apples MDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"fmt"
	"net/http"
	"time"

	//External Deps
	"github.com/gorilla/mux" // HTTP Router
	"github.com/groob/plist" //Plist Parsing

	// Internal Functions
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database"      //Mattrax Database
	errors "github.com/mattrax/Mattrax/internal/errors"     // Mattrax Error Handling
	mlg "github.com/mattrax/Mattrax/internal/logging"       //Mattrax Logging

	// Internal Modules
	restAPI "github.com/mattrax/Mattrax/appleMDM/api"     // The Apple MDM REST API
	structs "github.com/mattrax/Mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)

var pgdb = mdb.GetDatabase()
var log = mlg.GetLogger()
var config = mcf.GetConfig() // Get The Internal State

func Init() { log.Info("Loaded The Apple MDM Module") }

func Mount(r *mux.Router) {
	r.HandleFunc("/", genericResponse).Methods("GET")
	r.HandleFunc("/enroll", enrollHandler).Methods("GET")

	//REST API
	restAPI.Mount(r.PathPrefix("/api/").Subrouter())

	// MDM Device Endpoints
	r.Handle("/inform", errors.Handler(informHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.Handle("/server", errors.Handler(serverHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
}

func genericResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Apple Mobile Device Management Server!")
}
func enrollHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-apple-aspen-config")
	http.ServeFile(w, r, "data/enroll.mobileconfig")
}

// The "/inform" route is used to check if the device can join the mdm and update its push token to the server (In The Apple Docs This Is Refered To As The Check-In Route)
func informHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	//Parse The Request
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
				if _, err := pgdb.Model(&enrollingDevice).Where("uuid = ?", &cmd.UDID).Update(); err != nil {
					return 403, err
				}
			} else {
				if _, err := pgdb.Model(&enrollingDevice).Set("uuid = ?", &cmd.UDID).Insert(); err != nil {
					return 403, err
				}
			}

			fmt.Fprintf(w, "")
			return 200, nil
		}
	} else if cmd.MessageType == "TokenUpdate" {
		if cmd.Update.Token == nil && cmd.Update.PushMagic == "" && cmd.Update.UnlockToken == nil && (cmd.Update.AwaitingConfiguration == true || cmd.Update.AwaitingConfiguration == false) {
			return 403, errors.New("The Request To 'TokenUpdate' From The Device Is Malformed Or Their Device Is Pre IOS 9 or Is Missing The Device Information Permission In The Profile")
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
			if err := pgdb.Update(&device); err != nil {
				return 403, err
			}
			log.Info("A New Device Joined The MDM: " + device.UDID)
		} else if device.DeviceState == 4 {
			return 403, errors.New("A Not Enrolled Device Tried To Do A 'TokenUpdate'")
		} else if cmd.Update.AwaitingConfiguration { //Device Enrollment Program
			// TODO: Do DEP Prestage Enrollment By Pushing The Profiles Now Then Run The Fully Setup Thing Then Push The Finished Command
			return 403, errors.New("DEP Is Currently Not Supported But Is Coming Soon")
		}

		device.DeviceTokens.Token = cmd.Update.Token
		device.DeviceTokens.PushMagic = cmd.Update.PushMagic
		if cmd.Update.UnlockToken != nil {
			device.DeviceTokens.UnlockToken = cmd.Update.UnlockToken
		}

		if err := pgdb.Update(&device); err != nil {
			return 403, err
		}
		log.Debug("Device Updated Its APNS Keys: " + device.UDID)
		return 200, nil
	} else if cmd.MessageType == "CheckOut" {
		device.DeviceState = 4
		if err := pgdb.Update(&device); err != nil {
			return 403, err
		}
		log.Info("A Device Has Been Removed From The MDM: " + cmd.UDID)
		return 200, nil
	} else {
		return 403, errors.New("A Device Not In The Database Attempted An Action: " + cmd.MessageType)
	}
} //TODO: Double Check If It Tried To TokenUpdate Without DB Entry It Still Fails







//TODO: Redo Exit Resaving DB Handler (Maybe Add it To The errors Router Controller)
//TODO: Func Description -> Make Sure It Doesn't DOS Itself By Returning 403 Lots
func serverHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	//Parse The Request
	var cmd structs.Response
	if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { return 403, err }
	//Attempt To Get The Device From the Database
	var device structs.Device
	if err := pgdb.Model(&device).Where("uuid = ?", cmd.UDID).Select(); err != nil && errors.PgError(err) { return 403, err }
	//Handle The Request
	if device.DeviceState != 3 { return 403, errors.New("A Device That is Not Known To The MDM Tried To /server") } //TODO: Redo Message

	if cmd.Status == "Acknowledged" {
		if device.DevicePolicies.Inventory.CommandUUIDs[cmd.CommandUUID] != "" { //TODO: Cehck false isn't Returned By Blank
			requestType := device.DevicePolicies.Inventory.CommandUUIDs[cmd.CommandUUID] //TODO: Rename This var

			log.Warning("Inventory Req: ", cmd.CommandUUID)


			//Switch Here -> requestType

			if requestType == "ProfileList" {
				device.DeviceDetails.Profiles = cmd.ProfileList


			} else {
				log.Warning("Don't Know How to Handle That")
			}



			delete(device.DevicePolicies.Inventory.CommandUUIDs, cmd.CommandUUID)
		} else {
			log.Warning("Authentication For Non Inventory Thingo")

			//TODO




		}
	}












	if (time.Now().Unix()-device.DevicePolicies.LastUpdate > 30) { //TODO: Redo Timer Mechanic And Set A Sane Default. Allow Changing Through config.json (Bigger Deployments Will Need it)
		return doInventory(w, r, cmd, device)
	} else {
		if err := pgdb.Update(&device); err != nil { return 403, err }
	  return 200, nil
	}

	//Check For Update Policies To Update

	return 200, nil
}

func returnPlist(w http.ResponseWriter, payload *structs.Payload) error {
	plistCmd, err := plist.MarshalIndent(payload, "\t")
	if err != nil { return err }


	//log.Info(string(plistCmd)) //TEMP
	fmt.Fprintf(w, string(plistCmd)) //TODO: Can This be Done Without string() (Via Streams To Make It More Efficent)
	return nil
}





















func temp(w http.ResponseWriter, r *http.Request) (int, error) {
	request := new(structs.Command)
	request.RequestType = "DeviceLock"
	request.DeviceLock.Message = "Locked By Administrator" //Not Showing On Device

	payload, err := structs.NewPayload(request)
	if err != nil {
		return 403, err
	}

	// Encode in a plist and print to stdout
		// uses the github.com/groob/plist package
	/*encoder := plist.NewEncoder(os.Stdout)
	encoder.Indent("  ")
	if err := encoder.Encode(payload); err != nil {
		log.Fatal(err)
	}*/ //Return The Stream For Web Request






	plistCmd, err := plist.MarshalIndent(payload, "\t")
	if err != nil { return 403, err }


	log.Info(string(plistCmd)) //TEMP
	fmt.Fprintf(w, string(plistCmd))


	return 200, nil
}










var locked = false

func Disabled(w http.ResponseWriter, r *http.Request) (int, error) {
	response := &structs.Response{}
	if err := plist.NewXMLDecoder(r.Body).Decode(response); err != nil { return 403, err }

	if !locked {
		locked = true

		request := new(structs.Command)
		request.RequestType = "DeviceLock"
		request.DeviceLock.Message = "Locked By Administrator" //Not Showing On Device

		payload, err := structs.NewPayload(request)
		if err != nil {
			return 403, err
		}

		// Encode in a plist and print to stdout
	    // uses the github.com/groob/plist package
		/*encoder := plist.NewEncoder(os.Stdout)
		encoder.Indent("  ")
		if err := encoder.Encode(payload); err != nil {
			log.Fatal(err)
		}*/ //Return The Stream For Web Request






		plistCmd, err := plist.MarshalIndent(payload, "\t")
		if err != nil { return 403, err }


		log.Info(string(plistCmd)) //TEMP
		fmt.Fprintf(w, string(plistCmd))


		return 200, nil
	}


	log.Info(response.Status)
	if response.Status == "Acknowledged" {
		locked = false
	}


	return 200, nil
}




//TODO: Come Up With a Better name For The /server route and the /checkin route and set it up
//TODO: Reformat Apple MDM Code Files (MDM Routes Both In Independant Files)
//TODO: Redo All Log Messages, Error Codes Before v1 Release
//TODO: Handle Shutting Down In The Middle of An Inventory/Update And The The New Update Being Different
//TODO: System For Updater To Backup The DB, Install Update Restart And Do Checks And Monitoring For An Hour (Auto Rollback On Certain Errors) Before Returning To Normal
//TODO: Way To Using Postgres Get The Devices With A Policy
//Handle Devices DOSing The Server When it Keeps Failing -> Prevent It Fast And Alert The Admin

//TODO Maybe Do "/connect" Route Like The MDM Spec Wants For Trust Certs
//Read Full MDM Spec And Make Everything Bar School Manager Intergration Work
// VPP Support Through The API (Act As Proxy and Rate Limit, etc)
// Label All Functions In The appleMDM/structs/database.go File
//TODO: Is There A Need For A APNS Package?
//Software Update Caching Server Built In (Separate Module)
//TODO: Chnage /server to /checkin and /checkin to /register or something else
// Features:
//	 Auto Generate Profiles From Config Details Parsed In (Cache In Subdirectory)
//	 After Enrolling Do Inventory -> Get All The Devices Profiles, Apps, Details, Configuration, etc
//	 Detect Via APNS (And Logging) If Device Was Removed Without CheckOut Working -> Alert Admin
//	 Auto Prune Devices That Did Not Complete Enrollment -> Alert Admin
//	 Postgres Handle Database Lossing Connection and Stress Testing -> Handle Errors
//	 Clean APNS. Combine Multiple APNS Into One Request (For Bulk Without DOSing Apple)
//	 Prevent Forging Enrollment Certificates
//	 Prestage Enrollment (Template For Devices) -> DEP Support

//  See What Checkin Does If None Of The Core Values (4 Of Them) Are Not Given By The Client Does It Plist Parsing Error?
