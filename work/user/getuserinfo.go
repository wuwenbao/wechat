package user

import (
	"fmt"
	"net/http"

	"github.com/wuwenbao/wechat/util"
)

type Userinfo struct {
	UserId   string `json:"UserId"`
	DeviceId string `json:"DeviceId"`
	OpenId   string `json:"OpenId"`
}

func GetUserinfo(token, code string) (*Userinfo, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s`, token, code)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := new(Userinfo)
	if err := util.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	return data, nil
}
