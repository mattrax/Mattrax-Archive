package logging

import (
	// External Deps
	"github.com/rifflock/lfshook" // Controlling Outputs And Their Formatting
	"github.com/sirupsen/logrus"  // Logging

	// Internal Functions
	mcf "github.com/mattrax/mattrax/internal/configuration" //Mattrax Configuration
)

var config = mcf.GetConfig() // Get The Internal State
var log = logrus.New()       // The Logger

//TODO: Go Doc
func init() {
	log.Formatter = &logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "02/01/06 15:04:05",
		FullTimestamp:   true,
	}
	log.Hooks.Add(lfshook.NewHook(
		config.LogFile, //TODO: Append Data/Data+Number For Rolling Log Files Between Restarts
		&logrus.JSONFormatter{},
	))
}

//TODO: Go Doc
func GetLogger() *logrus.Logger { return log }
