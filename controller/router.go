package controller

import (
	"github.com/gin-gonic/gin"
	"theychat/middleware"
)

func Router() (r *gin.Engine) {
	r = gin.Default()

	ping := r.Group("/")
	ping.GET("/", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	chat := v1.Group("/chat")

	user.POST("/login", UserLogin)
	user.POST("/register", UserRegister)

	chat.Use(middleware.ChatAuth())
	chat.GET("", WsHandler)
	chat.GET("/message", ChatMessage)
	chat.GET("/group", ChatGroup)

	return
}
