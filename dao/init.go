package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"theychat/config"
	"theychat/model"
)

var DB *gorm.DB
var Machine *StoreMachine

func Init() {
	db, err := gorm.Open(mysql.Open(config.Conf.MysqlConfig.DNS), &gorm.Config{})
	DB = db
	err = DB.AutoMigrate(&model.User{}, &model.Group{}, &model.Message{})
	if err != nil {
		panic(err.Error())
	}

	Machine = &StoreMachine{
		MessageChan: make(chan *Message, 1000),
		GroupChan:   make(chan *Group, 100),
	}
	go Machine.Run()
}
