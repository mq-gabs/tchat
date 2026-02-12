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
	defer func() {
		conn.Close()
		mu.Lock()
		h.websocketConnService.RemoveConn(conn, chatID)
		mu.Unlock()
	}()

	h.websocketConnService.SaveConn(conn, chatID)

	for m := range h.messagesService.GetMessagesChannel(chatID) {
		for _, conn := range h.websocketConnService.GetConns(chatID) {
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
