package main

import (
	"net/http"

	"tchat.com/server/router"
)

func main() {
	h := router.GetHandler()

	http.ListenAndServe(":8080", h)
}
