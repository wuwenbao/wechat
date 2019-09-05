package template_message

import (
	"fmt"
	"io"

	"github.com/wuwenbao/wechat/util"

	"net/http"
)

const (
	SendApi = `https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s`
)

//Send 发送消息
func Send(token string, body io.Reader) error {
	resp, err := http.Post(fmt.Sprintf(SendApi, token), "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return util.ReadBody(resp.Body, nil)
}

