package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	mu sync.Mutex
)

func (h *Handlers) WebsocketChat(w http.ResponseWriter, r *http.Request) {
	mergedIDs := strings.TrimLeft(r.URL.Path, "/ws/chat/")
	if mergedIDs == "" {
		WriteBadRequest(w, errors.New("mergedIDs not provided"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		WriteInternalServerError(w, fmt.Errorf("cannot upgrade: %w", err))
		return
	}
	defer conn.Close()
	h.conn = append(h.conn, conn)

	for m := range h.newMessages {
		for _, conn := range h.conn {
			bytes, err := json.Marshal(m)
			if err != nil {
				fmt.Printf("cannot marshal message: %v\n", err)
				continue
			}

			if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
				fmt.Printf("cannot write message to socket: %v\n", err)
			}
		}
	}
}
