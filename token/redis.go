package token

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type redisToken struct {
	db func() redis.Conn
}

func (w *redisToken) GetToken(appId, secret string) (token string, err error) {
	conn := w.db()
	defer conn.Close()

	valueKey := fmt.Sprintf("AccessToken:%s+%s", appId, secret)
	token, err = redis.String(conn.Do("GET", valueKey))
	if err != nil {
		token, err = GetAccessToken(appId, secret)
		if err != nil {
			return "", err
		}
		if _, err = conn.Do("SET", valueKey, token, "EX", 7000); err != nil {
			return "", err
		}
	}

	return token, nil
}

func NewRedisToken(db func() redis.Conn) *redisToken {
	return &redisToken{
		db: db,
	}
}
