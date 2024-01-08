package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ClientId     string
	ClientSecret string
	Url          string
}

var Google ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("error while loading ini.")
	}
	Google = ConfigList{
		ClientId:     cfg.Section("google").Key("client_id").String(),
		ClientSecret: cfg.Section("google").Key("client_secret").String(),
		Url:          cfg.Section("google").Key("url").String(),
	}
}
