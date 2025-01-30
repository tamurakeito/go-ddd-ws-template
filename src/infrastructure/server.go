package infrastructure

import (
	"go-ddd-ws-template/src/domain/entity"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	Clients  map[entity.Client]bool
	Mutex    sync.Mutex
	Upgrader websocket.Upgrader
}

func NewServer() *Server {
	server := Server{
		Clients: make(map[entity.Client]bool),
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // 必要ならここでオリジン制限を加える
			},
		},
	}
	return &server
}

func (s *Server) BroadcastMessage(sender entity.Client, message string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	for client := range s.Clients {
		if client != sender {
			if err := client.SendMessage(message); err != nil {
				log.Printf("Error broadcasting message: %v", err)
				client.Close()
				delete(s.Clients, client)
			}
		}
	}
}
