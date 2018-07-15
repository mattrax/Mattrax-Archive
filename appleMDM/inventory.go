/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is The Apple MDM Core. It Manages The Webserver Routes and APNS for Apples MDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
	"net/http"
	"time"
  "errors"

	// Internal Modules
	structs "github.com/mattrax/Mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)

func doInventory(w http.ResponseWriter, r *http.Request, cmd structs.Response, device structs.Device) (int, error){
  currentState := device.DevicePolicies.Inventory.State

  // What Actions Are Done:
  //	ProfileList
  //	ProvisioningProfileList
  //	InstalledApplicationList
  //	DeviceInformation
  //	SecurityInfo
  //	Restrictions -> Maybe

  //TODO; Nearter Deployment System. Maybe Get Payloads Type From Array And Auto Do The Rest In Loop

  if currentState == 0 {
    log.Warning("Doing Inventory on Device: ", device.UDID)

    log.Warning("Deployed 1")
    if payload, err := structs.NewPayload(&structs.Command{ RequestType: "ProfileList" }); err != nil {
      return 403, err
    } else {
      device.DevicePolicies.Inventory.CommandUUIDs[payload.CommandUUID] = "ProfileList"
      //device.DevicePolicies.Inventory.CommandUUIDs[0] = payload.CommandUUID //TODO: Will This Cause Issues ????
      if err := returnPlist(w, payload); err != nil { return 403, err}
    }

  } else if currentState == 1 {

    log.Warning("Deployed 2")
    if payload, err := structs.NewPayload(&structs.Command{ RequestType: "ProvisioningProfileList" }); err != nil {
      return 403, err
    } else {
      device.DevicePolicies.Inventory.CommandUUIDs[payload.CommandUUID] = "ProvisioningProfileList"
      //device.DevicePolicies.Inventory.CommandUUIDs[1] = payload.CommandUUID //TODO: Will This Cause Issues ????
      if err := returnPlist(w, payload); err != nil { return 403, err}
    }


  } else if currentState == 2 { //Inventory Finished Cleanup
    log.Warning("Inventory Finished For Device: ", device.UDID)
    device.DevicePolicies.Inventory.State = 0
    device.DevicePolicies.Inventory.LastUpdate = time.Now().Unix()

    if err := pgdb.Update(&device); err != nil { return 403, err }
    return 200, nil
  } else {
    return 200, errors.New("An Invalid Inventory State Of '" + string(currentState) + "' Has Been Reached On Device: " + cmd.UDID) // TODO: Try And Save Daatabase By Taking One For Entry maybe
  }
  device.DevicePolicies.Inventory.State++
  if err := pgdb.Update(&device); err != nil { return 403, err }
  return 200, nil //TODO: Device Will DOS If Failed So Be Carefull Failing
}
