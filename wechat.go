package wechat

import (
	"github.com/wuwenbao/wechat/minigame"
	"github.com/wuwenbao/wechat/miniprogram"
	"github.com/wuwenbao/wechat/offiaccount"
	"github.com/wuwenbao/wechat/oplatform"
	"github.com/wuwenbao/wechat/payment"
	"github.com/wuwenbao/wechat/work"
)

//Offiaccount 公众号
func Offiaccount() *offiaccount.App {
	return offiaccount.New()
}

//Oplatform 开发平台
func Oplatform() *oplatform.App {
	return oplatform.New()
}

//Miniprogram 小程序
func Miniprogram(c miniprogram.Confer) *miniprogram.App {
	return miniprogram.New(c)
}

//Minigame 小游戏
func Minigame() *minigame.App {
	return minigame.New()
}

//Payment 微信支付
func Payment(c payment.Confer) *payment.App {
	return payment.New(c)
}

//Work 企业微信
func Work(c work.Confer) *work.App {
	return work.New(c)
}
