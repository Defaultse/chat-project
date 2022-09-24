package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Services) Login(c *gin.Context) {
	userEmail, userPassword, _ := c.Request.BasicAuth()

	fmt.Println(userEmail, userPassword)
	jwtToken, err := s.authService.SignIn(userEmail, userPassword)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, jwtToken)
}
