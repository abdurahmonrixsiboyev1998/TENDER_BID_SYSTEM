package notification

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections; modify this based on your needs
	},
}

type PushNotification struct {
	clients   map[int]map[*websocket.Conn]bool
	mutex     sync.RWMutex
	broadcast chan Message
}

type Message struct {
	UserID  int    `json:"user_id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func NewPushNotification() *PushNotification {
	return &PushNotification{
		clients:   make(map[int]map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}
}

func (pn *PushNotification) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Use the upgrader here
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade error: %v\n", err)
		return
	}
	defer ws.Close()

	userID := r.Context().Value("userID").(int)

	pn.mutex.Lock()
	if pn.clients[userID] == nil {
		pn.clients[userID] = make(map[*websocket.Conn]bool)
	}
	pn.clients[userID][ws] = true
	pn.mutex.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			pn.mutex.Lock()
			delete(pn.clients[userID], ws)
			pn.mutex.Unlock()
			break
		}
		pn.broadcast <- msg
	}
}

func (pn *PushNotification) HandleMessages() {
	for {
		msg := <-pn.broadcast
		pn.mutex.RLock()
		if clients, ok := pn.clients[msg.UserID]; ok {
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
		pn.mutex.RUnlock()
	}
}

func (pn *PushNotification) SendNotification(userID int, notificationType, content string) {
	msg := Message{
		UserID:  userID,
		Type:    notificationType,
		Content: content,
	}
	pn.broadcast <- msg
}
