package websocket

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

var animals = []string{
	"Panda",
	"Tiger",
	"Lemur",
	"Monkey",
	"Fish",
	"Cat",
	"Dog",
	"Dragon",
	"Deer",
	"Dog",
	"Horse",
	"Sheep",
	"Fairy",
	"Pixie",
	"Cow",
	"Ox",
	"Lion",
	"Bat",
	"Bear",
	"Bison",
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func WebsocketHandler(db *sql.DB, pool *Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Websocket Endpoint Hit!")

		conn, err := Upgrade(w, r)
		if err != nil {
			fmt.Fprintf(w, "%+V\n", err)
		}

		client := &Client{Username: "Anonymous " + animals[rand.Intn(len(animals))], Conn: conn, Pool: pool}

		pool.Register <- client
		client.Read()
	}
}
