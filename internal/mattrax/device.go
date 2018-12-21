package mattrax

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// DeviceID is a unique identifier for the device.
type DeviceID string

// Platform is the MDM protocol that is used to communicate with the device
type Platform string

// Device is a unique piece of technology under management
type Device struct {
	ID         DeviceID
	Platform   Platform
	PlatformID string // UDID (Apple), DeviceID (Windows)

	// Apple: Topic + PushMagic + Token -- Interface
	//      What does "Challenge []byte" Do?

	// Each of these will vary in format based on Platform
	OSVersion    string
	OSEdition    string // BuildVersion (Apple)
	DeviceName   string // TODO: omitempty
	SerialNumber string //TODO: Windows Doesn't Supply Originally

	//Policies []PolicyID

	AssignedTo     UserID
	EnrollmentTime time.Time
	EnrolledBy     UserID
	LatestUpdate   time.Time
}

// GenerateID creates a new unique ID for the device.
func (d *Device) GenerateID() error {
	randomID := uuid.NewV4()

	//TODO: Check For UUID Conflicts In The Repository

	/*if err != nil {
		return err
	}*/

	d.ID = DeviceID(randomID.String())
	return nil
}

// ErrUnknownDevice is used when a device could not be found.
var ErrUnknownDevice = errors.New("unknown device")

// DeviceRepository provides access to a device store.
type DeviceRepository interface {
	Create(device *Device) error
	Remove(device *Device) error
	Find(id DeviceID) (*Device, error)
}
