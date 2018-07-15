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
	//"time"

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

	micromdm "github.com/mattrax/Mattrax/appleMDM/structs/micromdm" // MicroMDM Structs TEMP: Redo This MSG
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
	r.Handle("/checkin", errors.Handler(checkinHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm-checkin")
	r.Handle("/server", errors.Handler(serverHandler)).Methods("PUT").HeadersRegexp("Content-Type", "application/x-apple-aspen-mdm")
}

func genericResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Apple Mobile Device Management Server!")
}
func enrollHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-apple-aspen-config")
	http.ServeFile(w, r, "data/enroll.mobileconfig")
}

// The "/checkin" route is used to check if the device can join the mdm and update its push token to the server
func checkinHandler(w http.ResponseWriter, r *http.Request) (int, error) {
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
			if err := pgdb.Update(&device); err != nil {
				return 403, err
			}
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

var run_commands = 1
var done = true
var locked = false

func serverHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	response := &micromdm.Response{}
	if err := plist.NewXMLDecoder(r.Body).Decode(response); err != nil { return 403, err }

	log.Info(response.Status)

	if !locked {
		locked = true


		// create a request
		//request := micromdm.DeviceLock{ Message: "Locked" } //PIN: "",  , PhoneNumber: "123-4567"
		/*request := &CommandRequest{
			RequestType: "DeviceInformation",
			Queries:     []string{"IsCloudBackupEnabled", "BatteryLevel"},
		}*/



		
		request := &micromdm.CommandRequest{ // Working Lock Device
			UDID: "HelloWorld",
			Command: micromdm.Command{
				RequestType: "DeviceLock",

			},
		}



		payload, err := micromdm.NewPayload(request)
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


		log.Info(string(plistCmd))
		fmt.Fprintf(w, string(plistCmd))


		return 200, nil
	}







	return 200, nil
}






func old(w http.ResponseWriter, r *http.Request) (int, error) {
	/*payload := structs.ServerCommand{
		CommandUUID: "BBA5879E-2649-43B1-9934-D0D26BBC0E5D", //TODO: Build Generator For These
		Command: structs.ServerPayload{
			RequestType: "DeviceLock",
		},
	}

	out, err := plist.MarshalIndent(payload, "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions
	fmt.Fprintf(w, string(out)) // TODO: This Stuff Is Tiwse (The Plist PArsing) Make That Not A Thing

	fmt.Println(string(out))
	return 200, nil*/





	//Parse The Request
	var cmd structs.DeviceStatus
	if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { return 403, err }
	//Attempt To Get The Device From the Database
	var device structs.Device
	if err := pgdb.Model(&device).Where("uuid = ?", cmd.UDID).Select(); err != nil && errors.PgError(err) { return 403, err }
	//Handle The Request
	if !(device.DeviceState <= 3) { return 403, errors.New("A Device Tried To /server Without Existing In The DB") } //TODO: make This make Sense
	if device.DeviceState != 3 { return 403, errors.New("A Device In A Not Fully Enrolled State Accessed The /server Route") } //TODO: make This make Sense

	if cmd.Status == "Idle" {
		log.Info("The Device Is Idle: ", device.UDID)
	} else {
		log.Warning("The Device Is Not Idle", cmd.Status) //TEMP
	}

	if !locked {
		locked = true

		out, _ := structs.ParsePayload(structs.ServerPayload{ //TODO Error Handling
			RequestType: "DeviceLock",
		})

		log.Info(out)

		fmt.Fprintf(w, out)
		return 200, nil


		/*payload := structs.ServerCommand{
			CommandUUID: "BBA5879E-2649-43B1-9934-D0D26BBC0E5D", //TODO: Build Generator For These
			Command: structs.ServerPayload{
				RequestType: "DeviceLock",
			},
		}

		out, err := plist.MarshalIndent(payload, "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions
		fmt.Fprintf(w, string(out)) // TODO: This Stuff Is Tiwse (The Plist PArsing) Make That Not A Thing

		fmt.Println(string(out))
		return 200, nil*/
	}

	return 200, nil











	//log.Info(cmd.InstalledApplicationList.Applications)

	if cmd.Status == "Idle" {
		log.Debug("Idle Device ", cmd.UDID)
	} else {
		log.Debug("Device In In State ", cmd.Status, ": ", cmd.UDID) // TODO: Redo Text

		fmt.Fprintf(w, "") // TEMP
		return 200, nil // TEMP
	}

	if done { //If Inventory Needs To Happen
		done = false
		log.Debug("Doing Inventory On Device: ", device.UDID)

		//micromdm.NewPayload

		/*payload := micromdm.Payload{u.String(),
			&Command{RequestType: requestType}}

		out, err := plist.MarshalIndent(payload, "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions
		fmt.Fprintf(w, string(out))
		log.Info(string(out))*/
		return 200, nil


		/* Kinda Temp Inventory */
		/*currentAction := structs.ServerCommand{
			CommandUUID: "741B39F2-649B-4BB6-A522-EE2YF2D99D26",
			Command: structs.ServerPayload{
				RequestType: "DeviceInformation",
				PayloadDeviceInformation: structs.PayloadDeviceInformation{
					Queries: []string{ "BatteryLevel" },
				},
			},
		}*/

		/*currentAction := structs.ServerCommand{
			CommandUUID: "741B39F2-649B-4BB6-A522-EE2YF2D99D26",
			Command: structs.ServerPayload{
				RequestType: "InstalledApplicationList",
			},
		}




		out, err := plist.MarshalIndent(currentAction, "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions
		fmt.Fprintf(w, string(out)) // TODO: This Stuff Is Tiwse (The Plist PArsing) Make That Not A Thing

		fmt.Println(string(out))
		return 200, nil*/
		/* End Kinda Temp Inventory */
	} else {
		log.Debug("No Inventory is needed")

		log.Warning(cmd)
	}


	//Send Stuff




	/*
	if device.DeviceState != 3 {
		//Error
	} else {
		if run_commands == 2 {
			log.Fatal(cmd)
		}
		run_commands = run_commands +1

		if cmd.Status != "Idle" {
			log.Debug("Sent Payload")

			currentAction := structs.ServerCommand{
				CommandUUID: "741B39F2-649B-4BB6-A522-EE2YF2D99D26",
				Command: structs.ServerPayload{
					RequestType: "DeviceInformation",
					PayloadDeviceInformation: structs.PayloadDeviceInformation{
						Queries: []string{ "BatteryLevel" },
					},
				},
			}

			out, err := plist.MarshalIndent(currentAction, "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions
			fmt.Fprintf(w, string(out)) // TODO: This Stuff Is Tiwse (The Plist PArsing) Make That Not A Thing


		}



		if device.DevicePolicies.CurrentAction.UDID != "" {
			log.Info("Device Is Currently Doing: ", device.DevicePolicies.CurrentAction.Name)

			out, err := plist.MarshalIndent(device.DevicePolicies.CurrentAction.Actions[0], "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions

			//log.Info(string(out))
			fmt.Fprintf(w, string(out))
			//TODO: Do This For all Actions Till Done (On Each Request). Remove Then When Device Sends Success Respone or Fails Twice In A Row
			//if //Inventory
			//Currently Running Something
		} else {
			log.Debug("Time Since Last Inventory ", time.Now().Unix()-device.DevicePolicies.LastUpdate) //TEMP For Development
			if device.DevicePolicies.LastUpdate == 0 { // || (time.Now().Unix()-device.DevicePolicies.LastUpdate > 30) { //TODO: Redo This Time Mechanic (Using Config Value) -> Currently Checkin every half a Minute
				log.Debug("Doing Inventory On Device: ", device.UDID)
				//Do Update



				device.DevicePolicies.CurrentAction = structs.DeviceCurrentAction{
					UDID: "UDID For Policy", //Generate This
					// TODO: Maybe Add Bool For If It Is A Policy
					Name: "Inventory",
					Actions: []structs.ServerCommand{}, //TEMP: Put Them Here And Don't Use Append Or Use Shotter Name/Reference
				}

				device.DevicePolicies.CurrentAction.Actions = append(device.DevicePolicies.CurrentAction.Actions, structs.ServerCommand{
					CommandUUID: "741B39F2-649B-4BB6-A522-EE2YF2D99D26",
					Command: structs.ServerPayload{
						RequestType: "DeviceInformation",
						PayloadDeviceInformation: structs.PayloadDeviceInformation{
							Queries: []string{ "BatteryLevel" },
						},
					},
				})

				out, err := plist.MarshalIndent(device.DevicePolicies.CurrentAction.Actions[0], "     "); if err != nil { return 403, err } //TODO: If Possible Remove Indents For Productions
				fmt.Fprintf(w, string(out)) // TODO: This Stuff Is Tiwse (The Plist PArsing) Make That Not A Thing






				//Generate And Add The Policys



				if err := pgdb.Update(&device); err != nil { return 403, err }


				// FIXME: Only Do This Once Checkin Is Done
				//device.DevicePolicies.LastUpdate = time.Now().Unix()
				//if err := pgdb.Update(&device); err != nil { return 403, err }
			}
		}


		//Check Policys and Deploy Them
	}
	*/




	//if checkIn {
		//Get Extra Device Details
		//Get Profiles
		//Get Application
	//}






	return 200, nil //TEMP (Disable Everything Below) //TODO: Next Few Comit This Will Be Redone For Pushing Policies

	//The Code That Goes Here Is In A File Called backup_server_handler Before The Redo (Error Handling Broke It)
}




func deviceInventory(w http.ResponseWriter, r *http.Request, device structs.Device) (int, error) {

	return 200, nil
}

//Handle Devices DOSing The Server When it Keeps Failing -> Prevent It Fast And Alert The Admin

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
