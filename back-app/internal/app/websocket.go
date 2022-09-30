package app

import (
	"chat-project-go/pkg/websocket"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Services) ServeWs(pool *websocket.Pool, conn *gin.Context, userId int64) {
	wsConn, err := websocket.Upgrade(conn.Writer, conn.Request)

	if err != nil {
		fmt.Fprintf(conn.Writer, "%+v\n", err)
	}

	userIdStr := strconv.Itoa(int(userId))

	client := &websocket.Client{
		ID:   userIdStr,
		Conn: wsConn,
		Pool: pool,
	}

	pool.Register <- client

	for {
		msgType, p, err := client.Conn.ReadMessage()

		if err != nil || msgType == 0 {
			pool.Unregister <- client
			client.Conn.Close()
			log.Println(err)
			break
		}

		result := make(map[string]interface{})
		json.Unmarshal([]byte(p), &result)

		message := websocket.Message{UserID: userIdStr, Body: result}

		go s.ProcessMessage(message)

		client.Conn.WriteJSON(message)
	}
}

func (s *Services) ProcessMessage(message websocket.Message) {
	switch message.Body["type"] {
	case 0:
		s.GetAllChatsLastMsg(message)
		// Start new conversation
	case 1:
		// Send message to user
		s.SendMessageToUser(message)
	case 2:
		s.GetConversationMsgs(message)
	case 3:
		// Start new conversation
		s.StartConversationWithUser(message)
	}
}
