package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"tchat.com/server/router/handlers"
	"tchat.com/server/store"
)

func GetHandler() http.Handler {
	h := handlers.NewHandler(store.NewCache())
	r := mux.NewRouter()

	r.HandleFunc("/users", h.FindUserByID).Methods("GET")
	r.HandleFunc("/users", h.SaveUser).Methods("POST")
	r.HandleFunc("/messages", h.ReadChat).Methods("GET")
	r.HandleFunc("/messages", h.SendMessage).Methods("POST")

	return r
}
