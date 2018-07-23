/**
 * Mattrax: An Open Source Device Management System
 * File Description: This Is The Apple MDM Server URL. it Is accessable From "/apple/server" And Is Used For Sending Payloads/Commands To Devices And Getting Thier Responses.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package appleMDM

import (
  //"encoding/json"
  "encoding/hex" // TODO Elminate This By Storing In Structs And DB As It
	"fmt"
	"net/http"
  "time"

	//External Deps
	"github.com/groob/plist" //Plist Parsing
  apns "github.com/RobotsAndPencils/buford/push" //APNS

	// Internal Functions
	errors "github.com/mattrax/Mattrax/internal/errors"     // Mattrax Error Handling

	// Internal Modules
	structs "github.com/mattrax/Mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)




var temp_done = false //TEMP

//TODO: Redo All Exit HTTP Codes
//TODO: Redo Exit Resaving DB Handler (Maybe Add it To The errors Router Controller)
//TODO: Func Description -> Make Sure It Doesn't DOS Itself By Returning 403 Lots
func serverHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	//Parse The Request
	var cmd structs.Response
	if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { return 403, err }
	//Attempt To Get The Device From the Database
	var device structs.Device
	if err := pgdb.Model(&device).Where("uuid = ?", cmd.UDID).Select(); err != nil && errors.PgError(err) { return 403, err }
	//Handle The Request
	if device.DeviceState != 3 { return 200, errors.New("A Device That is Not Known To The MDM Tried To /server") } //TODO: Redo Message
  if !apns.IsDeviceTokenValid(hex.EncodeToString(device.DeviceTokens.Token)) { return 200, errors.New("The APNS Token Is Invalid For Device: " + device.UDID )} //TODO: Report To Admin (Maybe Add New DeviceState Of Lost (When It Stops COmmunicating))
  //Handle Feedback From The Device
  if cmd.Status != "Acknowledged" {
    //Save Output And It Worked
  } else if cmd.Status != "Error" {
    //Save Output And It Failed
  } else if cmd.Status != "CommandFormatError" {
    log.Error("The Command Sent To The Device Was Malformed")

    /*log.WithFields(logrus.Fields{ //TODO: Fill With Lots Of Debugging Details
	    "animal": "walrus",
	    "size":   10,
	  }).Error("The Command Sent To The Device Was Malformed")*/


  } else if cmd.Status != "NotNow" {
    log.Fatal("This Isn't Working Yet") //TODO: Make This Mechanic
  } else if cmd.Status != "Idle" {
    log.Debug("The Device Is Idle")
  } else {
    //Handle Invalud Status From The Device
  }












	if cmd.Status == "Acknowledged" {
		if device.DevicePolicies.Inventory.Commands[cmd.CommandUUID] != "" { //TODO: Cehck false isn't Returned By Blank
			requestType := device.DevicePolicies.Inventory.Commands[cmd.CommandUUID] //TODO: Rename This var

			log.Warning("Inventory Req: ", cmd.CommandUUID)


			//Switch Here -> requestType

			if requestType == "ProfileList" {
				device.DeviceDetails.Profiles = cmd.ProfileList


			} else {
				log.Warning("Don't Know How to Handle That")
			}



			delete(device.DevicePolicies.Inventory.Commands, cmd.CommandUUID)
		} else {
			log.Warning("Authentication For Non Inventory Thingo")

			//TODO




		}
	}
  //Handle Other Status's
  log.Warning(cmd.Status)

  if cmd.Status == "Fail" {
    log.Error(cmd)
  }


  //TODO: Handle Policy Updates Before Inventory












  //Policys
  for uuid, devicePolicy := range device.DevicePolicies.Queued {
    if devicePolicy.Status == 0 {
      //Get The Policy From The Database
      var policy structs.Policy
    	if err := pgdb.Model(&policy).Where("uuid = ?", uuid).Select(); err != nil && errors.PgError(err) { return 403, err }
      //Generate And Push The Policy
      log.Info("Pushing Policy '" + policy.Config.Name + "' To Device '" + device.UDID + "'")

      var command = policy.Command
      command.RequestType = policy.Config.PolicyType
      if payload, err := structs.NewPayload(&command); err != nil {
        return 200, err
      } else {
        if err := returnPlist(w, payload); err != nil {
          return 200, err
        } else {
          log.Warning("Deployed")

          devicePolicy.Status = 1
          device.DevicePolicies.Queued[uuid] = devicePolicy
          if err := pgdb.Update(&device); err != nil && errors.PgError(err) { return 403, err } //TODO: Do I NEed erros.PgError For This Type Of Request
          return 200, nil


          //device.DevicePolicies.Queued[i].Status = 1


          //devicePolicy.Status = 1
          //if err := pgdb.Update(&devicePolicy); err != nil && errors.PgError(err) { return 403, err }

        }
        //return 200, returnPlist(w, payload)
      }







      break
    }
  }



  //Move To Installed Once Feedback is Recieved
  //delete(device.DevicePolicies.Queued, devicePolicy.UUID)
  //device.DevicePolicies.Queued[devicePolicy.UUID]
  //Status = 0

  //if err := pgdb.Update(&device); err != nil && errors.PgError(err) { return 403, err } //TODO: Do I NEed erros.PgError For This Type Of Request

  return 200, nil









  /*
  queuedPolicies := len(device.DevicePolicies.Queued)
  if queuedPolicies > 0 {
    for i, devicePolicy := range queuedPolicies {





    }


  } else { //TEMP Else And Content
    log.Warning("No Policies To Push To Device")
  }
  */





  /*
  for i, devicePolicy := range device.DevicePolicies.Queued { // TODO: Don't Use This It Would Push Multiple Policys In One Request
    if devicePolicy.Status == 0 {
      var policy structs.Policy
    	if err := pgdb.Model(&policy).Where("uuid = ?", devicePolicy.UUID).Select(); err != nil && errors.PgError(err) { return 403, err }

      //TODO: Check Device And Target Match to Be Safe
      //TODO: Handle No Device Being Returned -> Use The Custom If Erro handling Thingo

      log.Info("Pushing Policy '" + policy.Config.Name + "' To Device '" + device.UDID + "'")

      var command = policy.Command
      command.RequestType = policy.Config.PolicyType
      if payload, err := structs.NewPayload(&command); err != nil {
        return 200, err
      } else {
        if err := returnPlist(w, payload); err != nil {
          return 200, err
        } else {

          //device.DevicePolicies.Queued[i].Status = 1


          //devicePolicy.Status = 1
          //if err := pgdb.Update(&devicePolicy); err != nil && errors.PgError(err) { return 403, err }

        }
        //return 200, returnPlist(w, payload)
      }




      //Save devicePolicy

    } else {
      log.Warning("Already Sent To Device") //TEMP
    }

    return 200, nil
    //Up


    log.Info(devicePolicy.UUID)
    log.Info(devicePolicy.Status)

    //Get Policies Status




    //



  }*/

  /*if (time.Now().Unix()-device.DevicePolicies.LastUpdate > 30) { //TODO: Redo Timer Mechanic And Set A Sane Default. Allow Changing Through config.json (Bigger Deployments Will Need it)
		return doInventory(w, r, cmd, device)
	} else {
		if err := pgdb.Update(&device); err != nil { return 403, err }
	  return 200, nil
	}*/








    /*
    var policy structs.Policy
    if err := pgdb.Model(&policy).Where("uuid = ?", policy_name).Select(); err != nil { return 403, err }
    //TODO: Check PolicY.PolicyConfig.Target Is Valid With Device





    var command = policy.Command
    command.RequestType = "InstallApplication"
    if payload, err := structs.NewPayload(&command); err != nil {
      return 403, err
    } else {
      log.Info(payload.Command.InstallApplication.ITunesStoreID)
      return 200, returnPlist(w, payload)
    }






    //Track Status Of Policy Deployment 0 Not Sent, 1 Sent, 2 Got Reply















    return 200, nil


    */

	//Check For Update Policies To Update

	return 200, nil
}

func returnPlist(w http.ResponseWriter, payload *structs.Payload) error {
	plistCmd, err := plist.MarshalIndent(payload, "\t")
	if err != nil { return err }


	log.Info(string(plistCmd)) //TEMP
	fmt.Fprintf(w, string(plistCmd)) //TODO: Can This be Done Without string() (Via Streams To Make It More Efficent)
	return nil
}

func deleteSlice(a []string, i int) []string {
  copy(a[i:], a[i+1:])
  a[len(a)-1] = ""
  return a[:len(a)-1]
}





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
      device.DevicePolicies.Inventory.Commands[payload.CommandUUID] = "ProfileList"
      //device.DevicePolicies.Inventory.Commands[0] = payload.CommandUUID //TODO: Will This Cause Issues ????
      if err := returnPlist(w, payload); err != nil { return 403, err}
    }

  } else if currentState == 1 {

    log.Warning("Deployed 2")
    if payload, err := structs.NewPayload(&structs.Command{ RequestType: "ProvisioningProfileList" }); err != nil {
      return 403, err
    } else {
      device.DevicePolicies.Inventory.Commands[payload.CommandUUID] = "ProvisioningProfileList"
      //device.DevicePolicies.Inventory.Commands[1] = payload.CommandUUID //TODO: Will This Cause Issues ????
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
