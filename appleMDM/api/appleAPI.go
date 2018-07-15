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
	"github.com/mattrax/mattrax/appleMDM/apns" // Apple Push Notification Service

	// Internal Functions
	mcf "github.com/mattrax/mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/mattrax/internal/database"      //Mattrax Database
	mlg "github.com/mattrax/mattrax/internal/logging"       //Mattrax Logging

	// Internal Modules
	structs "github.com/mattrax/mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)

var pgdb = mdb.GetDatabase()
var log = mlg.GetLogger()
var config = mcf.GetConfig() // Get The Internal State

/* API Endpoints */

func Mount(r *mux.Router) {
	r.HandleFunc("/apns", pingApnsHandler).Methods("GET")
}

func pingApnsHandler(w http.ResponseWriter, r *http.Request) { // TEMP: This And APNS Handling Needs Redoing
	devices := structs.GetDevices()

	if devices == nil { //[]Device{} {
		log.Debug("Error Getting Devices Or There Are None")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error Getting Devices Or There Are None")
		return
	}

	for _, device := range devices {
		if device.DeviceState == 3 {
			log.Debug("APNS Update Sent To Device " + device.UDID)
			status := apns.DeviceUpdate(device)

			if !status { //Custom Error Handling (Detect Unenrolled Devices)

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
