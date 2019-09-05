package miniprogram

import (
	"io"

	"github.com/wuwenbao/wechat/miniprogram/template_message"
)

type TemplateMessage struct {
	Confer
}

//Send 发送消息
func (u *TemplateMessage) Send(body io.Reader) error {
	token, err := u.Token()
	if err != nil {
		return err
	}
	return template_message.Send(token, body)
}

//NewTemplateMessage 模版消息
func NewTemplateMessage(c Confer) *TemplateMessage {
	return &TemplateMessage{c}
}
