package miniprogram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/wuwenbao/wechat/internal/response"
)

type Wxacode struct {
	Confer
}

//NewWxacode 小程序码
func NewWxacode(c Confer) *Wxacode {
	return &Wxacode{
		Confer: c,
	}
}

const (
	GetApi          = `https://api.weixin.qq.com/wxa/getwxacode?access_token=%s`
	GetUnlimitedApi = `https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s`
	CreateQRCodeApi = `https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s`
)

func (w *Wxacode) Get(body io.Reader) (io.Reader, error) {
	return w.create(GetApi, body)
}

func (w *Wxacode) GetUnlimited(body io.Reader) (io.Reader, error) {
	return w.create(GetUnlimitedApi, body)
}

func (w *Wxacode) CreateQRCode(body io.Reader) (io.Reader, error) {
	return w.create(CreateQRCodeApi, body)
}

func (w *Wxacode) create(api string, body io.Reader) (io.Reader, error) {
	accessToken, err := w.Token()
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf(api, accessToken), "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := new(struct {
		ContentType string          `json:"contentType"`
		Buffer      json.RawMessage `json:"buffer"`
	})
	if err = response.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	return bytes.NewReader(data.Buffer), nil
}
