package http

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type SSEHub struct {
	clients map[*gin.Context]chan interface{}
	mu      sync.Mutex
}

func NewSSEHub() *SSEHub {
	return &SSEHub{
		clients: make(map[*gin.Context]chan interface{}),
	}
}

func (hub *SSEHub) AddClient(c *gin.Context, messageChannel chan interface{}) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[c] = messageChannel
}

func (hub *SSEHub) RemoveClient(c *gin.Context) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	delete(hub.clients, c)
}

func (hub *SSEHub) Broadcast(message interface{}) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for _, messageChannel := range hub.clients {
		messageChannel <- message
	}
}
