package miniprogram

import (
	"sync"
	"time"

	"github.com/wuwenbao/wechat/util"
)

type Confer interface {
	Appid() string
	Secret() string
	Token() (string, error)
}

type conf struct {
	appid     string
	secret    string
	tokenFunc util.TokenFunc
}

func Conf(appid, secret string, tf util.TokenFunc) *conf {
	c := &conf{
		appid:     appid,
		secret:    secret,
		tokenFunc: tf,
	}
	if c.tokenFunc == nil {
		c.tokenFunc = c.defaultTokenFunc()
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
	return c.tokenFunc(c.Appid(), c.Secret())
}

func (c *conf) defaultTokenFunc() util.TokenFunc {
	token := new(util.Token)
	mutex := new(sync.Mutex)
	return func(appid, secret string) (s string, e error) {
		mutex.Lock()
		defer mutex.Unlock()
		if token == nil || time.Now().After(token.ExpiresAt) {
			at, err := util.GetAccessToken(appid, secret)
			if err != nil {
				return "", err
			}
			token = at
			return at.AccessToken, nil
		}
		return token.AccessToken, nil
	}
}
