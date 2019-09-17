package customer_service_message

import (
	"fmt"
	"io"
	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

const (
	SendApi = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s`
)

//Send 发送消息
func Send(token string, body io.Reader) error {
	resp, err := http.Post(fmt.Sprintf(SendApi, token), "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return response.ReadBody(resp.Body, nil)
}
