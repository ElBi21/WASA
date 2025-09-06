package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createConversation retrieves all the chats where the user belongs
func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get body of request
	queryBody, _ := io.ReadAll(r.Body)
	var queryInput struct {
		IsPrivate        bool     `json:"is_private"`
		Users            []string `json:"users"`
		Name             string   `json:"name"`
		GroupDescription string   `json:"group_description"`
		Photo            string   `json:"photo"`
	}

	_ = json.Unmarshal(queryBody, &queryInput)

	// Execute query on DB
	_, err := rt.db.CreateConversation(queryInput.IsPrivate, queryInput.Users, queryInput.Name, queryInput.GroupDescription, queryInput.Photo)

	if err != nil {
		w.WriteHeader(http.StatusRequestTimeout)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	jsonReturn, _ := json.Marshal(queryInput)
	_, _ = w.Write(jsonReturn)
}

// getConversation retrieves the data for a conversation, together with the last sent message.
// TODO: Implement the Seen and Received flags, to do once I reach the /message stuff
func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	chatIdParam := ps.ByName("chat_id")

	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Retrieve the chat
	chat, err := rt.db.GetConversation(chatId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// Return the chat
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}
