package work

import (
	"strings"

	"github.com/wuwenbao/wechat/work/message"
)

type Message struct {
	Confer
	*message.SendParam
}

//ToUser 成员ID列表
func (m *Message) ToUser(user ...string) *Message {
	m.SendParam.ToUser = strings.Join(user, "|")
	return m
}

//ToParty 部门ID列表
func (m *Message) ToParty(party ...string) *Message {
	m.SendParam.ToParty = strings.Join(party, "|")
	return m
}

//ToTag 标签ID列表
func (m *Message) ToTag(tag ...string) *Message {
	m.SendParam.ToTag = strings.Join(tag, "|")
	return m
}

//Safe 	表示是否是保密消息
func (m *Message) Safe(b bool) *Message {
	if b {
		m.SendParam.Safe = 1
	}
	return m
}

//EnableIDTrans 表示是否开启id转译
func (m *Message) EnableIDTrans(b bool) *Message {
	if b {
		m.SendParam.EnableIDTrans = 1
	}
	return m
}

//Send 发送应用消息
func (m *Message) Send(msg message.Sender) (*message.SendResponse, error) {
	token, err := m.Token()
	if err != nil {
		return nil, err
	}
	return message.Send(token, strings.NewReader(msg.Send(m.SendParam)))
}

//NewMessage 发送应用消息
func NewMessage(c Confer, agentId int) *Message {
	return &Message{
		Confer: c,
		SendParam: &message.SendParam{
			AgentId: agentId,
		},
	}
}
