package nearby_poi

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type NearbyPoi struct {
	miniprogram.Confer
}

//New 附近的小程序
func New(c miniprogram.Confer) *NearbyPoi {
	return &NearbyPoi{Confer: c}
}
