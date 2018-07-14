/**
 * Mattrax: An Open Source Device Management System
 * File Descripton: This is the Code For Interfacing with APNS.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package apns

import (
  "encoding/hex"
  "encoding/json"
  //"crypto/tls"
  "github.com/RobotsAndPencils/buford/certificate"
  "github.com/RobotsAndPencils/buford/payload"
  "github.com/RobotsAndPencils/buford/push"

  "github.com/sirupsen/logrus" // Logging
	"github.com/go-pg/pg" // Database (Postgres)

  structs "github.com/mattrax/mattrax/appleMDM/structs"
)

var ( log *logrus.Logger; pgdb *pg.DB )
func Init(_pgdb *pg.DB, _log *logrus.Logger) { pgdb = _pgdb; log = _log }

//var cert tls.Certificate

func init() {

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

func DeviceUpdate(_device structs.Device) bool { //TODO: Clean This Up (Maybe Remove Client And Make It Global)
  cert, err := certificate.Load("data/PushCert.p12", "password") //TODO: Load This From Config (Maybe .env)

  if err != nil {
    log.Error(err) //TODO: Should Exit
    return false
  }

  client, err := push.NewClient(cert)
  if err != nil {
    log.Error(err)
    return false
  }

  service := push.NewService(client, push.Production)

  // construct a payload to send to the device:
  p := payload.MDM{
    Token: _device.DeviceTokens.PushMagic,
  }
  b, err := json.Marshal(p)
  if err != nil {
    log.Error(err)
    return false
  }

  // push the notification:
  deviceToken := hex.EncodeToString(_device.DeviceTokens.Token)

  if !push.IsDeviceTokenValid(deviceToken) {
    log.Warning("The Device Token Is Incorrect")
    return false
  }

  id, err := service.Push(deviceToken, nil, b)
  if err != nil {
    log.Error("1", err)
    return false
  }

  log.Debug("APNS ID: " + id)
  return true
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
