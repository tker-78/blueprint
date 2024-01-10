package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	SecurityKey  string
	ClientId     string
	ClientSecret string
	URL          string
}

var Google ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("error while loading ini.")
	}
	Google = ConfigList{
		SecurityKey:  cfg.Section("google").Key("security_key").String(),
		ClientId:     cfg.Section("google").Key("client_id").String(),
		ClientSecret: cfg.Section("google").Key("client_secret").String(),
		URL:          cfg.Section("google").Key("url").String(),
	}
}
