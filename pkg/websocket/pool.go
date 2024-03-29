package websocket

import (
	"context"
	"fmt"
	"time"

	"github.com/FreeJ1nG/freejing-be/dbquery"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start(queries *dbquery.Queries, ctx context.Context) {
	for {
		select {
		case originalClient := <-pool.Register:
			pool.Clients[originalClient] = true
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{SenderUsername: client.Username, Type: 1, Body: fmt.Sprintf("%v has joined the chat!", originalClient.Username)})
			}
			chatHistory, err := queries.GetChatHistory(ctx)
			if err != nil {
				fmt.Printf("GetChatHistory: %v\n", err)
				continue
			}
			if len(chatHistory) > 0 {
				fmt.Printf("Chat History: %v\n", chatHistory)
				originalClient.Conn.WriteJSON(chatHistory)
			}
		case originalClient := <-pool.Unregister:
			delete(pool.Clients, originalClient)
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{SenderUsername: client.Username, Type: 1, Body: fmt.Sprintf("%v has disconnected from chat :(", originalClient.Username)})
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			chat, err := queries.CreateChat(ctx, dbquery.CreateChatParams{Sender: message.SenderUsername, Message: message.Body, CreateDate: time.Now()})
			if err != nil {
				fmt.Printf("AddChatToHistory: %v\n", err)
				continue
			}
			fmt.Printf("New Chat: %v\n", chat)
			chatHistory, err := queries.GetChatHistory(ctx)
			if err != nil {
				fmt.Printf("GetChatHistory: %v\n", err)
				continue
			}
			fmt.Printf("Chat History: %v\n", chatHistory)
			for client := range pool.Clients {
				if len(chatHistory) > 0 {
					if err := client.Conn.WriteJSON(chatHistory); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
}
