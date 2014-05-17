package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"myapp/app/routes"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) Signin(username, pwd string) revel.Result {
	user,has := models.Signin(username, pwd)
	if !has {
		c.Flash.Error("密码或用户名错误！")
		return c.Redirect(routes.Admin.Index())
	}
	c.Session["username"] = user.UserName
	//c.Session["id"] = user.UserId
	return c.Redirect(routes.Admin.Index())
}

/*func (c Admin) AddUser(username,input string) revel.Result {
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.Flash.Error("Form invalid. Try again.")
		return c.Redirect(routes.App.ViewForm(username))  // <--- 反转路由
	}
	c.Flash.Success("Form processed!")
	return c.Redirect(routes.Admin.ViewConfirmation(username, input)) 
}*/