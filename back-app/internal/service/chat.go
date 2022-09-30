package service

import (
	"chat-project-go/internal/datastruct"
	"chat-project-go/internal/repository"
)

type ChatService interface {
	StartNewConversation()
	GetAllChatsById(userId string) ([]datastruct.ChatsLastMsgs, error)
	SendMessageToConversation(userId int64, conversation_id int64, message string) error
}

type chatService struct {
	chatRepo repository.ChatRepositoryContract
}

func NewChatService(chatRepo repository.ChatRepositoryContract) *chatService {
	return &chatService{
		chatRepo: chatRepo,
	}
}

func (c chatService) StartNewConversation() {

}

func (c chatService) SendMessageToConversation(userId int64, conversation_id int64, message string) error {
	err := c.chatRepo.CreateMessage(userId, conversation_id, message)

	if err != nil {
		return err
	}

	return nil
}

func (c chatService) GetAllChatsById(userId string) ([]datastruct.ChatsLastMsgs, error) {
	allChats, err := c.chatRepo.GetAllChatsByUserId(userId)

	if err != nil {
		return nil, err
	}

	return allChats, nil
}
