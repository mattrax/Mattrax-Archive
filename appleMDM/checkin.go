package appleMDM

import (
	"fmt"
	"net/http"

	"github.com/groob/plist"

	"encoding/hex"
	"encoding/json"
	"github.com/RobotsAndPencils/buford/certificate"
	"github.com/RobotsAndPencils/buford/payload"
	"github.com/RobotsAndPencils/buford/push"
)

// The "/checkin" route is used to check if the device can join the mdm and update its push token to the server
func checkinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Checkin Request")

	/*buf, err := ioutil.ReadAll(r.Body)
	  r.Body.Close()
	  if err != nil {
	      //Do something
	      fmt.Println(err)
	  }

	  fmt.Println(string(buf))

	  w.WriteHeader(http.StatusOK)

	  return*/

	var cmd CheckinCommand
	if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
		fmt.Println("Failed To Parse Checkin Request")
		fmt.Println(err)

		// TODO: Debug Event To Error Logs
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if cmd.MessageType == "Authenticate" {
		if cmd.auth.OSVersion != "" && cmd.auth.BuildVersion != "" && cmd.auth.ProductName != "" && cmd.auth.SerialNumber != "" && cmd.auth.IMEI != "" && cmd.auth.MEID != "" {
			fmt.Println("Authenticate Request")
			fmt.Println("Device Type: " + cmd.auth.ProductName)

			if true { //Device is Allowed To Join
				w.WriteHeader(http.StatusOK)
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else {
			fmt.Println("Authentication Request With Invalid Device (Pre IOS 9/Doesn't Have Perms)")

			// TODO: Debug Event To Error Logs (This Is A Pre IOS 9 or Doesn't Have Device Information access right in the profile)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

	} else if cmd.MessageType == "TokenUpdate" {
		fmt.Println("TokenUpdate Request")

		if cmd.update.Token != nil && cmd.update.PushMagic != "string" && cmd.update.UnlockToken != nil && (cmd.update.AwaitingConfiguration == true || cmd.update.AwaitingConfiguration == false) {
			if cmd.update.AwaitingConfiguration == true {
				//Do DEP Prestage Enrollment By Pushing The Profiles Now Then Run The Fully Setup Thing
				//TODO: Debug Event (Until Feature is Built)
				fmt.Println("Unsupported DEP Features")
				//TODO: Future Feature: If set to true, the device is awaiting a DeviceConfigured MDM command before proceeding through Setup Assistant.
			}
			if cmd.update.UnlockToken != nil {
				fmt.Println("Device Sent Unlock Token To Deal With")
			}
			fmt.Println("Device Updated Its Tokens")

			/* Contact APNS */

			cert, err := certificate.Load("PushCert.p12", "password")
			exitOnError(err)

			client, err := push.NewClient(cert)
			exitOnError(err)

			service := push.NewService(client, push.Production)

			// construct a payload to send to the device:
			p := payload.MDM{
				Token: cmd.update.PushMagic,
			}
			b, err := json.Marshal(p)
			exitOnError(err)

			// push the notification:
			deviceToken := hex.EncodeToString(cmd.update.Token)

			fmt.Println(cmd.update.Token) //"Device Token: " +

			if !push.IsDeviceTokenValid(deviceToken) {
				fmt.Println("The Device Token Is Incorrect")
				return
			}

			id, err := service.Push(deviceToken, nil, b)
			exitOnError(err)

			fmt.Println("apns-id:", id)

			/*
			   if e, ok := err.(*push.Error); ok {
			   	switch e.Reason {
			   	case push.ErrBadDeviceToken:
			   		// handle error
			   	}
			   }
			*/

			/* End Contact APNS */

			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println("Error: Incorrect Data For The TokenUpdate")
			//TODO: Debug Event
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if cmd.MessageType == "CheckOut" {
		fmt.Println("CheckOut Request")

		//TODO: Future Feature
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println("Invalid MessageType")

		// TODO: Debug Event To Error Logs
		w.WriteHeader(http.StatusInternalServerError) //TODO: Chnage HTTP Error Returned To Client Error
	}

	//To Force Enrollment To Fail For Development
	//w.WriteHeader(http.StatusInternalServerError)
	//w.Write([]byte("500 - Something bad happened!"))

	//fmt.Fprintf(w, "Apple MDM Checkin")
}
