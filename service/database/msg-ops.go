package database

import (
	"fmt"
	"wasatext/service/customstructs"
)

func (db *appdbimpl) CreateMessage(newMessage customstructs.PrimordialMessage) (customstructs.Message, error) {
	var message customstructs.Message

	_, err := db.c.Exec(
		"INSERT INTO Messages (sender, content, chat, timestamp, photo, forwarded, answer_to, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		newMessage.Sender, newMessage.Content, newMessage.ChatID, newMessage.Timestamp, newMessage.Photo,
		newMessage.Forwarded, newMessage.ReplyingTo, false)

	fmt.Println(err)

	return message, nil
}
