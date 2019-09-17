package response

import (
	"errors"
	"fmt"
)

type ReturnError struct {
	ReturnCode string `xml:"return_code"` //返回状态码
	ReturnMsg  string `xml:"return_msg"`  //返回信息
}

func (r *ReturnError) Check() error {
	if r.ReturnCode == "FAIL" {
		return errors.New(r.ReturnMsg)
	}
	return nil
}

type ResultError struct {
	ResultCode string `xml:"result_code"`  //业务结果
	ErrCode    string `xml:"err_code"`     //错误代码
	ErrCodeDes string `xml:"err_code_des"` //错误代码描述
}

func (r *ResultError) Check() error {
	if r.ResultCode == "FAIL" {
		return errors.New(r.ErrCode)
	}
	return nil
}

type ResponseError struct {
	ReturnError
	ResultError
}

func (r *ResponseError) Check() error {
	if err := r.ReturnError.Check(); err != nil {
		return err
	}
	if err := r.ResultError.Check(); err != nil {
		return err
	}
	return nil
}

//Fail 通知失败
func NotifyFail(msg error) string {
	f := `<xml><return_code><![CDATA[FAIL]]></return_code>return_msg><![CDATA[%s]]></return_msg></xml>`
	return fmt.Sprintf(f, msg)
}

//Success 通知成功
func NotifySuccess() string {
	msg := `<xml><return_code><![CDATA[SUCCESS]]></return_code>return_msg><![CDATA[OK]]></return_msg></xml>`
	return msg
}
