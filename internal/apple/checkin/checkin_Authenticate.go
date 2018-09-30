package checkin

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/mattrax/Mattrax/models"
)

func checkinAuthenticate(w http.ResponseWriter, request models.AppleAuthenticateDetails, db *sqlx.DB, device_exists bool, device models.Device) error {
	log.Println(request)
	//

	if device_exists { // The Device Exists In The Database

	}

	//TODO: Only Macs Can Enroll For Now
	if true { //Allowed To Enroll
		device.Status = 1 // TODO: Put Instead Of This Comment A Link To A Doc Explaining The Status Codes And There Meaning
		device.AppleAuthenticateDetails = request
	}

	//TODO: Handle AwaitingConfiguration For DEP -> Currently Just Bypass It By Sending The Done Payload
	if err := device.UpdateDB(db, `INSERT INTO devices (udid, topic, os_version, build_version, product_name, serial_number, status) VALUES (:udid, :topic, :os_version, :build_version, :product_name, :serial_number, :status)`); err != nil {
		return err
	}

	return nil
}

/* TEMP
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
*/
