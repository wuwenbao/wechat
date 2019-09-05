package payment

import (
	"bytes"
	"encoding/xml"

	"github.com/wuwenbao/wechat/payment/transfer"
	"github.com/wuwenbao/wechat/util"
)

type Transfer struct {
	Confer
}

//Balance 企业付款到用户零钱
func (t *Transfer) Balance(r *transfer.BalanceParam) error {
	r.Appid = t.Appid()
	r.MchId = t.MchId()
	//数据签名
	r.NonceStr = util.RandomStr(10)
	r.Sign = t.SignCheck(r)

	bts, err := xml.Marshal(struct {
		XMLName xml.Name `xml:"xml"`
		*transfer.BalanceParam
	}{BalanceParam: r})
	if err != nil {
		return err
	}
	//fmt.Println(string(bts))
	tlsConfig, err := t.GetTLSConfig()
	if err != nil {
		return err
	}
	return transfer.Balance(tlsConfig, bytes.NewReader(bts))
}


//NewTransfer 企业转账
func NewTransfer(c Confer) *Transfer  {
	return &Transfer{Confer: c}
}