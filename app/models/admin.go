package models

import (
	"fmt"
)

type AdminUser struct {
	UserId int16
	UserName string
	Email string
	Password string
	EcSolt string
	AddTime int64
	LastLogin int64
	LastIp string
	ActionList string
	NavList string
	LangType string
}

func Signin(username, pwd string) (*AdminUser, bool) {
	user := new(AdminUser)
	has, err := engine.Cols("user_id","user_name").Where("user_name=? and password=?", username, pwd).Get(user)
	if err != nil {
		//c.Flash.Error("非法查询")
		//return c.Redirect(routes.Admin.Index())
	}
	fmt.Println(user,has)
	return user,has
}