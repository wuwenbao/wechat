package encryptor

import "errors"

//水印
type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

func (w *Watermark) Check(appId string) error {
	if w.Appid != appId {
		return errors.New("数据不合法")
	}
	return nil
}
