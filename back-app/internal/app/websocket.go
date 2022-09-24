package app

import (
	"chat-project-go/pkg/websocket"
	"fmt"

	"github.com/gin-gonic/gin"
)

// func ServeWs(conn *gin.Context) {
// 	ws, err := websocket.Upgrade(conn.Writer, conn.Request)

// 	if err != nil {
// 		fmt.Fprintf(conn.Writer, "%+V\n", err)
// 	}

// 	go websocket.Writer(ws)

// 	websocket.Reader(ws)
// }

func (s *Services) ServeWs(pool *websocket.Pool, conn *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
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
