package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"theychat/cache"
	"theychat/config"
	"theychat/dao"
	"theychat/model"
	"time"
)

// action
const (
	SendMsg = iota
	AddSingleChat
	AddGroupChat
	JoinSingleChat
	JoinGroupChat
	Invite
)

// group type
const (
	SingleChat = 1
	GroupChat  = 2
)

type Client struct {
	*model.User
	Conn *websocket.Conn
}

/*
To: group id
Action: action
Join: target group's id
Message: message
Specify: userId, specify a person when raising a single chat | invite person id
GroupId: invite groupid
*/
type InputMessage struct {
	To      string `json:"to"`
	Action  int8   `json:"action"`
	Join    string `json:"join"`
	Message string `json:"message"`
	Specify uint   `json:"specify"`
	GroupId string `json:"group_id"`
}

type OutputMessage struct {
	From    *model.User `json:"from"` // UserId
	To      string      `json:"to"`   // GroupId
	Message string      `json:"message"`
	Time    time.Time   `json:"time"`
}

type ChatMessageArgs struct {
	UserId  uint
	Page    int    `form:"page"`
	GroupId string `form:"group_id"`
}

type ChatGroupArgs struct {
	UserId uint
	Page   int `form:"page"`
}

func Chat(client *Client, imsg *InputMessage) (resp *Resp, To []uint) {
	switch imsg.Action {
	case SendMsg:
		om := OutputMessage{
			From:    client.User,
			To:      imsg.To,
			Message: imsg.Message,
			Time:    time.Now(),
		}
		resp = RespOk(gin.H{
			"message": om,
			"action":  imsg.Action,
		})
		To = append(cache.GetGroupMember(imsg.To))
		dao.Machine.MessageChan <- &dao.Message{
			UserId:  client.ID,
			GroupId: imsg.To,
			Text:    imsg.Message,
		}
	case AddSingleChat:
		groupId := cache.AddGroup(client.ID)
		resp = RespOk(gin.H{
			"group_id": groupId,
			"action":   imsg.Action,
		})
		To = append(To, client.ID)
		dao.Machine.GroupChan <- &dao.Group{
			UserId:    client.ID,
			GroupId:   groupId,
			GroupType: SingleChat,
			Specify:   imsg.Specify,
		}
	case AddGroupChat:
		groupId := cache.AddGroup(client.ID)
		resp = RespOk(gin.H{
			"group_id": groupId,
			"action":   imsg.Action,
		})
		To = append(To, client.ID)
		dao.Machine.GroupChan <- &dao.Group{
			UserId:    client.ID,
			GroupId:   groupId,
			GroupType: GroupChat,
		}
	case JoinSingleChat:
		group, ok := dao.GetGroup(imsg.Join)
		ok = cache.JoinGroup(config.Conf.DetailConfig.SingleChatCap, imsg.Join, client.ID) &&
			group.IsSingleChat() &&
			group.IsSpecify(client.ID)
		if ok {
			resp = RespOk(gin.H{
				"join":   client.User,
				"action": imsg.Action,
			})
			To = append(cache.GetGroupMember(imsg.Join)) // send joiner info
		} else {
			resp = RespErr(403, "join fail")
			To = append(To, client.ID)
		}
	case JoinGroupChat:
		group, ok := dao.GetGroup(imsg.Join)
		ok = cache.JoinGroup(config.Conf.DetailConfig.GroupChatCap, imsg.Join, client.ID) && group.IsGroupChat()
		if ok {
			resp = RespOk(gin.H{
				"join":   client.User,
				"action": imsg.Action,
			})
			To = append(cache.GetGroupMember(imsg.Join)) // send joiner info
		} else {
			resp = RespErr(403, "join fail")
			To = append(To, client.ID)
		}
	case Invite:
		ok := cache.HasGroup(imsg.GroupId)
		if ok {
			resp = RespOk(gin.H{
				"group_id": imsg.GroupId,
				"action":   imsg.Action,
			})
			To = append(To, imsg.Specify)
		} else {
			resp = RespErr(403, "no such group")
			To = append(To, client.ID)
		}
	default:
		resp = RespErr(403, "error action")
		To = append(To, client.ID)
	}
	return
}

func ChatMessage(args *ChatMessageArgs) (resp *Resp) {
	if ok := cache.IsMember(args.UserId, args.GroupId); ok {
		ms, _ := dao.GetMessages(args.Page, args.GroupId)
		resp = RespOk(gin.H{
			"messages": ms,
		})
	} else {
		resp = RespErr(403, "you are not in this group")
	}
	return
}

func ChatGroup(args *ChatGroupArgs) (resp *Resp) {
	gs, ok := dao.GetGroups(args.Page, args.UserId)
	if ok {
		resp = RespOk(gin.H{
			"groups": gs,
		})
	} else {
		resp = RespErr(400, "fail")
	}
	return
}
