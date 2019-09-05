package miniprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/wuwenbao/wechat/miniprogram/encryptor"
)

type Encryptor struct {
	Confer
}

//NewEncryptor 消息解密
func NewEncryptor(c Confer) *Encryptor {
	return &Encryptor{
		Confer: c,
	}
}

//GetPhone 获取用户手机号
func (e *Encryptor) GetPhoneNumber(session, iv, encryptData string) (*encryptor.PhoneNumber, error) {
	bts, err := e.DecryptData(session, iv, encryptData)
	if err != nil {
		return nil, err
	}
	return encryptor.GetPhoneNumber(e.Appid(), bts)
}

//GetUserInfo 获取用户信息
func (e *Encryptor) GetUserInfo(session, iv, encryptData string) (*encryptor.UserInfo, error) {
	bts, err := e.DecryptData(session, iv, encryptData)
	if err != nil {
		return nil, err
	}
	return encryptor.GetUserInfo(e.Appid(), bts)
}

//DecryptData 消息解密
func (e *Encryptor) DecryptData(session, iv, encryptData string) ([]byte, error) {
	return decryptData(encryptData, iv, session)
}

func decryptData(encryptData, iv, key string) ([]byte, error) {
	if len(key) != 24 {
		return nil, fmt.Errorf("aesKey非法")
	}
	if len(iv) != 24 {
		return nil, fmt.Errorf("aesIv非法")
	}
	data, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, fmt.Errorf("encryptData decode error:%s", err)
	}
	keyBts, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("key decode error:%s", err)
	}
	ivBts, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, fmt.Errorf("iv decode error:%s", err)
	}
	dataLen := len(data)
	block, err := aes.NewCipher(keyBts)
	if err != nil {
		return nil, fmt.Errorf("new cipher error:%s", err)
	}
	blockMode := cipher.NewCBCDecrypter(block, ivBts)
	origData := make([]byte, dataLen)
	blockMode.CryptBlocks(origData, data)
	tint := int(origData[dataLen-1])
	if dataLen-tint > 0 {
		return origData[:(dataLen - tint)], nil
	}
	return origData, nil
}

type Checker interface {
	Check(string) error
}
