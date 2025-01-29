package core

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	Clients  map[*websocket.Conn]bool
	Mutex    sync.Mutex
	Upgrader websocket.Upgrader
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[*websocket.Conn]bool),
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // 必要ならここでオリジン制限を加える
			},
		},
	}
}

func (s *Server) BroadcastMessage(sender *websocket.Conn, message string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	for client := range s.Clients {
		if client != sender {
			if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Printf("Error broadcasting message: %v", err)
				client.Close()
				delete(s.Clients, client)
			}
		}
	}
}
