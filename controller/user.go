package controller

import (
	"github.com/gin-gonic/gin"
	"theychat/service"
)

func UserLogin(c *gin.Context) {
	var args service.UserLoginArgs
	err := c.ShouldBind(&args)
	if err != nil {
		c.JSON(400, "err args")
	}
	resp := service.UserLogin(&args)
	c.JSON(200, resp)
}

func UserRegister(c *gin.Context) {
	var args service.UserRegisterArgs
	err := c.ShouldBind(&args)
	if err != nil {
		c.JSON(400, "err args")
	}
	resp := service.UserRegister(&args)
	c.JSON(200, resp)
}
