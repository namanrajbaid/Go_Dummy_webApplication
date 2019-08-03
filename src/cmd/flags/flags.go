package flags

import (
	"flag"

	"cmd/config"
)

var (
	LogLevel string
	Port     int
	Verbose  bool
)

func init() {
	// initialize different flags in the system
	initLogLevel()
	initPort()
	//initVerbose()
	
	// Parse the flags passed as Command-Line-Args
	flag.Parse()
}

func initLogLevel() {
	flag.StringVar(
		&LogLevel,
		"loglevel",
		config.LogLevel,
		"Supported Log levels : [info, debug, error, warn, fatal, panic]")
}

func initPort() {
	flag.IntVar(
		&Port,
		"port",
		config.Port,
		"Port for Rna-Explorer App")
}