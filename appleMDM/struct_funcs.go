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
  err := pgdb.Model(&device).Where("uuid = ?", _UDID).Select()

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












func getDevices() []Device {
  var devices []Device
  err := pgdb.Model(&devices).Select()
  if err != nil {
		log.Error(err)
		return []Device{}
	}
  return devices
}

func newDevice(cmd CheckinCommand) *Device {
  return &Device{
    UDID: cmd.UDID,
    DeviceState: 0,
    DeviceDetails: DeviceDetails{
      OSVersion: cmd.auth.OSVersion,
      BuildVersion: cmd.auth.BuildVersion,
      ProductName: cmd.auth.ProductName,
      SerialNumber: cmd.auth.SerialNumber,
      IMEI: cmd.IMEI,
      MEID: cmd.MEID,
    },
    DeviceTokens: DeviceTokens{
      Token: []byte{},
      PushMagic: "",
      UnlockToken: []byte{},
    },

    /*UDID: cmd.UDID,
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
    Deployed: false,*/
  }
}

func editDevice(_device *Device, exists bool) bool {
  if _device == nil  {
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
      log.Warning("Postgres Error (Exists: ", exists, "): ", err);
       //TODO: Try Database Request Again Here
    }
    return false
  }
  return true
}


func deleteDevice(_device **Device) bool {

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
      log.Warning("Postgres Error: ", err);
       return false
    } else {
      return true
    }
  }
  return true
}
















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
