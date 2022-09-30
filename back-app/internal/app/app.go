package app

import (
	"chat-project-go/internal/service"
	"chat-project-go/pkg/websocket"
)

type Services struct {
	authService service.AuthService
	chatService service.ChatService
	wsPool      *websocket.Pool
}

func NewServices(
	authService service.AuthService,
	chatService service.ChatService,
	wsPool *websocket.Pool,
) *Services {
	return &Services{
		authService: authService,
		chatService: chatService,
		wsPool:      wsPool,
	}
}
