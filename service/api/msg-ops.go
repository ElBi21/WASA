package api

import (
	"encoding/json"
	"io"
	"net/http"
	"wasatext/service/customstructs"

	"github.com/julienschmidt/httprouter"
)

// sendMessage sends a message in the chat specified
func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get body of request
	queryBody, _ := io.ReadAll(r.Body)
	var newMsg customstructs.PrimordialMessage

	err := json.Unmarshal(queryBody, &newMsg)

	// If there have been errors (so missing fields or inexistent chat number) then return 418
	if err != nil || newMsg.Check() != nil {
		w.WriteHeader(http.StatusTeapot)
	}

	_, err = rt.db.CreateMessage(newMsg)

	if err != nil {
		w.WriteHeader(http.StatusTeapot)
	}

}
