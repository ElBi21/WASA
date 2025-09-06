package customstructs

import (
	"fmt"
)

func (msg PrimordialMessage) Check() error {
	// Check length of sender ID
	if !(len(msg.Sender) >= 1 && len(msg.Sender) <= 16) {
		return fmt.Errorf("sender ID has unadmissible length")
	}

	// Check length of content
	if !(len(msg.Content) >= 1 && len(msg.Content) <= 65536) {
		return fmt.Errorf("content has unadmissible length")
	}

	// Check length of chatID
	if !(msg.ChatID >= 1 && msg.ChatID <= 65536) {
		return fmt.Errorf("chat has unadmissible ID")
	}

	// Check length of timestamp
	if !(len(msg.Timestamp) <= 44294967296) {
		return fmt.Errorf("timestamp has unadmissible length")
	}

	// Check length of photo
	if !(len(msg.Photo) <= 44294967296) {
		return fmt.Errorf("photo has unadmissible length")
	}

	return nil
}
