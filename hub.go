package main

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.这里是广播站
type Hub struct {
	// Registered clients. 注册的客户端
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients. 注册来着client的请求
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register://注册client
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				//注销client并关闭其send通道
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast://如果有要广播的消息
			for client := range h.clients {
				//遍历已注册的客户端 船体消息
				fmt.Println(client)
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
