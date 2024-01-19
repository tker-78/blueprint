package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	SecurityKey  string
	ClientId     string
	ClientSecret string
	URL          string
}

var Google ConfigList

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error while loading ini.")
	}
	Google = ConfigList{
		SecurityKey:  os.Getenv("security_key"),
		ClientId:     os.Getenv("client_id"),
		ClientSecret: os.Getenv("client_secret"),
		URL:          os.Getenv("url"),
	}
}
