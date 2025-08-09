package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"wasatext/service/customstructs"

	"github.com/julienschmidt/httprouter"
)

// The doLogin function performs the login logic
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Fetch the body of the request and unmarshal in variable
	queryBody, _ := io.ReadAll(r.Body)
	var requestUser customstructs.User
	_ = json.Unmarshal(queryBody, &requestUser)

	// Check if the user is in the DB
	dbUser, dbErr := rt.db.GetUserByName(requestUser.Name)

	// If the user does not exist, create it
	if errors.Is(dbErr, sql.ErrNoRows) {
		// Check for name length
		if len(requestUser.Name) > 3 && len(requestUser.Name) < 16 {
			// Register in the DB
			dbUser, _ = rt.db.RegisterNewUser(requestUser.Name)
		}
	}

	// Respond to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	jsonReturn, _ := json.Marshal(dbUser)
	_, _ = w.Write(jsonReturn)
}
