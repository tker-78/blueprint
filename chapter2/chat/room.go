package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

// roomのチャンネルに値が送信されたら、ここで受信する
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

// httpのハンドラー
// rへのアクセスでsocketをupgradeして、その情報をもとにclientを生成して、
// チャネル間の通信を行う
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, 1024),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.read()
	client.write()
}
