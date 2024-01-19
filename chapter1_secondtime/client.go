package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

// websocketから読み込み
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// websocketに書き込み
// sendに入ってきたもの全てをwebsocketに書き込む
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
