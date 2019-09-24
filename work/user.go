package work

import "github.com/wuwenbao/wechat/work/user"

type User struct {
	Confer
}

//Get 读取成员
func (u *User) Get(userId string) (*user.User, error) {
	token, err := u.Token()
	if err != nil {
		return nil, err
	}
	return user.Get(token, userId)
}

//GetUserinfo 获取访问用户身份
func (u *User) GetUserinfo(code string) (*user.Userinfo, error) {
	token, err := u.Token()
	if err != nil {
		return nil, err
	}
	return user.GetUserinfo(token, code)
}

//SimpleList 获取访问用户身份
func (u *User) SimpleList(departmentId, fetchChild int) (*user.SimpleListResponse, error) {
	token, err := u.Token()
	if err != nil {
		return nil, err
	}
	return user.SimpleList(token, departmentId, fetchChild)
}

func NewUser(c Confer) *User {
	return &User{
		Confer: c,
	}
}
