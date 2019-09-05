package soter

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type Soter struct {
	miniprogram.Confer
}

//New 生物认证
func New(c miniprogram.Confer) *Soter {
	return &Soter{c}
}
