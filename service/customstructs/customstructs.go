package customstructs

type User struct {
	ID         int32
	Name       string
	Biography  string
	profilePic string
}

type Message struct {
	ID        int32
	Sender    User
	Content   string
	Timestamp string
	Photo     string
	Received  []User
	Read      []User
	Deleted   bool
	Forwarded bool
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
