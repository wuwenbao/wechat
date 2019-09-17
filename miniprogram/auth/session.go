package auth

import (
	"fmt"
	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

type Session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

//Code2Session 登录凭证校验
func Code2Session(appId, secret, code string) (*Session, error) {
	apiUrl := fmt.Sprintf(`https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`, appId, secret, code)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(Session)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}
