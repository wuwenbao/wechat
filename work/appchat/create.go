package appchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

type ParamCreate struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid`
}

type ResponseCreate struct {
	Chatid string `json:"chatid"`
}

func Create(token string, param *ParamCreate) (*ResponseCreate, error) {
	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(param); err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", buffer)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := new(ResponseCreate)
	if err := response.ReadBody(resp.Body, info); err != nil {
		return nil, err
	}
	return info, nil
}
