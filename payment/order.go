package payment

import (
	"bytes"
	"encoding/xml"

	"github.com/wuwenbao/wechat/payment/order"
	"github.com/wuwenbao/wechat/util"
)

type Order struct {
	Confer
}

//Unify 统一下单
func (o *Order) Unify(r *order.UnifyParam) (*order.UnifyResponse, error) {
	r.Appid = o.Appid()
	r.MchId = o.MchId()
	//数据签名
	r.NonceStr = util.RandomStr(10)
	r.Sign = o.SignCheck(r)

	bts, err := xml.Marshal(struct {
		XMLName xml.Name `xml:"xml"`
		*order.UnifyParam
	}{UnifyParam: r})
	if err != nil {
		return nil, err
	}
	return order.Unify(bytes.NewReader(bts))
}

//Refund 退款
func (o *Order) Refund(r *order.RefundParam) error {
	r.Appid = o.Appid()
	r.MchId = o.MchId()
	//数据签名
	r.NonceStr = util.RandomStr(10)
	r.Sign = o.SignCheck(r)

	bts, err := xml.Marshal(struct {
		XMLName xml.Name `xml:"xml"`
		*order.RefundParam
	}{RefundParam: r})
	if err != nil {
		return err
	}
	tlsConfig, err := o.GetTLSConfig()
	if err != nil {
		return err
	}
	return order.Refund(tlsConfig, bytes.NewReader(bts))
}

func NewOrder(c Confer) *Order {
	return &Order{Confer: c}
}
