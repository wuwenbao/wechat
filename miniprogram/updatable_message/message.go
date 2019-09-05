package updatable_message

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type UpdatableMessage struct {
	miniprogram.Confer
}

//New 动态消息
func New(c miniprogram.Confer) *UpdatableMessage {
	return &UpdatableMessage{c}
}
