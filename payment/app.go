package payment

type App struct {
	Confer
}

func New(c Confer) *App {
	if c == nil {
		c = Conf("", "", "")
	}
	return &App{
		Confer: c,
	}
}

//Refund 退款
func (a *App) Notify() *Notify {
	return NewNotify(a)
}

//Order 下单
func (a *App) Order() *Order {
	return NewOrder(a)
}

//Transfer 企业转账
func (a *App) Transfer() *Transfer {
	return NewTransfer(a)
}
