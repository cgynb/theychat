package dao

import (
	"theychat/config"
	"theychat/model"
)

func GetMessages(page int, groupId string) (ms []*model.Message, ok bool) {
	ok = true
	result := DB.Where("group_id = ?", groupId).
		Limit(config.Conf.DetailConfig.PageSize).Offset(config.Conf.DetailConfig.PageSize * (page - 1)).Find(&ms)
	if result.Error != nil {
		return nil, false
	}
	return
}

func CreateMessage(userId uint, groupId, text string) (m *model.Message, ok bool) {
	m = &model.Message{
		GroupId: groupId,
		UserId:  userId,
		Text:    text,
	}
	ok = true
	result := DB.Create(m)
	if result.Error != nil {
		return nil, false
	}
	return
}

func CreateGroup(userId uint, groupId string) (g *model.Group, ok bool) {
	ok = true
	g = &model.Group{
		GroupId: groupId,
		UserId:  userId,
	}
	result := DB.Create(g)
	if result.Error != nil {
		return nil, false
	}
	return
}

func GetGroup(page int, userId uint) (gs []*model.Group, ok bool) {
	ok = true
	result := DB.Where("user_id = ?", userId).
		Limit(config.Conf.DetailConfig.PageSize).Offset(config.Conf.DetailConfig.PageSize * (page - 1)).Find(&gs)
	if result.Error != nil {
		return nil, false
	}
	return
}
