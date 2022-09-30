package app

import "chat-project-go/pkg/websocket"

func (s *Services) SendMessageToUser(message websocket.Message) {
	// message.Body
	// s.chatService.SendMessageToConversation()
	// conn := s.wsPool.Clients["2016"]
	// conn.WriteMessage()
}

func (s *Services) StartConversationWithUser(message websocket.Message) {

}

func (s *Services) GetConversationMsgs(message websocket.Message) {

}

func (s *Services) GetAllChatsLastMsg(message websocket.Message) {

}
