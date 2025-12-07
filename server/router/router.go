package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"tchat.com/server/router/handlers"
	"tchat.com/server/store"
)

const (
	PathUsersGet     = "/users"
	PathUsersSave    = "/users"
	PathMessagesGet  = "/messages"
	PathMessagesSend = "/messages"
)

func GetHandler() http.Handler {
	h := handlers.NewHandler(store.NewCache())
	r := mux.NewRouter()

	r.HandleFunc(PathUsersGet, h.FindUserByID).Methods("GET")
	r.HandleFunc(PathUsersSave, h.SaveUser).Methods("POST")
	r.HandleFunc(PathMessagesGet, h.ReadChat).Methods("GET")
	r.HandleFunc(PathMessagesSend, h.SendMessage).Methods("POST")

	return r
}
