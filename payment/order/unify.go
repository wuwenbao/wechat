package order

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/wuwenbao/wechat/payment"
	"github.com/wuwenbao/wechat/util"
)

type UnifyResponse struct {
	util.ResponseError
	Appid      string `xml:"appid"`       //公众账号ID
	MchId      string `xml:"mch_id"`      //商户号
	DeviceInfo string `xml:"device_info"` //设备号
	NonceStr   string `xml:"nonce_str"`   //随机字符串
	Sign       string `xml:"sign"`        //签名`
	Openid     string `xml:"openid"`      //用户标识
	PrepayId   string `xml:"prepay_id"`   //预支付交易会话标识
	TradeType  string `xml:"trade_type"`  //交易类型
	CodeUrl    string `xml:"code_url"`    //二维码链接
}

func Unify(body io.Reader) (*UnifyResponse, error) {
	resp, err := http.Post(`https://api.mch.weixin.qq.com/pay/unifiedorder`, "", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(UnifyResponse)
	if err := xml.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	if err := response.Check(); err != nil {
		return nil, err
	}
	return response, nil
}

type UnifyParam struct {
	Appid          string        `xml:"appid"`
	MchId          string        `xml:"mch_id"`
	DeviceInfo     string        `xml:"device_info,omitempty"`
	NonceStr       string        `xml:"nonce_str"`
	Sign           string        `xml:"sign"`
	SignType       string        `xml:"sign_type,omitempty"`
	Body           string        `xml:"body"`
	Detail         payment.CDATA `xml:"detail,omitempty"`
	Attach         string        `xml:"attach,omitempty"`
	OutTradeNo     string        `xml:"out_trade_no"`
	FeeType        string        `xml:"fee_type,omitempty"`
	TotalFee       int           `xml:"total_fee"`
	SpbillCreateIp string        `xml:"spbill_create_ip"`
	TimeStart      string        `xml:"time_start,omitempty"`
	TimeExpire     string        `xml:"time_expire,omitempty"`
	GoodsTag       string        `xml:"goods_tag,omitempty"`
	NotifyUrl      string        `xml:"notify_url"`
	TradeType      string        `xml:"trade_type"`
	ProductId      string        `xml:"product_id,omitempty"`
	LimitPay       string        `xml:"limit_pay,omitempty"`
	Openid         string        `xml:"openid,omitempty"`
	SceneInfo      payment.CDATA `xml:"scene_info,omitempty"`
}

func (u *UnifyParam) SetDetail(str string) *UnifyParam {
	u.Detail = payment.CDATA{Text: str}
	return u
}

func (u *UnifyParam) SetSceneInfo(str string) *UnifyParam {
	u.SceneInfo = payment.CDATA{Text: str}
	return u
}

//JSAPI 微信网页端
func JSAPI() *UnifyParam {
	order := &UnifyParam{}
	order.TradeType = "JSAPI"
	return order
}
