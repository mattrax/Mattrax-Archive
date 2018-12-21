package windowsMDM

import (
	"log"

	"github.com/mattrax/Mattrax/internal/mattrax"
)

// Service is the interface that provides MDM enrollment methods.
type Service interface {
	// Enroll add a new Device under management
	Enroll(device *mattrax.Device) error

	// Unenroll removes a Device from managment
	Unenroll(device *mattrax.Device) error
}

type service struct {
	devices  mattrax.DeviceRepository
	policies mattrax.PolicyRepository
	users    mattrax.UserRepository
	//TODO: policies, etc
	// APNS Service, etc ----- Research more about this + the package supporting it ---- routingService shipping.RoutingService
}

func (s *service) Enroll(device *mattrax.Device) error {
	log.Println("Enrolling windows device")
	log.Println(device)

	s.devices.Create(device)

	return nil
}

// FUTURE TODO: Separate Removal and MDM Disabling
func (s *service) Unenroll(device *mattrax.Device) error {
	log.Println("Unenrolling windows device")
	log.Println(device)

	s.devices.Remove(device)

	return nil
}

func NewService(devices mattrax.DeviceRepository, policies mattrax.PolicyRepository, users mattrax.UserRepository) Service {
	return &service{
		devices,
		policies,
		users,
	}
}
