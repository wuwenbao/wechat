package uniform_message

import (
	"fmt"
	"io"

	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

const (
	SendApi = `https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s`
)

//Send 下发小程序和公众号统一的服务消息
func Send(token string, body io.Reader) error {
	resp, err := http.Post(fmt.Sprintf(SendApi, token), "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return response.ReadBody(resp.Body, nil)
}
