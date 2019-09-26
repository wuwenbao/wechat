package work

type App struct {
	Confer
}

func New(c Confer) *App {
	if c == nil {
		c = Conf("", "", nil)
	}
	return &App{
		Confer: c,
	}
}

//AppChat 成员相关
func (a *App) User() *User {
	return NewUser(a.Confer)
}

//Message 消息相关
func (a *App) Message(agentId int) *Message {
	return NewMessage(a.Confer, agentId)
}

//Checkin 打卡相关
func (a *App) Checkin() *Checkin {
	return NewCheckin(a.Confer)
}

//AppChat 群聊相关
func (a *App) AppChat() *AppChat {
	return NewAppChat(a.Confer)
}
