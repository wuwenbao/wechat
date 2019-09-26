package work

import (
	"strings"

	"github.com/wuwenbao/wechat/work/appchat"
)

type AppChat struct {
	Confer
}

//Create 创建群聊会话
func (u *AppChat) Create(param *appchat.ParamCreate) (*appchat.ResponseCreate, error) {
	token, err := u.Token()
	if err != nil {
		return nil, err
	}
	return appchat.Create(token, param)
}

//Update 修改群聊会话
func (u *AppChat) Update(param *appchat.ParamUpdate) error {
	token, err := u.Token()
	if err != nil {
		return err
	}
	return appchat.Update(token, param)
}

//Get 获取群聊会话
func (u *AppChat) Get(chatid string) (*appchat.ResponseGet, error) {
	token, err := u.Token()
	if err != nil {
		return nil, err
	}
	return appchat.Get(token, chatid)
}

//Send 应用推送消息
func (u *AppChat) Send(chatid string, msg appchat.Sender) error {
	token, err := u.Token()
	if err != nil {
		return err
	}
	return appchat.Send(token, strings.NewReader(msg.ChatSend(chatid)))
}

func NewAppChat(c Confer) *AppChat {
	return &AppChat{
		Confer: c,
	}
}
