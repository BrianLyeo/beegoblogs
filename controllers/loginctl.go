package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplNames = "login.tpl"
}

func (c *LoginController) Post() {
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