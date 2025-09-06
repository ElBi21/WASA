package customstructs

type User struct {
	Name        string `json:"user_id"`
	DisplayName string `json:"display_name"`
	Biography   string `json:"biography"`
	ProfilePic  string `json:"profile_pic"`
}

type Message struct {
	ID         int32
	Sender     User   `json:"sender"`
	Content    string `json:"content"`
	ChatID     int32  `json:"chat"`
	Timestamp  string `json:"timestamp"`
	Photo      string `json:"photo"`
	Forwarded  bool   `json:"forwarded"`
	ReplyingTo int32  `json:"replying"`

	Received  []User
	Read      []User
	Deleted   bool
	Reactions []Reaction
}

type PrimordialMessage struct {
	ID         int32
	Sender     string `json:"sender"`
	Content    string `json:"content"`
	ChatID     int32  `json:"chat"`
	Timestamp  string `json:"timestamp"`
	Photo      string `json:"photo"`
	Forwarded  bool   `json:"forwarded"`
	ReplyingTo int32  `json:"replying"`
}

type Reaction struct {
	ID         int32
	RefMessage Message
	Sender     User
	Content    string
}

type Chat struct {
	ID               int32
	IsPrivate        bool
	Users            []User
	Name             string
	GroupDescription string
	Photo            string
	LastSent         Message
}
