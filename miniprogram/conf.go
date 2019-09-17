package miniprogram

import (
	"github.com/wuwenbao/wechat/token"
)

type Confer interface {
	Appid() string
	Secret() string
	Token() (string, error)
}

type conf struct {
	appid        string
	secret       string
	tokenAdapter token.Adapter
}

func Conf(appid, secret string, tokenAdapter token.Adapter) *conf {
	c := &conf{
		appid:        appid,
		secret:       secret,
		tokenAdapter: tokenAdapter,
	}
	if c.tokenAdapter == nil {
		c.tokenAdapter = token.DefaultToken()
	}
	return c
}

func (c *conf) Appid() string {
	return c.appid
}

func (c *conf) Secret() string {
	return c.secret
}

func (c *conf) Token() (string, error) {
	return c.tokenAdapter.GetToken(c.Appid(), c.Secret())
}
