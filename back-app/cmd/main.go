package main

import (
	"chat-project-go/internal/app"
	"chat-project-go/internal/drivers/mssql"
	"chat-project-go/internal/repository"
	"chat-project-go/internal/service"
	"chat-project-go/pkg/websocket"

	"github.com/gin-gonic/gin"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

func main() {
	userRepository := repository.NewUserRepository(mssql.Connect)

	jwtTokenService := service.NewTokenManager(signingKey)
	authService := service.NewAuthService(jwtTokenService, userRepository)
	chatService := service.NewChatService()
	services := app.NewServices(authService, chatService)

	pool := websocket.NewPool()
	go pool.Start()

	router := gin.Default()

	router.POST("/register/", services.Register)
	router.POST("/login/", services.Login)
	// router.GET("/ws", websocket.ServeWs)
	router.GET("/ws", func(ctx *gin.Context) {
		services.ServeWs(pool, ctx)
	})
	router.GET("/ws/chat", func(ctx *gin.Context) {
		services.ServeWs(pool, ctx)
	})
	router.Run()
}
