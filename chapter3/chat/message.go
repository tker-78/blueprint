package main

import (
	"encoding/json"
	"log"
	"time"
)

type message struct {
	Name    string
	Message string
	When    time.Time
}

func (msg *message) MarshalJSON() ([]byte, error) {
	value, err := json.Marshal(&struct {
		Name    string
		Message string
		When    string
	}{
		Name:    msg.Name,
		Message: msg.Message,
		When:    msg.When.Format("2006年01月02日 15時04分"),
	})
	if err != nil {
		log.Println(err)
		return value, err
	}
	return value, err
}
