/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is the REST API For The Apple MDM Server. This Is Used By The Web Interface To Interface With The Backend.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleAPI

import (
	"fmt"
	"net/http"
	"encoding/json"

	// External Deps
	"github.com/gorilla/mux"                   // HTTP Router
	"github.com/mattrax/Mattrax/appleMDM/apns" // Apple Push Notification Service

	// Internal Functions
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database"      //Mattrax Database
	mlg "github.com/mattrax/Mattrax/internal/logging"       //Mattrax Logging

	// Internal Modules
	structs "github.com/mattrax/Mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)

var ( // Get The Internal State
	pgdb = mdb.GetDatabase()
	log = mlg.GetLogger()
	config = mcf.GetConfig()
)

/* API Endpoints */

func Mount(r *mux.Router) {
	r.HandleFunc("/apns", pingApnsHandler).Methods("GET")
}

func pingApnsHandler(w http.ResponseWriter, r *http.Request) { // TEMP: This And APNS Handling Needs Redoing
	var devices []structs.Device
	_ = pgdb.Model(&devices).Select() //TODO: Error Handling

	if devices == nil { //[]Device{} {
		log.Debug("Error Getting Devices Or There Are None")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error Getting Devices Or There Are None")
		return
	}

	for _, device := range devices {
		if device.DeviceState == 3 {
			if err := apns.DeviceUpdate(device); err != nil { //TODO Custom Error Handling (Detect Unenrolled Devices)
				log.Debug(err)
				//fmt.Fprintf(w, "Error Sending APNS Update To The Device: " + device.UDID)

				/*log.WithFields(mlg.Fields{
			    "udid": device.UDID,
					"DeviceState": device.DeviceState,
					"DeviceDetails": device.DeviceDetails,
			  }).Warning("Error Sending APNS Update")

				log.WithFields(mlg.Fields{
			    "udid": device.UDID,
					"DeviceDetails": device.DeviceDetails,
			  }).Debug("Error Sending APNS Update")*/


				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(JSONStatus{
					Success: true,
					Message: "Error Sending APNS Update To Device",
					Device: &structs.Device{
						UDID: device.UDID,
						DeviceState: device.DeviceState,
						DeviceDetails: device.DeviceDetails,
					},
				})

				return
			}
		}
	}

	json.NewEncoder(w).Encode(JSONStatus{
		Success: true,
		Message: "All Devices Have Been Told To Update",
	})
}


type JSONStatus struct { //TODO: Move To Struts Package
	Success bool `json:"status"`
	Message string `json:"message"`
	Device *structs.Device `json:"device,omitempty"`
}
