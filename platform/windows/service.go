package windowsMDM

import (
	"github.com/mattrax/Mattrax/internal/mattrax"
)

// Service is the interface that provides MDM enrollment methods.
type Service interface {
	// Enroll add a new Device under management
	Enroll(device *mattrax.Device) error

	// Update updates a devices details
	Update(device *mattrax.Device) error

	// Unenroll removes a Device from management
	Unenroll(device *mattrax.Device) error
}

type PlatformData struct {
}

type service struct {
	devices      mattrax.DeviceRepository
	policies     mattrax.PolicyRepository
	applications mattrax.ApplicationRepository
	users        mattrax.UserRepository
}

func (s *service) Enroll(device *mattrax.Device) error {
	s.devices.Create(device)

	return nil
}

func (s *service) Update(device *mattrax.Device) error {
	if err := s.devices.Create(device); err != nil {
		return err
	}

	return nil
}

// FUTURE TODO: Separate Removal and MDM Disabling
func (s *service) Unenroll(device *mattrax.Device) error {
	s.devices.Remove(device)

	return nil
}

func NewService(devices mattrax.DeviceRepository, policies mattrax.PolicyRepository, applications mattrax.ApplicationRepository, users mattrax.UserRepository) Service {
	return &service{
		devices,
		policies,
		applications,
		users,
	}
}
