package service

type ChatService interface {
}

type chatService struct{}

func NewChatService() *chatService {
	return &chatService{}
}

func (c chatService) GetChat() {

}

func (c chatService) GetChatMessages() {

}
