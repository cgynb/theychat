package model

import "gorm.io/gorm"

type Group struct {
	GroupId   string `json:"group_id"`
	UserId    uint   `json:"user_id"`
	GroupType int8   `json:"group_type"`
	Specify   uint   `json:"specify"`
	gorm.Model
}

type Message struct {
	GroupId string `json:"group_id"`
	UserId  uint   `json:"user_id"`
	Text    string `json:"text"`
	gorm.Model
}

func (g *Group) IsSingleChat() bool {
	return g.GroupType == 1
}

func (g *Group) IsGroupChat() bool {
	return g.GroupType == 2
}

func (g *Group) IsSpecify(userId uint) bool {
	return g.Specify == userId
}
