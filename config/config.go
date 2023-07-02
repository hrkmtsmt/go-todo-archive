package config

import (
	"log"

	"main/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	config, error := ini.Load("config.ini")

	if error != nil {
		log.Fatalln(error)
	}

	Config = ConfigList{
		Port:      config.Section("web").Key("port").MustString("8080"),
		SQLDriver: config.Section("db").Key("driver").String(),
		DBName:    config.Section("db").Key("name").String(),
		LogFile:   config.Section("web").Key("logfile").String(),
	}
}
