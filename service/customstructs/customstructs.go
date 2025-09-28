package customstructs

import "github.com/ucarion/emoji"

type User struct {
	Name        string `json:"user_id"`
	DisplayName string `json:"display_name"`
	Biography   string `json:"biography"`
	ProfilePic  string `json:"profile_pic"`
}

type Message struct {
	ID         int    `json:"message_id"`
	Sender     User   `json:"sender"`
	Content    string `json:"content"`
	ChatID     int    `json:"chat"`
	Timestamp  string `json:"timestamp"`
	Photo      string `json:"photo"`
	Forwarded  bool   `json:"forwarded"`
	ReplyingTo int    `json:"replying"`

	Received  []User     `json:"received"`
	Seen      []User     `json:"seen"`
	Deleted   bool       `json:"deleted"`
	Reactions []Reaction `json:"reactions"`
}

type PrimordialMessage struct {
	ID         int
	Sender     string `json:"sender"`
	Content    string `json:"message_content"`
	ChatID     int    `json:"chat"`
	Timestamp  string `json:"timestamp"`
	Photo      string `json:"photo"`
	Forwarded  bool   `json:"forwarded"`
	ReplyingTo int    `json:"replying"`
}

type Reaction struct {
	ID         int         `json:"reaction_id"`
	RefMessage int         `json:"message"`
	Sender     User        `json:"sender"`
	Content    emoji.Emoji `json:"reaction_content"`
}

type Chat struct {
	ID               int
	IsPrivate        bool
	Users            []User
	Name             string
	GroupDescription string
	Photo            string
	LastSent         Message
}
