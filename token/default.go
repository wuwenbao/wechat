package token

import (
	"sync"
	"time"
)

type defaultToken struct {
	token  string
	expire time.Time
	mutex  *sync.Mutex
}

func (d *defaultToken) GetToken(appId, secret string) (token string, err error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.token == "" || time.Now().After(d.expire) {
		token, err := GetAccessToken(appId, secret)
		if err != nil {
			return "", err
		}
		d.token = token
		d.expire = time.Now().Add(time.Second * 7000)
		return token, nil
	}
	return d.token, nil
}

func DefaultToken() *defaultToken {
	return &defaultToken{
		mutex: new(sync.Mutex),
	}
}
