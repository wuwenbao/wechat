package work

import (
	"fmt"
	"strings"

	"github.com/wuwenbao/wechat/work/checkin"
)

type Checkin struct {
	Confer
}

//GetCheckinData 获取打卡数据
func (c *Checkin) GetCheckinData(dataType, startTime, endTime int, user ...string) (*checkin.DataResponse, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	format := `
	{
	   "opencheckindatatype": %d,
	   "starttime": %d,
	   "endtime": %d,
	   "useridlist": ["%s"]
	}`
	data := fmt.Sprintf(format, dataType, startTime, endTime, strings.Join(user, `", "`))
	return checkin.GetCheckinData(token, strings.NewReader(data))
}

//GetCheckinOption 获取打卡规则
func (c *Checkin) GetCheckinOption(dateTime int, user ...string) (*checkin.OptionResponse, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	format := `
	{
		"datetime": %d,
		"useridlist": ["%s"]
	}`
	data := fmt.Sprintf(format, dateTime, strings.Join(user, `", "`))
	return checkin.GetCheckinOption(token, strings.NewReader(data))
}

func NewCheckin(c Confer) *Checkin {
	return &Checkin{
		Confer: c,
	}
}
