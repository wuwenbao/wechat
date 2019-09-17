package transfer

import (
	"crypto/tls"
	"encoding/xml"
	"io"
	"net/http"

	response2 "github.com/wuwenbao/wechat/internal/response"
	"github.com/wuwenbao/wechat/util"
)

type BalanceParam struct {
	Appid          string `xml:"appid"`
	MchId          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	Openid         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name"`
	Amount         int    `xml:"amount"`
	Desc           string `xml:"desc"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
}

//NewBalanceParam 企业付款到用户零钱参数
func NewBalanceParam(partnerTradeNo, openid, desc string, amount int) *BalanceParam {
	param := &BalanceParam{
		PartnerTradeNo: partnerTradeNo,
		Openid:         openid,
		Amount:         amount,
		Desc:           desc,
		SpbillCreateIp: "127.0.0.1",
		CheckName:      "NO_CHECK",
	}
	return param
}

//Balance 企业付款到用户零钱
func Balance(tlsConfig *tls.Config, body io.Reader) error {
	request, err := http.NewRequest(http.MethodPost, `https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers`, body)
	if err != nil {
		return err
	}
	resp, err := util.ClientTLS(request, tlsConfig)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response := new(response2.ResponseError)
	err = xml.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}
	return response.Check()
}
