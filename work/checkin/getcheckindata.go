package checkin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/wuwenbao/wechat/util"
)

type DataResponse struct {
	Checkindata []struct {
		Userid         string   `json:"userid"`
		Groupname      string   `json:"groupname"`
		CheckinType    string   `json:"checkin_type"`
		ExceptionType  string   `json:"exception_type"`
		CheckinTime    int      `json:"checkin_time"`
		LocationTitle  string   `json:"location_title"`
		LocationDetail string   `json:"location_detail"`
		Wifiname       string   `json:"wifiname"`
		Notes          string   `json:"notes"`
		Wifimac        string   `json:"wifimac"`
		Mediaids       []string `json:"mediaids"`
	} `json:"checkindata"`
}

//GetCheckinData 获取打卡数据
func GetCheckinData(token string, body io.Reader) (*DataResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(DataResponse)
	if err := util.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}
