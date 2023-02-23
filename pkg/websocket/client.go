package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Username string
	Pool     *Pool
}

type Message struct {
	SenderUsername string `json:"sender_username"`
	Type           int    `json:"type"`
	Body           string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{SenderUsername: c.Username, Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %v\n", message)
	}
}
