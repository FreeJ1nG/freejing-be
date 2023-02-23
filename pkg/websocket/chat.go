package websocket

import (
	"database/sql"
	"time"

	"github.com/FreeJ1nG.com/freejing-be/auth"
)

type Chat struct {
	Id         string `json:"id"`
	Sender     string `json:"sender"`
	Message    string `json:"message"`
	CreateDate string `json:"create_date"`
}

func AddChatToHistory(db *sql.DB, sender string, message string) (Chat, error) {
	var chat Chat

	id := auth.GenerateUuid()
	createDate := time.Now().Format(time.RFC3339)

	_, err := db.Exec("INSERT INTO chat_history (id, sender, message, create_date) VALUES ($1, $2, $3, $4)", id, sender, message, createDate)
	if err != nil {
		return chat, err
	}

	row := db.QueryRow("SELECT * FROM chat_history WHERE id = $1", id)
	if err := row.Scan(&chat.Id, &chat.Sender, &chat.Message, &chat.CreateDate); err != nil {
		return chat, err
	}

	return chat, nil
}

func DeleteChatFromHistory(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM chat_history WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetChatHistory(db *sql.DB) ([]Chat, error) {
	var chatHistory []Chat

	rows, err := db.Query("SELECT * FROM chat_history")
	if err != nil {
		return chatHistory, err
	}

	defer rows.Close()
	for rows.Next() {
		var chat Chat
		if err := rows.Scan(&chat.Id, &chat.Sender, &chat.Message, &chat.CreateDate); err != nil {
			return chatHistory, err
		}
		chatHistory = append(chatHistory, chat)
	}

	if err := rows.Err(); err != nil {
		return chatHistory, err
	}

	return chatHistory, nil
}
