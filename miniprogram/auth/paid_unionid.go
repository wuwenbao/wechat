package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/wuwenbao/wechat/util"
)

//GetPaidUnionId 用户支付完成后，获取该用户的unionid
func GetPaidUnionId(token, openid string, args ...string) (string, error) {
	var apiUrl string
	switch len(args) {
	case 1:
		apiUrl = fmt.Sprintf(`https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s&transaction_id=%s`, token, openid, args[0])
	case 2:
		apiUrl = fmt.Sprintf(`https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s&mch_id=%s&out_trade_no=%s`, token, openid, args[0], args[1])
	default:
		return "", errors.New("required: transaction_id / mch_id+out_trade_no")
	}
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := new(struct {
		Unionid string `json:"unionid"`
	})
	if err := util.ReadBody(resp.Body, data); err != nil {
		return "", err
	}
	return data.Unionid, nil
}
