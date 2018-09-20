package server

import "net/http"

// The Web Handler
func Handler() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		/*
		   var cmd ServerCommand
		   if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
		     log.Debug("Error Parsing Checkin Request: ", err)
		     w.WriteHeader(http.StatusBadRequest)
		     return
		   }
		   device := getDevice(cmd.UDID)

		   if device != nil && device.Deployed {
		     log.Debug("A Device Has Requested The Server: " + device.UDID)
		   } else {
		     log.Warning("A Device Attempted To Get Actions From Server Without Having Send APNS Tokens Yet")
		   }
		*/
		return nil
	}
}
