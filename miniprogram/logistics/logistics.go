package logistics

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type Logistics struct {
	miniprogram.Confer
}

//New 物流助手
func New(c miniprogram.Confer) *Logistics {
	return &Logistics{Confer: c}
}
