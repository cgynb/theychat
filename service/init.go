package service

import (
	"theychat/dao"
	"theychat/utils"
)

var ConnPool utils.SafeMap // safe map[uint]*websocket.Conn   (key: userId) - (value: conn)

func Init() {
	ConnPool.Init()
	go dao.Machine.Run()
}
