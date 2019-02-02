package models

import (
	"github.com/go-ini/ini"
	"log"
)

type Settings struct {
	DB_HOST     string
	DB_TYPE     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     int
}

func LoadSettings(addr string) Settings {
	settings := Settings{}
	cfg, err := ini.Load(addr)
	if err != nil {
		log.Fatal("Failed to read the settings")
	}
	settings.DB_HOST = cfg.Section("").Key("DB_HOST").String()
	settings.DB_TYPE = cfg.Section("").Key("DB_TYPE").String()
	settings.DB_NAME = cfg.Section("").Key("DB_NAME").String()
	settings.DB_USER = cfg.Section("").Key("DB_USER").String()
	settings.DB_PASSWORD = cfg.Section("").Key("DB_PASSWORD").String()
	settings.DB_PORT, err = cfg.Section("").Key("DB_PORT").Int()
	if err != nil {
		log.Fatal("Wrong DB port")
	}
	return settings
}
