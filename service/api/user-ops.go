package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// The doLogin function performs the login logic
//
// TODO: Implement the 403 error by checking for sessions (might want to delete it (?))
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Fetch the body of the request and unmarshal in variable
	queryBody, _ := io.ReadAll(r.Body)
	var resCode = http.StatusOK
	var queryInput struct {
		UserName string `json:"name"`
	}
	_ = json.Unmarshal(queryBody, &queryInput)

	// Check if the user is in the DB
	dbUser, dbErr := rt.db.GetUserByName(queryInput.UserName)

	// If the user does not exist, create it
	if errors.Is(dbErr, sql.ErrNoRows) {
		// Check for name length
		if len(queryInput.UserName) >= 3 && len(queryInput.UserName) <= 16 {
			// Register in the DB
			dbUser, _ = rt.db.RegisterNewUser(queryInput.UserName)
			resCode = http.StatusCreated
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	}

	// Respond to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)

	jsonReturn, _ := json.Marshal(dbUser)
	_, _ = w.Write([]byte(jsonReturn))
}

// The setMyUserName function changes the username of the user in the URI path
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Take the user ID from the path of the request
	userId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/user/"), "/id")
	var queryInput struct {
		NewUserName string `json:"newUsername"`
	}

	queryBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(queryBody, &queryInput)

	// Check for new username length
	if len(queryInput.NewUserName) >= 3 && len(queryInput.NewUserName) <= 16 {
		_ = rt.db.SetNewUserName(userId, queryInput.NewUserName)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusTeapot)
	}
}

// The setMyUserName function changes the username of the user in the URI path
func (rt *_router) setMyDisplayName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Take the user ID from the path of the request
	userId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/user/"), "/name")
	var queryInput struct {
		NewDisplayName string `json:"newDisplayName"`
	}

	queryBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(queryBody, &queryInput)

	// Check for new display name length
	if len(queryInput.NewDisplayName) >= 3 && len(queryInput.NewDisplayName) <= 32 {
		_ = rt.db.SetNewDisplayName(userId, queryInput.NewDisplayName)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusTeapot)
	}
}

// setMyBio sets a new biography
func (rt *_router) setMyBio(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Take the user ID from the path of the request
	userId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/user/"), "/bio")
	var queryInput struct {
		NewBio string `json:"newBio"`
	}

	queryBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(queryBody, &queryInput)

	// Check for new biography length
	if len(queryInput.NewBio) >= 1 && len(queryInput.NewBio) <= 512 {
		_ = rt.db.SetNewBiography(userId, queryInput.NewBio)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// getMyConversations retrieves all the chats where the user belongs
func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Take the user ID from the path of the request
	userID := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/user/"), "/chats")

	// Check if user exists, in case return 404
	_, err := rt.db.GetUserByName(userID)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
	}

	// Retrieve chats
	allChats := rt.db.GetConversations(userID)

	jsonChat, _ := json.Marshal(allChats)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonChat)
}
