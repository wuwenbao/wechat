package work

import (
	"github.com/wuwenbao/wechat/token"
)

type Confer interface {
	CorpId() (corpId string)
	CorpSecret() (corpSecret string)
	Token() (token string, err error)
}

type conf struct {
	corpId       string
	corpSecret   string
	tokenAdapter token.Adapter
}

func Conf(corpId, corpSecret string, tokenAdapter token.Adapter) *conf {
	c := &conf{
		corpId:       corpId,
		corpSecret:   corpSecret,
		tokenAdapter: tokenAdapter,
	}
	if c.tokenAdapter == nil {
		c.tokenAdapter = token.DefaultToken()
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
	return c.tokenAdapter.GetToken(c.CorpId(), c.CorpSecret())
}
