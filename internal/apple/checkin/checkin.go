package checkin

import "net/http"

// The Web Handler
func Handler() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
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
