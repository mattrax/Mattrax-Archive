/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is A File Of Function To Do Repeatable Actions With The Stucts.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleStruct

import (
	//External Deps
	"github.com/go-pg/pg" // Database (Postgres)
	"github.com/satori/go.uuid" // UUID Generation
	"github.com/groob/plist" //Plist Parsing

	// Internal Functions
	mdb "github.com/mattrax/Mattrax/internal/database" //Mattrax Database
	mlg "github.com/mattrax/Mattrax/internal/logging"  //Mattrax Logging


)

var pgdb = mdb.GetDatabase()
var log = mlg.GetLogger() // Get The Internal State

///// Devices Functions /////
func NewDevice(cmd CheckinCommand) Device {
	return Device{
		UDID:        cmd.UDID,
		DeviceState: 0, //cmd.DeviceState,
		DeviceDetails: DeviceDetails{
			OSVersion:    cmd.Auth.OSVersion,
			BuildVersion: cmd.Auth.BuildVersion,
			ProductName:  cmd.Auth.ProductName,
			SerialNumber: cmd.Auth.SerialNumber,
			IMEI:         cmd.IMEI,
			MEID:         cmd.MEID,
		},
		DeviceTokens: DeviceTokens{
			Token:       []byte{},
			PushMagic:   "",
			UnlockToken: []byte{},
		},
		DevicePolicies: DevicePolicies{
			Queued:    []string{},
			Installed: []string{},
		},
	}
}

///// Policies Functions /////
/*func getPolicy(uuid string) {

}*/

func ParsePolicy(policy Policy) (string, error) {

	// Returns The XML Output After Parsing The Inputted Policy
	return "hello world", nil
}

//TODO: Redo Error Hanling For This File. eg. Use: "err != nil && ierror.PgError(err) { return 403, err }"
//TODO: Keep This Line For Later: if err := pgdb.Delete(&device); err != nil && ierror.PgError(err) { return 405, err }

/*
if policy.Config.PolicyType == "InstallApplication" {


  AppPayload := ServerCommand{
    CommandUUID: "4424F929-BDD2-4D44-B518-393C0DABD56A", //TODO: Build Generator For These
    Command: ServerPayload{
      RequestType: "InstallApplication",
      PayloadInstallApplication: policy.Options.PayloadInstallApplication,
    },
  }

  out, err := plist.MarshalIndent(AppPayload, "     ") //TODO: Clean This Plist Parsing And Error Handling (And Other Ones Using The Same Code)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Fprintf(w, string(out))

  // Move Out Of The Queue

  device.DevicePolicies.Queued[index] //Delete It
  //Add To The installed


} else {
  fmt.Fprintf(w, "")
}
*/

/* End */

// getDevice()

/*func getDevice(_UDID string) (Device, error) {
  var device Device
  err := pgdb.Model(&device).Where("uuid = ?", _UDID).Select()
  if err != nil { return err }
  return device, nil
}*/

/*if err == pg.ErrNoRows || err == pg.ErrMultiRows {
  log.Debug("getDevice(): Searching Empty Database");
  return nil
} else
if err != nil {
  if err != pg.ErrNoRows && err != pg.ErrMultiRows {
    log.Warning("Postgres Error: ", err);
     //TODO: Try Database Request Again Here
  }

  return nil
}*/

func GetDevices() []Device {
	var devices []Device
	err := pgdb.Model(&devices).Select()
	if err != nil {
		log.Error(err)
		return []Device{}
	}
	return devices
}

func EditDevice(_device *Device, exists bool) bool {
	if _device == nil {
		log.Debug("editDevice() Parsed A Nil Device")
		return false
	}
	var err error
	if exists {
		err = pgdb.Update(_device)
	} else {
		//Create New
		_, err = pgdb.Model(_device).
			Set("uuid = ?", _device.UDID).
			Insert()
	}

	if err != nil {
		if err != pg.ErrNoRows && err != pg.ErrMultiRows {
			log.Warning("Postgres Error (Exists: ", exists, "): ", err)
			//TODO: Try Database Request Again Here
		}
		return false
	}
	return true
}

func DeleteDevice(_device **Device) bool {

	err := pgdb.Delete(&_device)

	//Eror Handle nil _device

	//err := pgdb.Delete(&_device)

	/*out, err := pgdb.
	  Model(Device{
	    UDID: _device.UDID,
	  }).
	  //Where("uuid = ?", _device.UDID).
	  //Select().
	  //Set("uuid = ?", _device.UDID).
	  //Select().
	  Delete() //*_device

	log.Debug(out)*/

	if err != nil {
		if err != pg.ErrNoRows && err != pg.ErrMultiRows {
			log.Warning("Postgres Error: ", err)
			return false
		} else {
			return true
		}
	}
	return true
}

















/* /Server Handling */
func ParsePayload(_payload ServerPayload) (string, error) {
	randomUUID, _ := uuid.NewV4() //TODO Erro hadnling
	payload := ServerCommand{
		CommandUUID: randomUUID.String(),
		Command: _payload,
	}

	out, err := plist.MarshalIndent(payload, "     "); if err != nil { return "", err } //TODO: If Possible Remove Indents For Productions
	return string(out), nil
}
