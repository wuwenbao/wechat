package token

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wuwenbao/wechat/internal/response"
)

type Adapter interface {
	GetToken(appId, secret string) (token string, err error)
}

//GetAccessToken 接口调用凭证
func GetAccessToken(appId, secret string) (string, error) {
	var apiUrl string
	if strings.Contains(appId, "wx") {
		apiUrl = fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, appId, secret)
	} else {
		apiUrl = fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s`, appId, secret)
	}
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := new(struct {
		AccessToken string `json:"access_token"`
	})
	if err := response.ReadBody(resp.Body, data); err != nil {
		return "", err
	}
	return data.AccessToken, nil
}
