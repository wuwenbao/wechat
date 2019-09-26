package appchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

type ParamUpdate struct {
	Chatid      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}

func Update(token string, param *ParamUpdate) error {
	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(param); err != nil {
		return err
	}
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/appchat/update?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return response.ReadBody(resp.Body, nil)
}
