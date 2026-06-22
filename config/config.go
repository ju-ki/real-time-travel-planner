package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	LogFile string
}

var AppConfig Config

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		os.Exit(1)
	}

	err = cfg.MapTo(&AppConfig)
	if err != nil {
		log.Fatalf("Failed to map config file: %v", err)
		os.Exit(1)
	}

	AppConfig.LogFile = cfg.Section("realTimeTravelPlanner").Key("logFile").String()

}
