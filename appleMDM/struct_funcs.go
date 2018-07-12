/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is Script Of Function To Do Repeatable Actions With The Stucts.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
  "github.com/go-pg/pg" // Database (Postgres)
)

func getDevice(_UDID string) *Device {
  var device Device
  err := pgdb.Model(&device).Where("udid = ?", _UDID).Select()

  /*if err == pg.ErrNoRows || err == pg.ErrMultiRows {
    log.Debug("getDevice(): Searching Empty Database");
    return nil
  } else*/
  if err != nil {
    if err != pg.ErrNoRows && err != pg.ErrMultiRows {
      log.Warning("Postgres Error: ", err);
       //TODO: Try Database Request Again Here
    }
    return nil
  }
  return &device
}

func getDevices() {












}

func newDevice(cmd CheckinCommand) *Device {
  return &Device{
    UDID: cmd.UDID,
    // Device Details
    OSVersion: cmd.auth.OSVersion,
    BuildVersion: cmd.auth.BuildVersion,
    ProductName: cmd.auth.ProductName,
    SerialNumber: cmd.auth.SerialNumber,
    IMEI: cmd.IMEI,
    MEID: cmd.MEID,
    // APNS
    Token: []byte{},
    PushMagic: "",
    UnlockToken: []byte{},
    //Status
    Deployed: false,
  }
}

func editDevice(_device *Device, exists bool) bool {
  var err error
  if exists {
    err = pgdb.Update(_device)
  } else {
    //Create New
    _, err = pgdb.Model(_device).
      Set("udid = ?", _device.UDID).
      Insert()
  }

  if err != nil {
    if err != pg.ErrNoRows && err != pg.ErrMultiRows {
      log.Warning("2Postgres Error: ", err);
       //TODO: Try Database Request Again Here
    }
    return false
  }
  return true

  /*_, err := pgdb.Model(_device).
    //OnConflict("(udid) DO UPDATE").
    Where("udid = ?", _device.UDID).
    Update()*/

  //log.Info(_device)

  //.Where("udid = ?", _device.UDID).OnConflict("(udid) DO UPDATE").Insert() //. Returning("udid")

  /*err := pgdb.Model(_device)OnConflict("(id) DO UPDATE").
        Set("udid = ?", _device.UDID).
        Create()
  */
        //.OnConflict("(udid) DO UPDATE").Create()
  //.Model(_device).
    //OnConflict("(udid) DO UPDATE").
    //Set("udid = ?", _device.UDID).


    //Column("id").
    //Where("udid = ?", _device.UDID).
    //OnConflict("DO NOTHING"). // OnConflict is optional
    //Returning("id").
    //SelectOrCreate()
    /*OnConflict("(udid) DO UPDATE").
    Set("udid = ?", _device.UDID).
    Insert()*/

  if err != nil {
    log.Warning("Postgres Error: ", err);
    return false //TODO: Try Again Here
  }

  //log.Info("After Update: ", getDevice(_device.UDID).Deployed)
  return true
}


func deleteDevice() {

}

//update/add device
