package miniprogram

type App struct {
	Confer
}

func New(c Confer) *App {
	if c == nil {
		c = Conf("", "", nil)
	}
	app := &App{
		Confer: c,
	}
	return app
}

//Auth 开放接口
func (a *App) Auth() *Auth {
	return NewAuth(a)
}

//Wxacode 小程序二维码
func (a *App) Wxacode() *Wxacode {
	return NewWxacode(a)
}

//Encryptor 消息解密
func (a *App) Encryptor() *Encryptor {
	return NewEncryptor(a)
}

//UniformMessage 统一服务消息
func (a *App) UniformMessage() *UniFormMessage {
	return NewUniFormMessage(a)
}

//TemplateMessage 模版消息
func (a *App) TemplateMessage() *TemplateMessage {
	return NewTemplateMessage(a)
}

//CustomerServiceMessage 客服消息
func (a *App) CustomerServiceMessage() *CustomerServiceMessage {
	return NewCustomerServiceMessage(a)
}

//UpdatableMessage 动态消息
func (a *App) UpdatableMessage() {

}

//Soter 生物认证
func (a *App) Soter() {
}

//Security 内容安全
func (a *App) Security() {
}

//Logistics 物流助手
func (a *App) Logistics() {
}

//Analysis 数据分析
func (a *App) Analysis() {
}

//PluginManager 插件管理
func (a *App) PluginManager() {
}

//NearbyPoi 附近的小程序
func (a *App) NearbyPoi() {
}
