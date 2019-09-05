package plugin_manager

import (
	"github.com/wuwenbao/wechat/miniprogram"
)

type PluginManager struct {
	miniprogram.Confer
}

//New 插件管理
func New(c miniprogram.Confer) *PluginManager {
	return &PluginManager{Confer: c}
}
