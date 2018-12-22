package mattrax

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// DeviceID is a unique identifier for the device.
type DeviceID string

// Platform defines which MDM protocol is used to communicate with the device
type Platform string

var ApplePlatform = Platform("apple")

var WindowsPlatform = Platform("windows")

// Device is a unique piece of technology under management
type Device struct {
	ID         DeviceID
	Platform   Platform
	PlatformID string // UDID (Apple), DeviceID (Windows)

	PlatformData interface{}

	// Apple: Topic + PushMagic + Token -- Interface
	//      What does "Challenge []byte" Do?

	// Each of these will vary in format based on Platform
	OSVersion    string
	OSEdition    string // BuildVersion (Apple)
	DeviceName   string // TODO: omitempty
	SerialNumber string //TODO: Windows Doesn't Supply Originally

	Policies     []PolicyID
	Applications []ApplicationID

	AssignedTo     UserID
	LatestUpdate   time.Time
	EnrolledBy     UserID
	EnrollmentTime time.Time
}

// GenerateID creates a new unique ID for the device.
func GenerateDeviceID() DeviceID {
	randomID := uuid.NewV4() // TODO: THere Is No Error Handling????

	//TODO: Check For UUID Conflicts In The Repository

	return DeviceID(randomID.String())
}

// ErrInvalidDevice is used when a device has empty values which shouldn't be empty.
var ErrInvalidDeviceValues = errors.New("invalid device values, should not be empty")

// ErrUnknownDevice is used when a device could not be found.
var ErrUnknownDevice = errors.New("unknown device")

//TODO: Ancient Device Error

// DeviceRepository provides access to a device store.
type DeviceRepository interface {
	Create(device *Device) error
	Update(device *Device) error
	Remove(device *Device) error
	Find(id DeviceID) (Device, error)
	FindAll() ([]Device, error)
}
