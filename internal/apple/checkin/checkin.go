package checkin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/mattrax/Mattrax/models"
)

// The Web Handler
func Handler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		request := models.AppleAuthenticateDetails{}                //TODO: Maybe Move The Model To This Package
		if err := request.PopulateRequestData(r.Body); err != nil { //TODO: Try And MOve This Error Handling Centeral Based On If The Error Is From Plist/Model Giving Client Errors
			log.Println("Error Parsing Checkin Request: ", err)
			http.Error(w, "Error Parsing The Data From The Device", 404) // TODO: Get Correct 4xx Error Code For This
			return nil                                                   // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above //TODO: Stop Writting This Same Thing And Link To A Github Issue
		}

		device := models.Device{}
		if request.UDID == "" {
			http.Error(w, "Invalid Request", 401) //TODO: Replace With Returned Central Handling/Logging
			return nil
		}
		device.AppleAuthenticateDetails.UDID = request.UDID

		// Send It To The Function Handling Each Checkin Action
		switch request.DeviceRequest.MessageType { //TODO: Check MessageType Varible Exists
		case "Authenticate": //TODO: Make Sure All The Requires Fields In The Request Are There
			// Load The Device From The Database
			err := device.LoadFromDB(db)
			if err != nil && err != sql.ErrNoRows {
				return err
			}

			// Call The Function To Verify & Enroll/Reject The Device
			return checkinAuthenticate(w, request, db, err == sql.ErrNoRows, device)
		case "TokenUpdate":
			// Load The Device From The Database
			err := device.LoadFromDB(db)
			if err != sql.ErrNoRows {
				http.Error(w, "Device Not Found", 401) //TODO: Replace With Returned Central Handling/Logging
				return nil
			} else if err != nil {
				return err
			}

			// Call The Function To Verify & Enroll/Reject The Device
			return checkinTokenUpdate(w, request, db, err == sql.ErrNoRows, device)
		case "CheckOut":
			//TODO: CheckOut
		default:
			http.Error(w, "That Operation Is Not Supported By Mattrax", 500) //TODO: Get Correct 4xx Error Code For This
			return nil                                                       // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above
		}

		/*
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
					// TEMP Bypass
					device.DeviceState = 3
					if err := pgdb.Update(&device); err != nil {
						return 403, err
					}
					log.Info("A New Device Joined The MDM: " + device.UDID)
					// End Bypass

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
		*/
		return nil
	}
}

// Parse The Plist Data Coming From The Device
/*var cmd models.CheckinRequest
if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { //TODO: Does It Handle A Blank BOdy?
	log.Println("Error Parsing Checkin Request: ", err)
	http.Error(w, "Error Parsing The Data From The Device", 500) // TODO: Get Correct 4xx Error Code For This
	return nil                                                   // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above
}

log.Println(cmd)*/

/*var cmd models.CheckinCommand
if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { //TODO: Does It Handle A Blank BOdy?
	log.Println("Error Parsing Checkin Request: ", err)
	http.Error(w, "Error Parsing The Data From The Device", 500) // TODO: Get Correct 4xx Error Code For This
	return nil                                                   // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above
}*/

// Retrieve The Device From The Database (The Output Could Be Nil If One Is Not Found)
