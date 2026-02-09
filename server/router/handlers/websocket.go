package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"tchat.com/server/utils"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	mu sync.Mutex
)

func (h *Handlers) WebsocketChat(w http.ResponseWriter, r *http.Request) {
	chatID := utils.ChatID(strings.TrimLeft(r.URL.Path, "/ws/chat/"))
	if chatID == "" {
		WriteBadRequest(w, errors.New("chatID not provided"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		WriteInternalServerError(w, fmt.Errorf("cannot upgrade: %w", err))
		return
	}
	defer conn.Close()

	mu.Lock()
	h.conn[chatID] = append(h.conn[chatID], conn)
	mu.Unlock()

	for m := range h.newMessages[chatID] {
		for _, conn := range h.conn[chatID] {
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
