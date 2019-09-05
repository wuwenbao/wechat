package miniprogram

import (
	"github.com/wuwenbao/wechat/miniprogram/auth"
)

type Auth struct {
	Confer
}

//Token 获取小程序全局唯一后台接口调用凭据
func (a *Auth) GetAccessToken() (string, error) {
	return a.Token()
}

//Code2Session 登录凭证校验
func (a *Auth) Code2Session(code string) (*auth.Session, error) {
	return auth.Code2Session(a.Appid(), a.Secret(), code)
}

//GetPaidUnionId 用户支付完成后，获取该用户的unionid
func (a *Auth) GetPaidUnionId(openid string, args ...string) (string, error) {
	token, err := a.Token()
	if err != nil {
		return "", err
	}
	return auth.GetPaidUnionId(token, openid, args...)
}

//NewAuth 开放接口
func NewAuth(c Confer) *Auth {
	return &Auth{Confer: c}
}
