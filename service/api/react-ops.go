package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"wasatext/service/customstructs"

	"github.com/julienschmidt/httprouter"
	"github.com/ucarion/emoji"
)

func returnEmptyReaction(w http.ResponseWriter, errorCode int) {
	var emptyReaction customstructs.Reaction
	w.WriteHeader(errorCode)
	jsonReturn, _ := json.Marshal(emptyReaction)
	_, _ = w.Write(jsonReturn)
}

// commentMessage comments a message with a given reaction
func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	auth = strings.TrimPrefix(auth, "Bearer ")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyReaction(w, http.StatusUnauthorized)
		return
	}

	var queryInput struct {
		Message int    `json:"message"`
		Sender  string `json:"sender"`
		Content string `json:"reaction_content"`
	}

	queryBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(queryBody, &queryInput)

	// Check if emoji is valid
	reactionContent, isItValid := emoji.Lookup(queryInput.Content)

	if err != nil || !isItValid {
		returnEmptyReaction(w, http.StatusBadRequest)
		return
	}

	// Check if message exists
	_, errMsg := rt.db.GetMessage(queryInput.Message, true)

	if errMsg != nil {
		returnEmptyReaction(w, http.StatusBadRequest)
		return
	}

	reactions, errReact := rt.db.GetReactions(queryInput.Message)

	// Check if user reacted already
	if errReact == nil {
		// If the user did react, remove the reaction and create a new one
		for _, reaction := range reactions {
			if reaction.Sender.Name == queryInput.Sender {
				err = rt.db.DeleteReaction(reaction.ID)

				if err != nil {
					returnEmptyReaction(w, http.StatusBadRequest)
					return
				}
			}
		}
	}

	// Create the reaction
	reaction, err := rt.db.CreateReaction(queryInput.Message, reactionContent, queryInput.Sender)

	if err != nil {
		returnEmptyReaction(w, http.StatusBadRequest)
		return
	}

	// Reaction created, return it through API
	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(reaction)
	_, _ = w.Write(jsonReturn)
}

// uncommentMessage removes a reaction specified in the path of the URI
func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	auth = strings.TrimPrefix(auth, "Bearer ")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyReaction(w, http.StatusUnauthorized)
		return
	}

	reactionToDelete := ps.ByName("comment_id")
	reactionID, err := strconv.Atoi(reactionToDelete)

	if err != nil {
		returnEmptyReaction(w, http.StatusBadRequest)
		return
	}

	// Remove the reaction
	err = rt.db.DeleteReaction(reactionID)

	if err != nil {
		returnEmptyReaction(w, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
