package security

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type Security struct {
	miniprogram.Confer
}

//New 内容安全
func New(c miniprogram.Confer) *Security {
	return &Security{c}
}
