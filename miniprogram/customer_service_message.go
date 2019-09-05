package miniprogram

import (
	"io"

	"github.com/wuwenbao/wechat/miniprogram/customer_service_message"
)

type CustomerServiceMessage struct {
	Confer
}

//Send 发送消息
func (u *CustomerServiceMessage) Send(body io.Reader) error {
	token, err := u.Token()
	if err != nil {
		return err
	}
	return customer_service_message.Send(token, body)
}

//New 模版消息
func NewCustomerServiceMessage(c Confer) *CustomerServiceMessage {
	return &CustomerServiceMessage{ Confer: c}
}
