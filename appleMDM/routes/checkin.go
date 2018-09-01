package routes

import (
	//TODO WTF Does This Import Statement Need To Look Like This
	"net/http"
	"time"

	log "github.com/labstack/gommon/log" //TODO: WTF Again

	"../../internal/pgsql" //TODO: Full Path

	devices "../../internal/devices" //TODO: WTF Does This Need The Start To Import With Go-plus
	"../structs"                     //TODO: Full Path
	"github.com/groob/plist"
	"github.com/labstack/echo"
)

var pgdb = pgsql.GetDB()

func CheckinHandler(c echo.Context) error { //TODO: Is Returning 500 On Error Fine -> Cause It Is Not In THe Spec
	// Parse The Plist From The Client
	var cmd structs.CheckinCommand
	if err := plist.NewXMLDecoder(c.Request().Body).Decode(&cmd); err != nil {
		return err
	}

	// Verify The Request Is Valid
	if cmd.UDID == "" { //TODO: Maybe Verify "cmd.Topic" Here
		return c.String(http.StatusUnauthorized, "Incorrect Payload Parameters")
	}

	// Attempt To Get The Device From the Database
	var device devices.Computer
	if err := pgdb.Model(&device).Where("uuid = ?", cmd.UDID).Select(); err != nil && pgsql.NotFound(err) {
		return err
	}

	if cmd.MessageType == "Authenticate" {
		return authenticate(device, cmd, c)
	} else if cmd.MessageType == "TokenUpdate" {
		return tokenUpdate(device, cmd, c)
	} else if cmd.MessageType == "CheckOut" {
		return checkOut(device, cmd, c)
	}
	return c.String(http.StatusUnauthorized, "")
}

func authenticate(device devices.Computer, cmd structs.CheckinCommand, c echo.Context) error {
	if deviceState, ok := device.DeviceState.(structs.MacOS_DeviceState); ok { //TODO: Not Working Right
		log.Info(deviceState.EnrollmentState)
		//&& deviceState.EnrollmentState == 0
		return c.String(http.StatusUnauthorized, "The Device Is Partly Enrolled")
	}

	// Check That The Device Doesn't Already Exist
	if device.UUID != "" {
		return c.String(http.StatusUnauthorized, "Device Already Known To The MDM")
	}

	// Check The Information Sent By The Client
	if cmd.Auth.OSVersion == "" && cmd.Auth.BuildVersion == "" && cmd.Auth.ProductName == "" && cmd.Auth.SerialNumber == "" && cmd.Auth.IMEI == "" && cmd.Auth.MEID == "" {
		return c.String(http.StatusUnauthorized, "Incorrect Payload Parameters")
	}

	//TODO: Handle The "Topic" Sent From The Client
	//TODO: Handle The Device States And Ones That Are Not Fully Enrolled

	// Create The New Device
	enrollingDevice := devices.Computer{
		UUID: cmd.UDID,
		DeviceState: structs.MacOS_DeviceState{
			Token:           []byte{},
			PushMagic:       "",
			UnlockToken:     []byte{},
			LastUpdate:      time.Now().Unix(),
			EnrollmentState: 0,
		},
		DeviceInfo: structs.MacOS_DeviceInfo{
			OSVersion:    cmd.Auth.OSVersion,
			BuildVersion: cmd.Auth.BuildVersion,
			ProductName:  cmd.Auth.ProductName,
			SerialNumber: cmd.Auth.SerialNumber,
			IMEI:         cmd.IMEI,
			MEID:         cmd.MEID,
		},
	}

	// Add The Device To The Database
	if err := pgdb.Insert(&enrollingDevice); err != nil {
		return err
	}

	// Log The Event & Return Success To The Client
	log.Info(cmd.SerialNumber + "  - Enrolled With The MDM!")
	return c.String(200, "")
}

func tokenUpdate(device devices.Computer, cmd structs.CheckinCommand, c echo.Context) error {
	return c.String(200, "")
}

func checkOut(device devices.Computer, cmd structs.CheckinCommand, c echo.Context) error {
	log.Info(cmd.SerialNumber + "  - Removed From The MDM!")
	return c.String(200, "")
}
