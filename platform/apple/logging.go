package appleMDM

import (
	"github.com/mattrax/Mattrax/internal/mattrax"
	log "github.com/sirupsen/logrus"
)

type loggingService struct {
	next Service
}

func NewLoggingService(s Service) Service {
	return &loggingService{s}
}

func (s *loggingService) Enroll(device *mattrax.Device) error {
	log.WithFields(log.Fields{
		"id":           device.ID,
		"platform":     device.Platform,
		"platformID":   device.PlatformID,
		"deviceName":   device.DeviceName,
		"serialnumber": device.SerialNumber,
	}).Info("An Device Requested Enrollment")
	return s.next.Enroll(device)
}

func (s *loggingService) Update(device *mattrax.Device) error {
	log.Info("") //TODO:
	return s.next.Enroll(device)
}

func (s *loggingService) Unenroll(device *mattrax.Device) error {
	log.Info("") //TODO:
	return s.next.Enroll(device)
}
