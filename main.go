package main

import (
	"theychat/cache"
	"theychat/config"
	"theychat/controller"
	"theychat/dao"
	"theychat/service"
)

func main() {
	config.Init()
	cache.Init()
	dao.Init()
	controller.Init()
	service.Init()
	router := controller.Router()
	_ = router.Run(":8000")
}
