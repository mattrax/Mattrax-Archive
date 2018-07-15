package mdm

//go:generate go run generate_marshaler_code.go -out marshaler.go
//go:generate go run generate_payload_code.go -out new_payload.go

import (
	"github.com/satori/go.uuid"
)

// CommandRequest represents an MDM command request
type CommandRequest struct {
	UDID string `json:"udid"`
	Command
}

// Payload is an MDM payload
type Payload struct {
	CommandUUID string
	Command     *Command
}

type Command struct {
	RequestType string `json:"request_type"`
	DeviceInformation
	InstallApplication
	AccountConfiguration
	ScheduleOSUpdateScan
	ScheduleOSUpdate
	InstallProfile
	RemoveProfile
	InstallProvisioningProfile
	RemoveProvisioningProfile
	InstalledApplicationList
	DeviceLock
	ClearPasscode
	EraseDevice
	RequestMirroring
	DeleteUser
	EnableLostMode
	ApplyRedemptionCode
	InstallMedia
	RemoveMedia
	Settings
}

// The following commands are in the order provided by the apple documentation.

// InstallProfile is an InstallProfile MDM Command
type InstallProfile struct {
	Payload []byte `plist:",omitempty" json:"payload,omitempty"`
}

type RemoveProfile struct {
	Identifier string `plist:",omitempty" json:"identifier,omitempty"`
}

type InstallProvisioningProfile struct {
	ProvisioningProfile []byte `plist:",omitempty" json:"provisioning_profile,omitempty"`
}

type RemoveProvisioningProfile struct {
	UUID string `plist:",omitempty" json:"uuid,omitempty"`
}

type InstalledApplicationList struct {
	Identifiers     []string `plist:",omitempty" json:"identifiers,omitempty"`
	ManagedAppsOnly bool     `plist:",omitempty" json:"managed_apps_only,omitempty"`
}

// DeviceInformation is a DeviceInformation MDM Command
type DeviceInformation struct {
	Queries []string `plist:",omitempty" json:"queries,omitempty"`
}

type DeviceLock struct {
	PIN         string `plist:",omitempty" json:"pin,omitempty"`
	Message     string `plist:",omitempty" json:"message,omitempty"`
	PhoneNumber string `plist:",omitempty" json:"phone_number,omitempty"`
}

type ClearPasscode struct {
	UnlockToken []byte `plist:",omitempty" json:"unlock_token,omitempty"`
}

type EraseDevice struct {
	PIN string `plist:",omitempty" json:"pin,omitempty"`
}

type RequestMirroring struct {
	DestinationName     string `plist:",omitempty" json:"destination_name,omitempty"`
	DestinationDeviceID string `plist:",omitempty" json:"destination_device_id,omitempty"`
	ScanTime            string `plist:",omitempty" json:"scan_time,omitempty"`
	Password            string `plist:",omitempty" json:"password,omitempty"`
}

type Restrictions struct {
	ProfileRestrictions bool `plist:",omitempty" json:"restrictions,omitempty"`
}

type DeleteUser struct {
	UserName      string `plist:",omitempty" json:"user_name,omitempty"`
	ForceDeletion bool   `plist:",omitempty" json:"force_deletion,omitempty"`
}

type EnableLostMode struct {
	Message     string `plist:",omitempty" json:"message,omitempty"`
	PhoneNumber string `plist:",omitempty" json:"phone_number,omitempty"`
	Footnote    string `plist:",omitempty" json:"footnote,omitempty"`
}

// InstallApplication is an InstallApplication MDM Command
type InstallApplication struct {
	ITunesStoreID         int                        `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	Identifier            string                     `plist:",omitempty" json:"identifier,omitempty"`
	ManifestURL           string                     `plist:",omitempty" json:"manifest_url,omitempty"`
	ManagementFlags       int                        `plist:",omitempty" json:"management_flags,omitempty"`
	NotManaged            bool                       `plist:",omitempty" json:"not_managed,omitempty"`
	ChangeManagementState string                     `plist:",omitempty" json:"change_management_state,omitempty"`
	Options               *InstallApplicationOptions `plist:",omitempty" json:"options,omitempty"`
	// TODO: add remaining optional fields
}

type InstallApplicationConfiguration struct {
	// TODO: managed app config
}

type InstallApplicationOptions struct {
	NotManaged     bool `plist:",omitempty" json:"not_managed,omitempty"`
	PurchaseMethod int  `plist:",omitempty" json:"purchase_method,omitempty"`
}

type ApplyRedemptionCode struct {
	Identifier     string `plist:",omitempty" json:"identifier,omitempty"`
	RedemptionCode string `plist:",omitempty" json:"redemption_code,omitempty"`
}

type ManagedApplicationList struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type RemoveApplication struct {
	Identifier string `json:"identifier,omitempty"`
}

type InviteToProgram struct {
	ProgramID     string `json:"program_id,omitempty"`
	InvitationURL string `json:"invitation_url,omitempty"`
}

type ValidateApplications struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type InstallMedia struct {
	ITunesStoreID int    `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	MediaURL      string `plist:",omitempty" json:"media_url,omitempty"`
	MediaType     string `json:"media_type"`
	// TODO: media url fields
}

type RemoveMedia struct {
	MediaType     string `json:"media_type"`
	ITunesStoreID int    `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	PersistentID  string `plist:",omitempty" json:"persistent_id,omitempty"`
}

type Settings struct {
	Settings []Setting `plist:",omitempty" json:"settings,omitempty"`
}

type ManagedApplicationConfiguration struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type ApplicationConfiguration struct {
	Identifier    string            `json:"identifier,omitempty"`
	Configuration map[string]string `plist:",omitempty" json:"configuration,omitempty"` // TODO: string map is temporary
}

type ManagedApplicationAttributes struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type ManagedApplicationFeedback struct {
	Identifiers    []string `plist:",omitempty" json:"identifiers,omitempty"`
	DeleteFeedback bool     `plist:",omitempty" json:"delete_feedback,omitempty"`
}

// AccountConfiguration is an MDM command to create a primary user on OS X
// It allows skipping the UI to set up a user.
type AccountConfiguration struct {
	SkipPrimarySetupAccountCreation     bool           `plist:",omitempty" json:"skip_primary_setup_account_creation,omitempty"`
	SetPrimarySetupAccountAsRegularUser bool           `plist:",omitempty" json:"skip_primary_setup_account_as_regular_user,omitempty"`
	AutoSetupAdminAccounts              []AdminAccount `plist:",omitempty" json:"auto_setup_admin_accounts,omitempty"`
}

// AdminAccount is the configuration for the
// Admin account created during Setup Assistant
type AdminAccount struct {
	ShortName    string `plist:"shortName" json:"short_name"`
	FullName     string `plist:"fullName,omitempty" json:"full_name,omitempty"`
	PasswordHash data   `plist:"passwordHash" json:"password_hash"`
	Hidden       bool   `plist:"hidden,omitempty" json:"hidden,omitempty"`
}

type OSUpdate struct {
	ProductKey string `json:"product_key"`
	/*
		One of the following:
		Default: Download and/or install the software update, depending on the current device state. See the UpdateResults dictionary, below, to determine which InstallAction is scheduled.
		DownloadOnly: Download the software update without installing it.
		InstallASAP: Install an already downloaded software update.
		NotifyOnly: Download the software update and notify the user via the App Store (macOS only).
		InstallLater: Download the software update and install it at a later time (macOS only).
	*/
	InstallAction string `json:"install_action"`
}

// ScheduleOSUpdate runs update(s) immediately
type ScheduleOSUpdate struct {
	Updates []OSUpdate `plist:",omitempty" json:"updates,omitempty"`
}

// ScheduleOSUpdateScan schedules a (background) OS SoftwareUpdate check
type ScheduleOSUpdateScan struct {
	Force bool `plist:",omitempty" json:"force,omitempty"`
}

type data []byte

func newPayload(requestType string) *Payload {
	u, _ := uuid.NewV4() //TODO Handle Error
	return &Payload{u.String(),
		&Command{RequestType: requestType}}
}
