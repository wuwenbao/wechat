package work

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/wuwenbao/wechat/util"
)

type Confer interface {
	CorpId() (corpId string)
	CorpSecret() (corpSecret string)
	Token() (token string, err error)
}

type conf struct {
	corpId     string
	corpSecret string
	tokenFunc  util.TokenFunc
}

func Conf(corpId, corpSecret string, tf util.TokenFunc) *conf {
	c := &conf{
		corpId:     corpId,
		corpSecret: corpSecret,
		tokenFunc:  tf,
	}
	if c.tokenFunc == nil {
		c.tokenFunc = c.defaultTokenFunc()
	}
	return c
}

func (c *conf) CorpId() string {
	return c.corpId
}

func (c *conf) CorpSecret() string {
	return c.corpSecret
}

func (c *conf) Token() (string, error) {
	return c.tokenFunc(c.CorpId(), c.CorpSecret())
}

func (c *conf) defaultTokenFunc() util.TokenFunc {
	token := new(Token)
	mutex := new(sync.Mutex)
	return func(corpId, corpSecret string) (s string, e error) {
		mutex.Lock()
		defer mutex.Unlock()
		if token == nil || time.Now().After(token.ExpiresAt) {
			at, err := getAccessToken(corpId, corpSecret)
			if err != nil {
				return "", err
			}
			token = at
			return at.AccessToken, nil
		}
		return token.AccessToken, nil
	}
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   time.Time
}

func getAccessToken(corpId, corpSecret string) (*Token, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s`, corpId, corpSecret)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := new(Token)
	if err := util.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	data.ExpiresAt = time.Now().Add(time.Second * 7000)
	return data, nil
}
