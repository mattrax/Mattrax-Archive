/**
 * Mattrax: An Open Source Device Management System
 * File Description: This is the Code For Interfacing with APNS.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package apns

import (
	"encoding/hex"
	"encoding/json"

	//External Deps
	"github.com/RobotsAndPencils/buford/certificate" // Apple Push Notification Service -> Certificates
	"github.com/RobotsAndPencils/buford/payload"     // Apple Push Notification Service -> Payloads
	"github.com/RobotsAndPencils/buford/push"        // Apple Push Notification Service -> Push

	errors "github.com/mattrax/Mattrax/internal/errors" // Mattrax Error Handling
	// Internal Functions
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
	mdb "github.com/mattrax/Mattrax/internal/database"      //Mattrax Database
	mlg "github.com/mattrax/Mattrax/internal/logging"       //Mattrax Logging

	// Internal Modules
	structs "github.com/mattrax/Mattrax/appleMDM/structs" // Apple MDM Structs/Functions
)

var pgdb = mdb.GetDatabase()
var log = mlg.GetLogger()
var config = mcf.GetConfig() // Get The Internal State

func init() {
	//Load The Certficate From where The Config Said
	log.Debug(config.Name) //This Is To Not Cause Unused var Error
}

/*
func ExInit() { // TODO: Fix This - The Old not Functioal init()
  var err error
  cert, err = certificate.Load("PushCert.p12", "password") //TODO: Load This From Config (Maybe .env)

  if err != nil {
    fmt.Println(err) //TODO: Should Exit
  }

  fmt.Println("Loaded The APNS Certificate") //Logging Hasn;t Loaded Yet Make This Later
}*/
/*
func loadAPNSCertificate(certFile string, password string) *tls.Certificate {
  cert, err := certificate.Load(certFile, password)

  if err != nil {
    log.Fatal(err)
    return nil
  }

  return cert
}
*/

func DeviceUpdate(_device structs.Device) error { //TODO: Clean This Up (Maybe Remove Client And Make It Global)
	cert, err := certificate.Load("data/PushCert.p12", "password") //TODO: Load This From Config (Maybe .env)

	if err != nil {
		return err //TODO: Should Exit
	}

	client, err := push.NewClient(cert)
	if err != nil {
		return err
	}

	service := push.NewService(client, push.Production)

	// construct a payload to send to the device:
	p := payload.MDM{
		Token: _device.DeviceTokens.PushMagic,
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// push the notification:
	deviceToken := hex.EncodeToString(_device.DeviceTokens.Token)

	if !push.IsDeviceTokenValid(deviceToken) {
		log.Warning("The Device Token Is Incorrect")
		return errors.New("Invalid Token")
	}

	id, err := service.Push(deviceToken, nil, b)
	if err != nil {
		return err
	}

	log.Debug("APNS ID: " + id)
	return nil
}

/*cert, err := certificate.Load("PushCert.p12", "password")
if err != nil {
  log.Error(err)
  w.WriteHeader(http.StatusBadRequest)
  return
}

client, err := push.NewClient(cert)
if err != nil {
  log.Error(err)
  w.WriteHeader(http.StatusBadRequest)
  return
}

service := push.NewService(client, push.Production)

// construct a payload to send to the device:
p := payload.MDM{
  Token: device.PushMagic,
}
b, err := json.Marshal(p)
if err != nil {
  log.Error(err)
  w.WriteHeader(http.StatusBadRequest)
  return
}

// push the notification:
deviceToken := hex.EncodeToString(device.Token)

if !push.IsDeviceTokenValid(deviceToken) {
  fmt.Println("The Device Token Is Incorrect")
  return
}

id, err := service.Push(deviceToken, nil, b)
if err != nil {
  log.Error(err)
  w.WriteHeader(http.StatusBadRequest)
  return
}

log.Debug("APNS ID: " + id)*/
