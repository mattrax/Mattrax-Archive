package endpoints

import (
	"context"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/groob/plist"
	"github.com/mastahyeti/cms"
	"github.com/mattrax/Mattrax/platform/apple/endpoints/structs"
	"github.com/rs/zerolog/log"
)

// TODO: Clean All Log Output + log.Fatal() or panic()
// TODO: Purge Half Enrolled Devices
// TODO: Structure Into Multiple Files
// TODO: Last Seen Varible On The Device

const ContextUserKey string = "token" // TEMP
func (svc *Service) authenticate(next http.HandlerFunc) http.HandlerFunc {
	scepCertPool := x509.NewCertPool()
	scepCertPool.AddCert(svc.CertStore.ScepCert)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body) // TODO: Does this cause errors for the plist
		if err != nil {
			// TODO: return nil, errors.Wrap(err, "decoding the messages body")
			return
		}
		signature := r.Header.Get("Mdm-Signature")

		if signature == "" {
			// TODO: Handle Not Allowing
			return
		}

		sig, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			// TODO: return nil, errors.Wrap(err, "decode MDM SignMessage header")
			return
		}

		signedCert, err := cms.ParseSignedData(sig)
		if err != nil {
			// TODO: return nil, errors.Wrap(err, "decode MDM SignMessage certificate")
			log.Info().Err(err).Msg("FAIL2")
			return
		}

		// TODO: Checks For If The Certificate Has Been Revoked

		_, err = signedCert.VerifyDetached(body, x509.VerifyOptions{Roots: scepCertPool})
		if err != nil {
			// TODO: return nil, errors.Wrap(err, "something something")
			log.Info().Err(err).Msg("FAIL")
			return
		}

		log.Info().Msg("PASS")

		// TODO: Verify Against The Root Certificate

		/*p7, err := pkcs7.Parse(sig)
		if err != nil {
			// TODO: return nil, errors.Wrap(err, "CMS parse decoded MDM SignMessage signature")
			return
		}

		p7.Content = body
		if err := p7.Verify(); err != nil {
			// TODO: return nil, errors.Wrap(err, "CMS verify MDM Signed Message")
			return
		}
		cert := p7.GetOnlySigner()
		if cert == nil {
			// TODO: return nil, errors.New("invalid or missing CMS signer")
			return
		}*/

		//fmt.Println(sig)

		ctx := context.WithValue(r.Context(), ContextUserKey, "theuser")
		next(w, r.WithContext(ctx))
	}
}

func (svc *Service) checkinHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Context().Value(ContextUserKey))

		//////////////

		// TODO: Make This a middleware
		var cmd structs.CheckinCommand
		if err := plist.NewXMLDecoder(r.Body).Decode(&cmd); err != nil {
			log.Debug().Str("protocol", "apple/checkin").Msg("FAILED Checkin. Invalid Request Body.")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		device, fetchErr := svc.Storage.FindDevice(cmd.UDID)
		if fetchErr == sql.ErrNoRows {
			// TODO: Check For MDM Signature Header
		} else if fetchErr != nil {
			log.Fatal().Err(fetchErr).Msg("Error creating the device") // TODO: Clean Error Message
		}
		/// END MIDDLEWARE HERE

		switch cmd.MessageType {
		case "Authenticate":
			if cmd.Topic == "" || cmd.UDID == "" || cmd.OSVersion == "" || cmd.BuildVersion == "" || cmd.ProductName == "" || cmd.SerialNumber == "" {
				log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Interface("Command", cmd).Msg("FAILED Checkin. Invalid Fields.") //TODO: Dump Fields
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if fetchErr != sql.ErrNoRows && device.Enrolled { // If Device Exists and Completed Enrollment
				log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Msg("FAILED Checkin. The Device Already Exists.")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			//enrolledDevice := NewDeviceFromCheckinAuthentication()

			//device := storage.Device{}
			//device.PopulateFromCheckinCommand(cmd)

			//////////////////////

			// TODO: Handle device.Enrolled = false
			/*
				err := svc.Storage.CreateDevice(storage.Device{
					UUID:           cmd.UDID,
					Enrolled:       false,
					AwaitingConfig: cmd.AwaitingConfiguration,
					Topic:          cmd.Topic,
					OSVersion:      cmd.OSVersion,
					BuildVersion:   cmd.BuildVersion,
					ProductName:    cmd.ProductName,
					SerialNumber:   cmd.SerialNumber,
					IMEI:           cmd.IMEI,
					MEID:           cmd.MEID,
					DeviceName:     cmd.DeviceName,
					Challenge:      cmd.Challenge, // TOOD: Should This Be Put Into The database?
					Model:          cmd.Model,
					ModelName:      cmd.ModelName,
					CreatedBy:      "00000000-0000-0000-0000-000000000000", // TODO: Implemented This
					CreatedAt:      time.Now(),
				})
				log.Info().Msg("A New Device Enrolled") // TODO: Add Helpfull Fields
				if err != nil {
					log.Fatal().Err(err).Msg("Error creating the device") // TODO: Clean Error Message
					w.WriteHeader(http.StatusUnauthorized)
					return
				}*/

		case "TokenUpdate": // TODO: Gracefull Recover From DEP AwaitingConfig - Until I Can get access to DEP to implemented it
			if cmd.Token == nil || cmd.PushMagic == "" {
				log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Interface("Command", cmd).Msg("FAILED Checkin. Invalid Fields.") //TODO: Dump Fields
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if cmd.AwaitingConfiguration {
				log.Error().Msg("Mattrax currently doesn't support DEP")
				log.Error().Msg("If you have access to DEP and a machines serial number contact me about helping build this feature")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if fetchErr == sql.ErrNoRows { // If Device Exists
				log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Msg("FAILED Checkin. Device Isn't Enrolled.")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !device.Enrolled {
				device.Enrolled = true
			}

			fmt.Println(cmd.UserID)
			if cmd.UserID != "" {
				log.Error().Msg("User Based Token Updates Are not Implimented") // TODO;
				return
			}

			// TODO: Handle User TokenUpdate

			device.Token = cmd.Token
			device.PushMagic = cmd.PushMagic
			device.UnlockToken = cmd.UnlockToken

			// TEMP: ioutil.WriteFile("test.token", cmd.Token, 0644)
			// TEMP: fmt.Println(string(cmd.PushMagic))

			err := svc.Storage.UpdateDevice(device)
			if err != nil {
				log.Fatal().Err(err).Msg("Error updating devices tokens") // TODO: Clean Error Message + Dangerous Fatal
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Debug().Msg("A Device Token Updated") // TODO: Log Devices Info
		case "CheckOut":
			if fetchErr == sql.ErrNoRows { // If Device Exists
				log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Msg("FAILED Checkin. Device Isn't Enrolled.")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// TODO: Save Device Info And Alert The Admin - Delete If Admin Confirms

			// err := svc.Storage.DeleteDevice(cmd.UDID)
			// if err != nil {
			// 	log.Fatal().Err(err).Msg("Error updating devices tokens") // TODO: Clean Error Message + Dangerous Fatal
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	return
			// }
			log.Info().Msg("A Device Checked Out") // TODO: Log Devices Info
		default:
			log.Debug().Str("protocol", "apple/checkin").Str("MessageType", cmd.MessageType).Msg("FAILED Checkin. Invalid MessageType.")
			w.WriteHeader(http.StatusUnauthorized)
			log.Info().Msg("An Invalid Checkin MessageType") // TODO: Log The MessageType
			return
		}
	}
}

func (svc *Service) Authenticate(cmd structs.CheckinCommand, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
}
