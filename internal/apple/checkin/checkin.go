package checkin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/mattrax/Mattrax/models"
)

// The Web Handler
func Handler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		device := models.Device{} //TODO: Move THis Model Back TO This Package

		request := models.Device{}
		if err := request.PopulateRequestData(r.Body); err != nil { //TODO: Try And MOve This Error Handling Centeral Based On If The Error Is From Plist/Model Giving Client Errors
			log.Println("Error Parsing Checkin Request: ", err)
			http.Error(w, "Error Parsing The Data From The Device", 404) // TODO: Get Correct 4xx Error Code For This
			return nil                                                   // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above //TODO: Stop Writting This Same Thing And Link To A Github Issue
		}

		cleanup := func() {
			_, err := db.NamedExec("INSERT INTO devices VALUES (:udid, :topic, :os_version, :build_version, :product_name, :serial_number, :imei, :meid, :token, :push_magic, :unlock_token, :awaiting_configuration) ON CONFLICT (udid) DO UPDATE SET push_magic = EXCLUDED.push_magic, unlock_token = EXCLUDED.unlock_token;", device)
			if err != nil {
				log.Println(err)
				panic(err) //TODO: Better Error Handling
			}
		}

		// Send It To The Function Handling Each Checkin Action
		switch device.DeviceRequest.MessageType { //TODO: Check MessageType Varible Exists
		case "Authenticate":
			if err := device.LoadFromDB(db); err != nil && err != sql.ErrNoRows {
				return err
			}

			if true { //Allowed To Enroll
				device = request
			}

			log.Println(device)

			defer cleanup()
			return checkinAuthenticate(w, r)
		case "TokenUpdate":
			/*topic := device.Topic
			udid := device.UDID
			token := device.Token
			push_magic := device.PushMagic
			unlock_token := device.UnlockToken
			awaiting_configuration := device.AwaitingConfiguration*/
			if err := device.LoadFromDB(db); err != nil && err != sql.ErrNoRows {
				return err
			}
			device.PushMagic = request.PushMagic

			defer cleanup()
			return checkinTokenUpdate(w, r)
		case "CheckOut":
			if err := device.LoadFromDB(db); err != nil && err != sql.ErrNoRows {
				return err
			}
			return checkout(w, r)
		default:
			http.Error(w, "That Operation Is Not Supported By Mattrax", 500) //TODO: Get Correct 4xx Error Code For This
			return nil                                                       // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above
		}

		/*
		    var cmd CheckinCommand
		    if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
		      log.Debug("Error Parsing Checkin Request: ", err)
		      w.WriteHeader(http.StatusBadRequest)
		      return
		    }
		  	device := getDevice(cmd.UDID)

		    if cmd.MessageType == "Authenticate" {
		      if cmd.auth.OSVersion != "" && cmd.auth.BuildVersion != "" && cmd.auth.ProductName != "" && cmd.auth.SerialNumber != "" && cmd.auth.IMEI != "" && cmd.auth.MEID != "" {
		  			if device == nil {
		  				device = newDevice(cmd)

		  				if status := editDevice(device, false); status == false {
		  	        log.Debug("Failure To Add New Device To The Database")
		  	        w.WriteHeader(http.StatusUnauthorized) //TODO: Check This Kills The Client Joining
		  	      } else {
		  	        w.WriteHeader(http.StatusOK)
		  	      }
		  			} else {
		  				log.Warning("An Existing Device Has Requested To Enroll. -> 403 (Unauthorized)")
		  				w.WriteHeader(http.StatusUnauthorized)
		  			}
		      }
		    } else if cmd.MessageType == "TokenUpdate" && device != nil {
		      if cmd.update.Token != nil && cmd.update.PushMagic != "string" && cmd.update.UnlockToken != nil && (cmd.update.AwaitingConfiguration == true || cmd.update.AwaitingConfiguration == false) {
		  			if device.Deployed == false {
		  				device.Deployed = true

		  				if status := editDevice(device, true); status == false {
		  	        log.Debug("Failure To Update The Devices Deployment Status")
		  	        w.WriteHeader(http.StatusUnauthorized)
		  					return
		  	      } else {
		  					log.Info("A New Device Joined The MDM: " + device.UDID)
		  	      }
		  			}

		  			if cmd.update.AwaitingConfiguration == true {
		          // TODO: Do DEP Prestage Enrollment By Pushing The Profiles Now Then Run The Fully Setup Thing Then Push The Finished Command
		          log.Error("Unsupported DEP Features")
		  				w.WriteHeader(http.StatusUnauthorized)
		  				return
		        }

		        device.Token = cmd.update.Token
		        device.PushMagic = cmd.update.PushMagic
		        if cmd.update.UnlockToken != nil {
		          device.UnlockToken = cmd.update.UnlockToken
		        }

		  			if status := editDevice(device, true); status == false {
		  				log.Debug("Failure To Update The Devices APNS Tokens")
		  				w.WriteHeader(http.StatusUnauthorized)
		  				return
		  			} else {
		  				log.Debug("Device Updated Its APNS Keys: " + device.UDID)
		  				w.WriteHeader(http.StatusOK)
		  			}
		      } else {
		        log.Warning("A Device Requested To Join With An Invalid Setup (Pre IOS 9 or Doesn't Have Perms)")
		        w.WriteHeader(http.StatusUnauthorized)
		      }
		  	} else if cmd.MessageType == "CheckOut" && device != nil {
		  		err := pgdb.Delete(device)
		  	  if err != nil {
		  	      log.Fatal("Failure To CheckOut Device. The Device Will Not Try Again", err)
		  				w.WriteHeader(http.StatusBadRequest)
		  				return
		  	  }
		  		log.Debug("A Device Has Been Removed From The MDM: " + cmd.UDID)
		  		w.WriteHeader(http.StatusOK)
		    } else {
		  		if device != nil {
		  			log.Warning("Unkown Checkin MessageType of: " + cmd.MessageType)
		  		} else {
		  			log.Warning("A Device Not In The Database Attempted An Action: " + cmd.MessageType)
		  		}
		  		w.WriteHeader(http.StatusBadRequest)
		    }
		*/
		return nil
	}
}

// Parse The Plist Data Coming From The Device
/*var cmd models.CheckinRequest
if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { //TODO: Does It Handle A Blank BOdy?
	log.Println("Error Parsing Checkin Request: ", err)
	http.Error(w, "Error Parsing The Data From The Device", 500) // TODO: Get Correct 4xx Error Code For This
	return nil                                                   // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above
}

log.Println(cmd)*/

/*var cmd models.CheckinCommand
if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil { //TODO: Does It Handle A Blank BOdy?
	log.Println("Error Parsing Checkin Request: ", err)
	http.Error(w, "Error Parsing The Data From The Device", 500) // TODO: Get Correct 4xx Error Code For This
	return nil                                                   // This Does Not Return The Error Due To The MDM Requiring Specific Return Codes With Is Sent Above
}*/

// Retrieve The Device From The Database (The Output Could Be Nil If One Is Not Found)
