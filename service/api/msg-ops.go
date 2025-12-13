package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wasatext/service/customstructs"

	"github.com/julienschmidt/httprouter"
)

// returnEmptyMessage returns an empty message through the API
func returnEmptyMessage(w http.ResponseWriter, errorCode int) {
	var emptyReaction customstructs.Message
	w.WriteHeader(errorCode)
	jsonReturn, _ := json.Marshal(emptyReaction)
	_, _ = w.Write(jsonReturn)
}

// sendMessage sends a message in the chat specified
func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	auth = strings.TrimPrefix(auth, "Bearer ")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyMessage(w, http.StatusUnauthorized)
		return
	}

	// Get body of request
	queryBody, _ := io.ReadAll(r.Body)
	var newMsg customstructs.PrimordialMessage
	newMsg.Timestamp = time.Now().Format("2006-01-02 15:04:05.000")
	newMsg.Forwarded = false

	err := json.Unmarshal(queryBody, &newMsg)

	// If there have been errors (so missing fields or inexistent chat number) then return 418
	if err != nil || newMsg.Check() != nil {
		returnEmptyMessage(w, http.StatusTeapot)
		return
	}

	message, err := rt.db.CreateMessage(newMsg)

	// Return error code if failed to create message on the database
	if err != nil {
		returnEmptyMessage(w, http.StatusTeapot)
		return
	}

	chatObj, _ := rt.db.GetConversation(message.ChatID)

	// Place the message in the tables of received and read from the sender
	for _, groupUser := range chatObj.Users {
		_ = rt.db.AddReceivedMessage(message.ID, groupUser.Name)
		message.Received = append(message.Received, groupUser)
	}

	// _ = rt.db.AddSeenMessage(message.ID, message.Sender.Name)
	// message.Seen = append(message.Seen, message.Sender)

	// Return the message through the API
	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(message)
	_, _ = w.Write(jsonReturn)
}

// forwardMessage forwards a message to another chat
func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	auth = strings.TrimPrefix(auth, "Bearer ")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyMessage(w, http.StatusUnauthorized)
		return
	}

	var queryInput struct {
		Sender string `json:"sender"`
		ChatID int    `json:"chat_id"`
	}

	queryBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(queryBody, &queryInput)

	if err != nil {
		returnEmptyMessage(w, http.StatusForbidden)
		return
	}

	// Retrieve ID of message to forward
	messageIDParam := ps.ByName("message_id")
	messageID, err := strconv.Atoi(messageIDParam)

	if err != nil {
		returnEmptyMessage(w, http.StatusForbidden)
		return
	}

	// Get the message and the sender
	message, errMessage := rt.db.GetMessage(messageID, true)
	sender, errSender := rt.db.GetUserByName(queryInput.Sender)

	// If message doesn't exist, return error
	if errMessage != nil || errSender != nil {
		returnEmptyMessage(w, http.StatusNotFound)
		return
	}

	// If the user is trying to forward a deleted message, abort
	if message.Deleted {
		returnEmptyMessage(w, http.StatusNotFound)
		return
	}

	// Prepare message and send it to DB
	message.Timestamp = time.Now().Format("2006-01-02 15:04:05.000")
	message.Sender = sender
	message.ChatID = queryInput.ChatID
	message.Forwarded = true

	newID, err := rt.db.ForwardMessage(message)

	if err != nil {
		returnEmptyMessage(w, http.StatusNotFound)
		return
	}

	// Update the ID
	message.ID = newID

	// Add message to received and seen tables
	_ = rt.db.AddReceivedMessage(message.ID, message.Sender.Name)
	message.Received = append(message.Received, message.Sender)
	_ = rt.db.AddSeenMessage(message.ID, message.Sender.Name)
	message.Seen = append(message.Seen, message.Sender)

	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(message)
	_, _ = w.Write(jsonReturn)
}

// deleteMessage deletes a message by flagging it in the DB as "removed". The message's ID
// is specified in the URI path
func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	auth = strings.TrimPrefix(auth, "Bearer ")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyMessage(w, http.StatusUnauthorized)
		return
	}

	param := ps.ByName("message_id")
	messageID, err := strconv.Atoi(param)

	if err != nil {
		returnEmptyMessage(w, http.StatusBadRequest)
		return
	}

	// Flag the message as deleted
	message, err := rt.db.DeleteMessage(messageID)

	if err != nil {
		returnEmptyMessage(w, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(message)
	_, _ = w.Write(jsonReturn)
}
