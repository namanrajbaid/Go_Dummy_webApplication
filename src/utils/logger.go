package utils

import (
	"cmd/flags"
	log "github.com/sirupsen/logrus"
)

func InitializeLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	level, err := log.ParseLevel(flags.LogLevel)
	if err == nil {
		log.SetLevel(level)
	}
}