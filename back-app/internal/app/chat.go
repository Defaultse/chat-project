package app

// func (s *Services) ChatReader(conn *gin.Context) {
// 	fmt.Println("here")

// 	ws, err := upgrader.Upgrade(conn.Writer, conn.Request, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer ws.Close()

// 	go func(ws *websocket.Conn) {
// 		for {
// 			//Read Message from client
// 			mt, message, err := ws.ReadMessage()
// 			if err != nil {
// 				fmt.Println(err)
// 				break
// 			}
// 			//If client message is ping will return pong
// 			if string(message) == "ping" {
// 				message = []byte("pong")
// 			}
// 			//Response message to client
// 			err = ws.WriteMessage(mt, message)
// 			if err != nil {
// 				fmt.Println(err)
// 				break
// 			}
// 		}
// 	}(ws)
// }

// func wsEndoint(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool {
// 		return true
// 	}

// 	ws, err := upgrader.Upgrade(w, r, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("Client Successfully Connected...")

// 	err = ws.WriteMessage(1, []byte("Hello, connected!"))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	Reader(ws)
// }
