/**
 * Mattrax: An Open Source Device Management System
 * File Description: This File Has A Helper For Generating A New Payload (For The Server Handler).
 * Package Description: These Are The Structs and Helpers For Device Communication, The API and Database Communication.
 * A HUGE Thanks To MicroMDM. This Package Is A Modied Version Of The (github.com/micromdm/mdm) Package. It Is Used Under The MIT Licence and The Original Work Is Copyright Of MicroMDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package structs

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func NewPayload(request *Command) (*Payload, error) {
	requestType := request.RequestType
	var payload *Payload
	if uuid, err := uuid.NewV4(); err != nil {
		return nil, err
	} else {
		payload = &Payload{uuid.String(),
			&Command{RequestType: requestType}}
	}

	switch requestType {

	case "DeviceInformation":
		payload.Command.DeviceInformation = request.DeviceInformation

	case "InstallApplication":
		payload.Command.InstallApplication = request.InstallApplication

	/*case "AccountConfiguration":
		payload.Command.AccountConfiguration = request.AccountConfiguration

	case "ScheduleOSUpdateScan":
		payload.Command.ScheduleOSUpdateScan = request.ScheduleOSUpdateScan

	case "ScheduleOSUpdate":
		payload.Command.ScheduleOSUpdate = request.ScheduleOSUpdate

	case "InstallProfile":
		payload.Command.InstallProfile = request.InstallProfile

	case "RemoveProfile":
		payload.Command.RemoveProfile = request.RemoveProfile

	case "InstallProvisioningProfile":
		payload.Command.InstallProvisioningProfile = request.InstallProvisioningProfile

	case "RemoveProvisioningProfile":
		payload.Command.RemoveProvisioningProfile = request.RemoveProvisioningProfile

	case "InstalledApplicationList":
		payload.Command.InstalledApplicationList = request.InstalledApplicationList

	case "DeviceLock":
		payload.Command.DeviceLock = request.DeviceLock

	case "ClearPasscode":
		payload.Command.ClearPasscode = request.ClearPasscode

	case "EraseDevice":
		payload.Command.EraseDevice = request.EraseDevice

	case "RequestMirroring":
		payload.Command.RequestMirroring = request.RequestMirroring

	case "DeleteUser":
		payload.Command.DeleteUser = request.DeleteUser

	case "EnableLostMode":
		payload.Command.EnableLostMode = request.EnableLostMode

	case "ApplyRedemptionCode":
		payload.Command.ApplyRedemptionCode = request.ApplyRedemptionCode

	case "InstallMedia":
		payload.Command.InstallMedia = request.InstallMedia

	case "RemoveMedia":
		payload.Command.RemoveMedia = request.RemoveMedia

	case "Settings":
		payload.Command.Settings = request.Settings
		*/
	case "ProfileList",
		"ProvisioningProfileList",
		"CertificateList",
		"SecurityInfo",
		"StopMirroring",
		"ClearRestrictionsPassword",
		"UserList",
		"LogOutUser",
		"DisableLostMode",
		"DeviceLocation",
		"ManagedMediaList",
		"OSUpdateStatus",
		"DeviceConfigured",
		"AvailableOSUpdates",
		"Restrictions",
		"ShutDownDevice",
		"RestartDevice":
		return payload, nil

	default:
		return nil, fmt.Errorf("Unsupported MDM RequestType %v", requestType)
	}
	return payload, nil
}
