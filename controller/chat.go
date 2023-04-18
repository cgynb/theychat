package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"theychat/model"
	"theychat/service"
	"time"
)

var upGrader = websocket.Upgrader{
	WriteBufferSize:  1024,
	ReadBufferSize:   1024,
	HandshakeTimeout: time.Hour,
	CheckOrigin: func(r *http.Request) bool {
		//c, ok := utils.ParseToken(r.Header.Get("Authorization"))
		//id := strconv.Itoa(int(userClaims.ID))
		//user, ok := dao.GetUser("id", id)
		return true
	},
}

func WsHandler(c *gin.Context) {
	var client *service.Client
	var user *model.User
	var imsg *service.InputMessage
	var resp *service.Resp
	var to []uint
	imsg = new(service.InputMessage)

	// bind conn, user to client
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*model.User)
	} else {
		return
	}
	client = &service.Client{
		User: user,
		Conn: conn,
	}
	service.ConnPool.Set(client.ID, client.Conn)
	defer disconnect(client)

	for {
		err := conn.ReadJSON(imsg)
		if err != nil {
			return
		}
		resp, to = service.Chat(client, imsg)
		broadCast(resp, to)
	}
}

func disconnect(client *service.Client) {
	client.Conn.Close()
	if _, notEmpty := service.ConnPool.Get(client.ID); notEmpty {
		service.ConnPool.Del(client.ID)
	}
}

func broadCast(resp *service.Resp, to []uint) {
	for _, clientId := range to {
		// send msg
		if val, ok := service.ConnPool.Get(clientId); ok && val != nil {
			conn, _ := val.(*websocket.Conn)
			conn.WriteJSON(resp)
		} else {
			fmt.Println("err", clientId)
		}
	}
}

func ChatMessage(c *gin.Context) {
	var args service.ChatMessageArgs
	var user *model.User
	err := c.ShouldBind(&args)
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*model.User)
		args.UserId = user.ID
	}
	if err != nil {
		c.JSON(400, "err args")
	}
	resp := service.ChatMessage(&args)
	c.JSON(200, resp)
}

func ChatGroup(c *gin.Context) {
	var args service.ChatGroupArgs
	var user *model.User
	err := c.ShouldBind(&args)
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*model.User)
		args.UserId = user.ID
	}
	if err != nil {
		c.JSON(400, "err args")
	}
	resp := service.ChatGroup(&args)
	c.JSON(200, resp)
}
