/**
 * Mattrax: An Open Source Device Management System
 * File Description: This File Has All of The Helper Functions For The Structs In The "database.go" File.
 * Package Description: These Are The Structs and Helpers For Device Communication, The API and Database Communication.
 * A HUGE Thanks To MicroMDM. This Package Is A Modied Version Of The (github.com/micromdm/mdm) Package. It Is Used Under The MIT Licence and The Original Work Is Copyright Of MicroMDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package structs

func NewDevice(cmd CheckinCommand) Device {
	return Device{
		UDID:        cmd.UDID,
		DeviceState: 0, //cmd.DeviceState,
		DeviceDetails: DeviceDetails{
			OSVersion:    cmd.Auth.OSVersion,
			BuildVersion: cmd.Auth.BuildVersion,
			ProductName:  cmd.Auth.ProductName,
			SerialNumber: cmd.Auth.SerialNumber,
			IMEI:         cmd.IMEI,
			MEID:         cmd.MEID,
			Profiles: []ProfileListItem{},
		},
		DeviceTokens: DeviceTokens{
			Token:       []byte{},
			PushMagic:   "",
			UnlockToken: []byte{},
		},
		DevicePolicies: DevicePolicies{
			Installed:    []string{},
			LastUpdate: 	0,
			Queued:       []string{},
			Inventory:		DevicePoliciesInventory{
				State: 0,
				CommandUUIDs: map[string]string{},
				LastUpdate: 0,
			},
		},
	}
}
