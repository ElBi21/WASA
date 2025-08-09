package customstructs

type User struct {
	Name        string
	DisplayName string
	Biography   string
	ProfilePic  string
}

type Message struct {
	ID        int32
	Sender    User
	Content   string
	ChatID    int32
	Timestamp string
	Photo     string
	Forwarded bool
	AnswerTo  int32

	Received  []User
	Read      []User
	Deleted   bool
	Reactions []Reaction
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
	Messages         []Message
	LastSent         Message
}
