package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplNames = "register.tpl"
}

func (c *RegisterController) Post() {
	passwd := c.GetString("password")
	if passwd == "" || passwd[0] != 'p' {
		c.Data["Timeout"] = "2"
		c.Data["URL"] = "/static/page/login_failed.html"
	} else {
		c.Data["Timeout"] = "2"
		c.Data["URL"] = "/login_success"
	}
	c.TplNames = "login_redirect.tpl"
}
