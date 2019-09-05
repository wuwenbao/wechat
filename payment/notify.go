package payment

import (
	"net/http"

	"github.com/wuwenbao/wechat/payment/notify"
	"github.com/wuwenbao/wechat/util"
)

type Notify struct {
	Confer
}

//Refund 退款通知
func (n *Notify) Refund(w http.ResponseWriter, r *http.Request, notifyHandle func(info *notify.RefundReqInfo) error) {
	defer r.Body.Close()
	info, err := notify.RefundSign(n.MchKey(), r.Body)
	if err != nil {
		w.Write([]byte(util.NotifyFail(err)))
		return
	}
	if err := notifyHandle(info); err != nil {
		w.Write([]byte(util.NotifyFail(err)))
		return
	}
	w.Write([]byte(util.NotifySuccess()))
}

//Paid 支付通知
func (n *Notify) Paid(w http.ResponseWriter, r *http.Request, notifyHandle func(res *notify.PaidResponse) error) {
	defer r.Body.Close()
	res, err := notify.PaidSign(n.SignCheck, r.Body)
	if err != nil {
		w.Write([]byte(util.NotifyFail(err)))
		return
	}
	if err := notifyHandle(res); err != nil {
		w.Write([]byte(util.NotifyFail(err)))
		return
	}
	w.Write([]byte(util.NotifySuccess()))
}

//NewNotify 通知相关
func NewNotify(c Confer) *Notify {
	return &Notify{Confer: c}
}
