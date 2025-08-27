package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// createConversation retrieves all the chats where the user belongs
func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get body of request
	queryBody, _ := io.ReadAll(r.Body)
	var queryInput struct {
		IsPrivate        bool     `json:"isPrivate"`
		Users            []string `json:"users"`
		Name             string   `json:"name"`
		GroupDescription string   `json:"groupDescription"`
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
