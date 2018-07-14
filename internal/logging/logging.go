package logging

import (
  "github.com/sirupsen/logrus" // Logging
  "github.com/rifflock/lfshook" // Logging -> Console and File Output With Different Formattings
)

//var Logrus = logrus // TODO: Test This

var log = logrus.New()

//TODO: Go Doc
func init() {
  log.Formatter = &logrus.TextFormatter{
      DisableColors: false,
      TimestampFormat : "02/01/06 15:04:05",
      FullTimestamp:true,
  }
	log.Hooks.Add(lfshook.NewHook(
		"data/log.txt", // FIXME: config.LogFile  //TODO: Append Data/Data+Number For Rolling Log Files Between Restarts
		&logrus.JSONFormatter{},
	))
}

//TODO: Go Doc
func GetLogger() *logrus.Logger { return log }
