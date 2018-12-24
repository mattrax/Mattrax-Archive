package windowsMDM

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
		"device": device,
	}).Info("A Windows Device Requested Enrollment")
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
