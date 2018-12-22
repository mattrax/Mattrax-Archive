package windowsMDM

import (
	"github.com/mattrax/Mattrax/internal/mattrax"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("mattrax_windows_service")

type loggingService struct {
	next Service
}

func NewLoggingService(s Service) Service {
	return &loggingService{s}
}

func (s *loggingService) Enroll(device *mattrax.Device) error {
	log.Info("") //TODO: I need a structed logger
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
