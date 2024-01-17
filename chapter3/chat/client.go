package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

type message struct {
	Name      string
	Message   string
	When      time.Time
	AvatarURL string
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			// if avatarURL, ok := c.userData["avatar_url"]; ok {
			// 	msg.AvatarURL = avatarURL.(string)
			// }
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			c.room.forward <- msg
		} else {
			break
		}
	}

	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(&msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

func (msg *message) MarshalJSON() ([]byte, error) {
	value, err := json.Marshal(&struct {
		Name      string
		Message   string
		When      string
		AvatarURL string
	}{
		Name:      msg.Name,
		Message:   msg.Message,
		When:      msg.When.Format("2006年01月02日 15時04分"),
		AvatarURL: msg.AvatarURL,
	})
	if err != nil {
		log.Println(err)
		return value, err
	}
	return value, err
}
