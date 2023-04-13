package model

import "gorm.io/gorm"

type Group struct {
	GroupId string `json:"group_id"`
	UserId  uint   `json:"user_id"`
	gorm.Model
}

type Message struct {
	GroupId string `json:"group_id"`
	UserId  uint   `json:"user_id"`
	Text    string `json:"text"`
	gorm.Model
}
