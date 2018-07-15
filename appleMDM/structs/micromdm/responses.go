/**
 * Mattrax: An Open Source Device Management System
 * File Description: This File Has All of The Structs For Parsing The Devices Response From The Server Handler.
 * Package Description: These Are The Structs For The Plist (Device Communication) and JSON (API) Data.
 * A HUGE Thanks To MicroMDM. This Package Is A Modied Version Of The (github.com/micromdm/mdm) Package. It Is Used Under The MIT Licence and The Original Work Is Copyright Of MicroMDM.
 * Protcol Documentation: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html
 * Copyright (C) 2018-2018 Oscar Beaumont <oscartbeaumont@gmail.com>
 */

package mdm

import "time"

// Response is an MDM Command Response
type Response struct {
	UDID                     string
	UserID                   *string `json:"user_id,omitempty" plist:"UserID,omitempty"`
	Status                   string
	CommandUUID              string
	RequestType              string                           `json:"request_type,omitempty" plist:",omitempty"`
	ErrorChain               []ErrorChainItem                 `json:"error_chain" plist:",omitempty"`
	QueryResponses           QueryResponses                   `json:"query_responses,omitempty" plist:",omitempty"`
	SecurityInfo             SecurityInfo                     `json:"security_info,omitempty" plist:",omitempty"`
	CertificateList          CertificateList                  `json:"certificate_list,omitempty" plist:",omitempty"`
	InstalledApplicationList InstalledApplicationListResponse `json:"installed_application_list,omitempty" plist:",omitempty"`
	ScheduleOSUpdateScan     ScheduleOSUpdateScanResponse     `json:"schedule_os_update_scan"`
	OSUpdateStatus           OSUpdateStatusResponse           `json:"os_update_status,omitempty" plist:",omitempty"`
	AvailableOSUpdates       AvailableOSUpdatesResponse       `json:"available_os_updates,omitempty" plist:",omitempty"`
	ProfileList              ProfileList                      `json:"profile_list,omitempty" plist:",omitempty"`
}

type AvailableOSUpdatesResponse []AvailableOSUpdatesResponseItem

type AvailableOSUpdatesResponseItem struct {
	ProductKey        string  `json:"product_key"`
	HumanReadableName string  `json:"human_readable_name"`
	ProductName       string  `json:"product_name"`
	Version           string  `json:"version"`
	Build             string  `json:"build"`
	DownloadSize      int `json:"download_size"`
	InstallSize       float64 `json:"install_size"`

	// Each entry represents an app identifier that is closed to install this update (macOS only).
	AppIdentifiersToClose []string `json:"app_identifiers_to_close"`

	IsCritical                bool `json:"is_critical"`
	IsConfigurationDataUpdate bool `json:"is_configuration_data_update"`
	IsFirmwareUpdate          bool `json:"is_firmware_update"`
	RestartRequired           bool `json:"restart_required"`
	AllowsInstallLater        bool `json:"allows_install_later"`
}

type ScheduleOSUpdateScanResponse struct {
	ScanInitiated bool `json:"scan_initiated"`
}

type OSUpdateStatusResponse []OSUpdateStatusResponseItem

type OSUpdateStatusResponseItem struct {
	ProductKey              string  `json:"product_key"`
	IsDownloaded            bool    `json:"is_downloaded"`
	DownloadPercentComplete float64 `json:"download_percent_complete"`
	/*
		The status of this update. Possible values are:
		Idle: No action is being taken on this software update.
		Downloading: The software update is being downloaded.
		Installing: The software update is being installed. This status may not be returned if the device must reboot during installation.
	*/
	Status string `json:"status"`
}

type ProvisioningProfileListItem struct {
	Name       string    `plist:",omitempty" json:"name,omitempty"`
	UUID       string    `plist:",omitempty" json:"uuid,omitempty"`
	ExpiryDate time.Time `plist:",omitempty" json:"expiry_date,omitempty"`
}

type ProvisioningProfileListResponse []ProvisioningProfileListItem

type CertificateListItem struct {
	CommonName string `json:"common_name"`
	Data       []byte `json:"data"`
	IsIdentity bool   `json:"is_identity"`
}

// CertificateList is the CertificateList MDM Command Response
type CertificateList []CertificateListItem

type InstalledApplicationListItem struct {
	Identifier   string `plist:",omitempty" json:"identifier,omitempty"`
	Version      string `plist:",omitempty" json:"version,omitempty"`
	ShortVersion string `plist:",omitempty" json:"short_version,omitempty"`
	Name         string `json:"name,omitempty"`
	BundleSize   uint32 `plist:",omitempty" json:"bundle_size,omitempty"`
	DynamicSize  uint32 `plist:",omitempty" json:"dynamic_size,omitempty"`
	IsValidated  bool   `plist:",omitempty" json:"is_validated,omitempty"`
}

type InstalledApplicationListResponse []InstalledApplicationListItem

// CommonQueryResponses has a list of query responses common to all device types
type CommonQueryResponses struct {
	UDID                  string            `json:"udid"`
	Languages             []string          `json:"languages,omitempty"`              // ATV 6+
	Locales               []string          `json:"locales,omitempty"`                // ATV 6+
	DeviceID              string            `json:"device_id"`                        // ATV 6+
	OrganizationInfo      map[string]string `json:"organization_info,omitempty"`      // IOS 7+
	LastCloudBackupDate   time.Time         `json:"last_cloud_backup_date,omitempty"` // IOS 8+
	AwaitingConfiguration bool              `json:"awaiting_configuration"`           // IOS 9+
	// iTunes
	ITunesStoreAccountIsActive bool   `json:"itunes_store_account_is_active"` // IOS 7+ OSX 10.9+
	ITunesStoreAccountHash     string `json:"itunes_store_account_hash"`      // IOS 8+ OSX 10.10+

	// Device
	DeviceName                    string  `json:"device_name"`
	OSVersion                     string  `json:"os_version"`
	BuildVersion                  string  `json:"build_version"`
	ModelName                     string  `json:"model_name"`
	Model                         string  `json:"model"`
	ProductName                   string  `json:"product_name"`
	SerialNumber                  string  `json:"serial_number"`
	DeviceCapacity                float32 `json:"device_capacity"`
	AvailableDeviceCapacity       float32 `json:"available_device_capacity"`
	BatteryLevel                  float32 `json:"battery_level,omitempty"`           // IOS 5+
	CellularTechnology            int     `json:"cellular_technology,omitempty"`     // IOS 4+
	IsSupervised                  bool    `json:"is_supervised"`                     // IOS 6+
	IsDeviceLocatorServiceEnabled bool    `json:"is_device_locator_service_enabled"` // IOS 7+
	IsActivationLockEnabled       bool    `json:"is_activation_lock_enabled"`        // IOS 7+ OSX 10.9+
	IsDoNotDisturbInEffect        bool    `json:"is_dnd_in_effect"`                  // IOS 7+
	EASDeviceIdentifier           string  `json:"eas_device_identifier"`             // IOS 7 OSX 10.9
	IsCloudBackupEnabled          bool    `json:"is_cloud_backup_enabled"`           // IOS 7.1

	// Network
	BluetoothMAC string   `json:"bluetooth_mac"`
	WiFiMAC      string   `json:"wifi_mac"`
	EthernetMACs []string `json:"ethernet_macs"` // Surprisingly works in IOS
}

// AtvQueryResponses contains AppleTV QueryResponses
type AtvQueryResponses struct {
}

// IosQueryResponses contains iOS QueryResponses
type IosQueryResponses struct {
	IMEI                 string `json:"imei"`
	MEID                 string `json:"meid"`
	ModemFirmwareVersion string `json:"modem_firmware_version"`
	IsMDMLostModeEnabled bool   `json:"is_mdm_lost_mode_enabled,omitempty"` // IOS 9.3
	MaximumResidentUsers int    `json:"maximum_resident_users"`             // IOS 9.3

	// Network
	ICCID                    string `json:"iccid,omitempty"` // IOS
	CurrentCarrierNetwork    string `json:"current_carrier_network,omitempty"`
	SIMCarrierNetwork        string `json:"sim_carrier_network,omitempty"`
	SubscriberCarrierNetwork string `json:"subscriber_carrier_network,omitempty"`
	CarrierSettingsVersion   string `json:"carrier_settings_version,omitempty"`
	PhoneNumber              string `json:"phone_number,omitempty"`
	VoiceRoamingEnabled      bool   `json:"voice_roaming_enabled,omitempty"`
	DataRoamingEnabled       bool   `json:"data_roaming_enabled,omitempty"`
	IsRoaming                bool   `json:"is_roaming,omitempty"`
	PersonalHotspotEnabled   bool   `json:"personal_hotspot_enabled,omitempty"`
	SubscriberMCC            string `json:"subscriber_mcc,omitempty"`
	SubscriberMNC            string `json:"subscriber_mnc,omitempty"`
	CurrentMCC               string `json:"current_mcc,omitempty"`
	CurrentMNC               string `json:"current_mnc,omitempty"`
}

// OSUpdateSettingsResponse contains information about macOS update settings.
type OSUpdateSettingsResponse struct {
	AutoCheckEnabled                bool      `json:"auto_check_enabled"`
	AutomaticAppInstallationEnabled bool      `json:"automatic_app_installation_enabled"`
	AutomaticOSInstallationEnabled  bool      `json:"automatic_os_installation_enabled"`
	AutomaticSecurityUpdatesEnabled bool      `json:"automatic_security_updates_enabled"`
	BackgroundDownloadEnabled       bool      `json:"background_download_enabled"`
	CatalogURL                      string    `json:"catalog_url"`
	IsDefaultCatalog                bool      `json:"is_default_catalog"`
	PerformPeriodicCheck            bool      `json:"perform_periodic_check"`
	PreviousScanDate                time.Time `json:"previous_scan_date"`
	PreviousScanResult              int       `json:"previous_scan_result"`
}

// MacosQueryResponses contains macOS queryResponses
type MacosQueryResponses struct {
	OSUpdateSettings   OSUpdateSettingsResponse // OSX 10.11+
	LocalHostName      string                   `json:"local_host_name,omitempty"` // OSX 10.11
	HostName           string                   `json:"host_name,omitempty"`       // OSX 10.11
	ActiveManagedUsers []string                 `json:"active_managed_users"`      // OSX 10.11
}

// QueryResponses is a DeviceInformation MDM Command Response
type QueryResponses struct {
	CommonQueryResponses
	MacosQueryResponses
	IosQueryResponses
	AtvQueryResponses
}

// SecurityInfo is the SecurityInfo MDM Command Response
type SecurityInfo struct {
	FDEEnabled                     bool `json:"fde_enabled,omitempty"` // OSX
	FDEHasPersonalRecoveryKey      bool `json:"fde_has_personal_recovery_key,omitempty"`
	FDEHasInstitutionalRecoveryKey bool `json:"fde_has_institutional_recovery_key,omitempty"`

	HardwareEncryptionCaps        int  `json:"hardware_encryption_caps,omitempty"` // iOS
	PasscodeCompliant             bool `json:"passcode_compliant,omitempty"`
	PasscodeCompliantWithProfiles bool `json:"passcode_compliant_with_profiles,omitempty"`
	PasscodePresent               bool `json:"passcode_present,omitempty"`
}

type RequestMirroringResponse struct {
	MirroringResult string `json:"mirroring_result,omitempty"`
}

//type GlobalRestrictions struct {
//	RestrictedBool map[string]bool `plist:"restrictedBool,omitempty" json:"restricted_bool,omitempty"`
//	RestrictedValue map[string]int `plist:"restrictedValue,omitempty" json:"restricted_value,omitempty"`
//	Intersection map[string]string `plist:"intersection,omitempty" json:"intersection,omitempty"` // TODO: not actually string values
//	Union map[string]string `plist:"union,omitempty" json:"union,omitempty"` // TODO: not actually string values
//}

type UsersListItem struct {
	UserName      string `json:"user_name,omitempty"`
	HasDataToSync bool   `json:"has_data_to_sync,omitempty"`
	DataQuota     int    `json:"data_quota,omitempty"`
	DataUsed      int    `json:"data_used,omitempty"`
}

type UsersListResponse []UsersListItem

// Represents a single error in the error chain response
type ErrorChainItem struct {
	ErrorCode            int    `json:"error_code,omitempty"`
	ErrorDomain          string `json:"error_domain,omitempty"`
	LocalizedDescription string `json:"localized_description,omitempty"`
	USEnglishDescription string `json:"us_english_description,omitempty"`
}

type ErrorChain []ErrorChainItem
type PayloadContentItem struct {
	PayloadDescription       string `json:"payload_description"`
	PayloadDisplayName       string `json:"payload_display_name"`
	PayloadIdentifier        string `json:"payload_identifier"`
	PayloadOrganization      string `json:"payload_organization"`
	PayloadVersion           int    `json:"payload_version"`
}

type ProfileListItem struct {
	HasRemovalPasscode       bool                 `json:"has_removal_passcode"`
	IsEncrypted              bool                 `json:"is_encrypted"`
	PayloadContent           []PayloadContentItem `json:"payload_content,omitempty" plist:",omitempty"`
	PayloadDescription       string               `json:"payload_description"`
	PayloadDisplayName       string               `json:"payload_display_name"`
	PayloadIdentifier        string               `json:"payload_identifier"`
	PayloadOrganization      string               `json:"payload_organization"`
	PayloadRemovalDisallowed bool                 `json:"payload_removal_disallowed"`
	PayloadUUID              string               `json:"payload_uuid"`
	PayloadVersion           int                  `json:"payload_version"`
}
type ProfileList []ProfileListItem
