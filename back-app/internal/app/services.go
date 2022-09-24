package app

import "chat-project-go/internal/service"

type Services struct {
	authService service.AuthService
	chatService service.ChatService
}

func NewServices(
	authService service.AuthService,
	chatService service.ChatService,
) *Services {
	return &Services{
		authService: authService,
		chatService: chatService,
	}
}
