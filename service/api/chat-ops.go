package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"wasatext/service/customstructs"

	"github.com/julienschmidt/httprouter"
)

// returnEmptyUser creates an empty user and puts it into the response
func returnEmptyUser(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	var user customstructs.User
	jsonReturn, _ := json.Marshal(user)
	_, _ = w.Write(jsonReturn)
}

// returnEmptyChat creates an empty chat and puts it into the response
func returnEmptyChat(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	var chat customstructs.Chat
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}

// createConversation retrieves all the chats where the user belongs
func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

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
	var err error

	// Execute query on DB
	if queryInput.IsPrivate {
		// If the private chat has not defined both ends, then return error (or if more than 2 users are added)
		if len(queryInput.Users) != 2 {
			returnEmptyChat(w, http.StatusBadRequest)
			return
		}

		queryInput.Name = ""
		queryInput.GroupDescription = ""
		queryInput.Photo = ""
	}

	_, err = rt.db.CreateConversation(
		queryInput.IsPrivate, queryInput.Users, queryInput.Name, queryInput.GroupDescription, queryInput.Photo)

	if err != nil {
		returnEmptyChat(w, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsonReturn, _ := json.Marshal(queryInput)
	_, _ = w.Write(jsonReturn)
}

// getConversation retrieves the data for a conversation, together with the last sent message.
func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

	chatIdParam := ps.ByName("chat_id")

	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		returnEmptyChat(w, http.StatusBadRequest)
		return
	}

	// Retrieve the chat
	chat, err := rt.db.GetConversation(chatId)

	if err != nil {
		returnEmptyChat(w, http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// Return the chat
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}

// addToGroup adds a user to a group, specified in the URI of the request
func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

	chatParam := ps.ByName("chat_id")
	chatID, errParam := strconv.Atoi(chatParam)

	queryBody, _ := io.ReadAll(r.Body)
	var queryInput struct {
		NewUser string `json:"new_user"`
	}

	errInput := json.Unmarshal(queryBody, &queryInput)

	if errParam != nil || errInput != nil {
		returnEmptyUser(w, http.StatusNotFound)
		return
	}

	// Check if the user and the chat do exist
	user, errUser := rt.db.GetUserByName(queryInput.NewUser)
	chat, errChat := rt.db.GetConversation(chatID)

	if errUser != nil || errChat != nil {
		returnEmptyUser(w, http.StatusNotFound)
		return
	}

	if len(chat.Users) != 0 {
		// Check if user is already in the chat
		for _, item := range chat.Users {
			if user.Name == item.Name {
				returnEmptyUser(w, http.StatusBadRequest)
				return
			}
		}
	}

	// If arrived here, it means that the user is not in the chat already, hence, we can add it
	err := rt.db.AddUserToGroup(chatID, user.Name)

	if err != nil {
		returnEmptyUser(w, http.StatusNotFound)
		return
	}

	// Add the user to the return object
	chat.Users = append(chat.Users, user)

	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}

// setGroupDescription sets a new description (max length of 4098) for a group, which must be a non-private chat
func (rt *_router) setGroupDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

	chatParam := ps.ByName("chat_id")
	chatID, errParam := strconv.Atoi(chatParam)

	queryBody, _ := io.ReadAll(r.Body)
	var queryInput struct {
		NewDescription string `json:"new_description"`
	}

	errInput := json.Unmarshal(queryBody, &queryInput)

	if errParam != nil || errInput != nil || len(queryInput.NewDescription) > 4098 {
		returnEmptyChat(w, http.StatusBadRequest)
		return
	}

	// Check if the chat is private
	chat, errChat := rt.db.GetConversation(chatID)

	if errChat != nil || chat.IsPrivate {
		returnEmptyChat(w, http.StatusBadRequest)
		return
	}

	// Change description
	err := rt.db.SetNewGroupDescription(chatID, queryInput.NewDescription)

	if err != nil {
		returnEmptyChat(w, http.StatusBadRequest)
		return
	}

	chat.GroupDescription = queryInput.NewDescription
	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}

// leaveGroup removes a user from a group
func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

	chatParam := ps.ByName("chat_id")
	leaveUserID := ps.ByName("user_id")

	user, errUser := rt.db.GetUserByName(leaveUserID)
	chatID, errParamChat := strconv.Atoi(chatParam)

	if errParamChat != nil || errUser != nil {
		returnEmptyUser(w, http.StatusBadRequest)
		return
	}

	// Check if the chat exists and if the user is in the chat
	chat, errChat := rt.db.GetConversation(chatID)

	if errChat != nil {
		returnEmptyUser(w, http.StatusBadRequest)
		return
	}

	userIsInChat := false
	for _, user := range chat.Users {
		if user.Name == leaveUserID {
			userIsInChat = true
		}
	}

	if !userIsInChat {
		returnEmptyUser(w, http.StatusBadRequest)
		return
	}

	// 2 paths: if the user is in a group, then just remove it. If the chat is private, delete the chat
	if chat.IsPrivate {
		for _, user := range chat.Users {
			// Remove users from ChatsUsers table
			err := rt.db.RemoveUserFromGroup(chatID, user.Name)

			if err != nil {
				returnEmptyUser(w, http.StatusBadRequest)
				return
			}

			// Delete private chat between them
			errChatDelete, errMsgDelete := rt.db.DeleteChatAndMessages(chat)

			if errChatDelete != nil || errMsgDelete != nil {
				returnEmptyUser(w, http.StatusBadRequest)
				return
			}
		}
	} else {
		err := rt.db.RemoveUserFromGroup(chatID, leaveUserID)

		if err != nil {
			returnEmptyUser(w, http.StatusBadRequest)
			return
		}

		// If the user is the last one in the group, delete the group and the messages
		if len(chat.Users) == 1 && chat.Users[0].Name == leaveUserID {
			errChatDelete, errMsgDelete := rt.db.DeleteChatAndMessages(chat)

			if errChatDelete != nil || errMsgDelete != nil {
				returnEmptyUser(w, http.StatusBadRequest)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(user)
	_, _ = w.Write(jsonReturn)
}

// setGroupName sets a new group name for a group
func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

	chatParam := ps.ByName("chat_id")
	chatID, errParam := strconv.Atoi(chatParam)

	queryBody, _ := io.ReadAll(r.Body)
	var queryInput struct {
		NewGroupName string `json:"new_name"`
	}

	errInput := json.Unmarshal(queryBody, &queryInput)

	if errParam != nil || errInput != nil {
		returnEmptyChat(w, http.StatusNotFound)
		return
	}

	// Retrieve the chat from the DB
	chat, errChat := rt.db.GetConversation(chatID)

	// If the chat does not exist, or it's private, we can't continue, thus we abort
	if errChat != nil || chat.IsPrivate {
		returnEmptyChat(w, http.StatusNotFound)
		return
	}

	// Change the name of the group
	err := rt.db.SetNewGroupName(chatID, queryInput.NewGroupName)

	if err != nil {
		returnEmptyChat(w, http.StatusNotFound)
		return
	}

	chat.Name = queryInput.NewGroupName

	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}

// setGroupName sets a new group picture for a group
func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := r.Header.Get("Authorization")
	_, errAuth := rt.db.GetUserByName(auth)
	if errAuth != nil {
		returnEmptyChat(w, http.StatusUnauthorized)
		return
	}

	chatParam := ps.ByName("chat_id")
	chatID, errParam := strconv.Atoi(chatParam)

	queryBody, _ := io.ReadAll(r.Body)
	var queryInput struct {
		NewPhoto string `json:"new_photo"`
	}

	errInput := json.Unmarshal(queryBody, &queryInput)

	if errParam != nil || errInput != nil {
		returnEmptyChat(w, http.StatusNotFound)
		return
	}

	// Retrieve the chat from the DB
	chat, errChat := rt.db.GetConversation(chatID)

	// If the chat does not exist, or it's private, we can't continue, thus we abort
	if errChat != nil || chat.IsPrivate {
		returnEmptyChat(w, http.StatusNotFound)
		return
	}

	// Change the photo of the group
	err := rt.db.SetNewGroupPhoto(chatID, queryInput.NewPhoto)

	if err != nil {
		returnEmptyChat(w, http.StatusNotFound)
		return
	}

	chat.Photo = queryInput.NewPhoto

	w.WriteHeader(http.StatusOK)
	jsonReturn, _ := json.Marshal(chat)
	_, _ = w.Write(jsonReturn)
}
