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
)

type Client struct {
	*model.User
	Conn *websocket.Conn
}

type InputMessage struct {
	To      string `json:"to"` //  GroupId
	Action  int8   `json:"action"`
	Join    string `json:"join"` // target join GroupId
	Message string `json:"message"`
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
		resp = RespOk(om)
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
		})
		To = append(To, client.ID)
		dao.Machine.GroupChan <- &dao.Group{
			UserId:  client.ID,
			GroupId: groupId,
		}
	case AddGroupChat:
		groupId := cache.AddGroup(client.ID)
		resp = RespOk(gin.H{
			"group_id": groupId,
		})
		To = append(To, client.ID)
		dao.Machine.GroupChan <- &dao.Group{
			UserId:  client.ID,
			GroupId: groupId,
		}
	case JoinSingleChat:
		ok := cache.JoinGroup(config.Conf.DetailConfig.SingleChatCap, imsg.Join, client.ID)
		if ok {
			resp = RespOk(gin.H{
				"join": client.User,
			})
			To = append(cache.GetGroupMember(imsg.Join)) // send joiner info
		} else {
			resp = RespErr(403, "this group is full")
			To = append(To, client.ID)
		}
	case JoinGroupChat:
		ok := cache.JoinGroup(config.Conf.DetailConfig.GroupChatCap, imsg.Join, client.ID)
		if ok {
			resp = RespOk(gin.H{
				"join": client.User,
			})
			To = append(cache.GetGroupMember(imsg.Join)) // send joiner info
		} else {
			resp = RespErr(403, "this group is full")
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
	gs, ok := dao.GetGroup(args.Page, args.UserId)
	if ok {
		resp = RespOk(gin.H{
			"groups": gs,
		})
	} else {
		resp = RespErr(400, "fail")
	}
	return
}
