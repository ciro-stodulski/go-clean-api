package http

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type SSEHub struct {
	clients map[*gin.Context]chan any
	mu      sync.Mutex
}

func NewSSEHub() *SSEHub {
	return &SSEHub{
		clients: make(map[*gin.Context]chan any),
	}
}

func (hub *SSEHub) AddClient(c *gin.Context, messageChannel chan any) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[c] = messageChannel
}

func (hub *SSEHub) RemoveClient(c *gin.Context) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	delete(hub.clients, c)
}

func (hub *SSEHub) Broadcast(message any) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for _, messageChannel := range hub.clients {
		messageChannel <- message
	}
}
