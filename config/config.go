package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	LogFile   string
	DbName    string
	SQLDriver string
	Port      int
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
	AppConfig.DbName = cfg.Section("db").Key("name").String()
	AppConfig.SQLDriver = cfg.Section("db").Key("driver").String()
	AppConfig.Port = cfg.Section("web").Key("port").MustInt()

}
