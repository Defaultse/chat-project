package app

import (
	"chat-project-go/pkg/websocket"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Services) ServeWs(pool *websocket.Pool, conn *gin.Context) {
	fmt.Println("WebSocket Endpoint")
	wsConn, err := websocket.Upgrade(conn.Writer, conn.Request)

	if err != nil {
		fmt.Fprintf(conn.Writer, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: wsConn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
