package miniprogram

import (
	"io"

	"github.com/wuwenbao/wechat/miniprogram/uniform_message"
)

type UniFormMessage struct {
	Confer
}

//Send 下发小程序和公众号统一的服务消息
func (u *UniFormMessage) Send(body io.Reader) error {
	token, err := u.Token()
	if err != nil {
		return err
	}

	return uniform_message.Send(token, body)
}

//NewUniFormMessage 统一服务消息
func NewUniFormMessage(c Confer) *UniFormMessage {
	return &UniFormMessage{c}
}
