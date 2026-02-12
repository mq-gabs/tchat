package service

import (
	"sync"

	"github.com/gorilla/websocket"
	"tchat.com/server/utils"
)

type WebsocketConnService struct {
	mu   sync.Mutex
	conn map[utils.ChatID][]*websocket.Conn
}

func NewWebsocketConnService() *WebsocketConnService {
	return &WebsocketConnService{
		conn: make(map[utils.ChatID][]*websocket.Conn),
	}
}

func (ws *WebsocketConnService) GetConns(chatID utils.ChatID) []*websocket.Conn {
	conns, ok := ws.conn[chatID]
	if ok {
		return conns
	}

	ws.mu.Lock()
	ws.conn[chatID] = make([]*websocket.Conn, 0)
	ws.mu.Unlock()

	return ws.conn[chatID]
}

func (ws *WebsocketConnService) SaveConn(conn *websocket.Conn, chatID utils.ChatID) {
	ws.mu.Lock()
	ws.conn[chatID] = append(ws.conn[chatID], conn)
	ws.mu.Unlock()
}

func (ws *WebsocketConnService) RemoveConn(conn *websocket.Conn, chatID utils.ChatID) {
	ws.conn[chatID] = utils.Filter(ws.conn[chatID], func(c *websocket.Conn) bool {
		return c != conn
	})
}
