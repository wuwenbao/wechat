package order

import (
	"crypto/tls"
	"encoding/xml"
	"io"
	"net/http"

	response2 "github.com/wuwenbao/wechat/internal/response"
	"github.com/wuwenbao/wechat/util"
)

type RefundParam struct {
	Appid         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      int    `xml:"total_fee"`
	RefundFee     int    `xml:"refund_fee"`
	RefundFeeType string `xml:"refund_fee_type,omitempty"`
	RefundDesc    string `xml:"refund_desc,omitempty"`
	RefundAccount string `xml:"refund_account,omitempty"`
	NotifyUrl     string `xml:"notify_url,omitempty"`
}

//Refund 退款
func Refund(tlsConfig *tls.Config, body io.Reader) error {
	request, err := http.NewRequest(http.MethodPost, `https://api.mch.weixin.qq.com/secapi/pay/refund`, body)
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

//ByOutTradeNumber 根据商户订单号退款
func ByOutTradeNumber(transactionId, outRefundNo string, totalFee, refundFee int) *RefundParam {
	param := &RefundParam{
		TransactionId: transactionId,
		OutRefundNo:   outRefundNo,
		TotalFee:      totalFee,
		RefundFee:     refundFee,
	}
	return param
}

//ByTransactionId 根据微信订单号退款
func ByTransactionId(outTradeNo, outRefundNo string, totalFee, refundFee int) *RefundParam {
	param := &RefundParam{
		OutTradeNo:  outTradeNo,
		OutRefundNo: outRefundNo,
		TotalFee:    totalFee,
		RefundFee:   refundFee,
	}
	return param
}
