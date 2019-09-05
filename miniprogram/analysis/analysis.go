package analysis

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type Analysis struct {
	miniprogram.Confer
}

//New 数据分析
func New(c miniprogram.Confer) *Analysis {
	return &Analysis{Confer: c}
}
