package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve data from database
}
