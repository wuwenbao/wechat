package appchat

import (
	"fmt"
	"io"
	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

type Sender interface {
	ChatSend(chatid string) string
}

//Send 发送应用消息
func Send(token string, body io.Reader) error {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return response.ReadBody(resp.Body, nil)
}
