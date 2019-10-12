package appchat

import (
	"fmt"
	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

type ResponseGet struct {
	ChatInfo *ChatInfo `json:"chat_info"`
}
type ChatInfo struct {
	Chatid   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
}

func Get(token, chatid string) (*ResponseGet, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=%s&chatid=%s`, token, chatid)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := new(ResponseGet)
	if err := response.ReadBody(resp.Body, info); err != nil {
		return nil, err
	}
	return info, nil
}
