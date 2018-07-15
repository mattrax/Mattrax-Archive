package logging

import (
	// External Deps
	"github.com/rifflock/lfshook" // Controlling Outputs And Their Formatting
	logrus "github.com/sirupsen/logrus" // Logging

	// Internal Functions
	mcf "github.com/mattrax/Mattrax/internal/configuration" //Mattrax Configuration
)

var (
	config = mcf.GetConfig() // Get The Internal State
 	log = logrus.New()       // The Logger
)
type Fields = logrus.Fields  // Export Logrus Fields (So It Does Have To Be Imported By Another Package)

//TODO: Go Doc
func init() {
	if config.Verbose { log.Level = logrus.DebugLevel }
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
func GetLogger() *logrus.Logger { return log  }
