package dao

import "fmt"

type Message struct {
	UserId  uint
	GroupId string
	Text    string
}
type Group struct {
	UserId    uint
	GroupId   string
	GroupType int8
	Specify   uint
}
type StoreMachine struct {
	MessageChan chan *Message
	GroupChan   chan *Group
}

func (stm *StoreMachine) Run() {
	fmt.Println("machine running")
	for {
		select {
		case group := <-stm.GroupChan:
			CreateGroup(group.UserId, group.GroupId, group.GroupType, group.Specify)
		case msg := <-stm.MessageChan:
			CreateMessage(msg.UserId, msg.GroupId, msg.Text)
		}
	}
}
