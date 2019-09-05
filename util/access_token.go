package util

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   time.Time
}

//Token 接口调用凭证
func GetAccessToken(appid, secret string) (*Token, error) {
	apiUrl := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, appid, secret)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := new(Token)
	if err := ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	data.ExpiresAt = time.Now().Add(time.Second * 7000)
	return data, nil
}

type TokenFunc func(appId, secret string) (string, error)

//defaultTokenFunc 默认token方法 暂时需要保留不能删除
func DefaultTokenFunc() TokenFunc {
	tokenSyncMap := new(sync.Map)
	return func(appid, secret string) (s string, e error) {
		var (
			value interface{}
			ok    bool
			token Token
		)
		f := func() (s string, e error) {
			at, err := GetAccessToken(appid, secret)
			if err != nil {
				return "", err
			}
			tokenSyncMap.Store(appid, *at)
			return at.AccessToken, nil
		}
		if value, ok = tokenSyncMap.Load(appid); !ok {
			return f()
		}
		if token, ok = value.(Token); !ok {
			return f()
		}
		if time.Now().After(token.ExpiresAt) {
			return f()
		}
		return token.AccessToken, nil
	}
}
