/**
 * Mattrax: An Open Source Device Management System
 * File Description: This Is The Apple MDM Checkin URL. it Is accessible From "/apple/checkin" And Is Used For Joining Device And Updating Their APNS Details.
 * Important Notes: In The Apple Docs (And Even Parts of This Code) This ("Inform") is Referred To As Checkin Which Is Its Official Name
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"fmt"
	"net/http"

	//External Deps
	"github.com/groob/plist" //Plist Parsing

	// Internal Functions
	errors "github.com/mattrax/Mattrax/internal/errors" // Mattrax Error Handling

	// Internal Modules
	structs "github.com/mattrax/Mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)

// The "/checkin" route is used to check if the device can join the mdm and update its push token to the server (In The Apple Docs This Is Referred To As The Check-In Route)
func checkinHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	//Parse The Request
	var cmd structs.CheckinCommand
	if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
		return 403, err
	}
	//Attempt To Get The Device From the Database
	var device structs.Device
	if err := pgdb.Model(&device).Where("uuid = ?", cmd.UDID).Select(); err != nil && errors.PgError(err) {
		return 403, err
	}
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
