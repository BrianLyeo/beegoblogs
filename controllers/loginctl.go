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
	c.TplNames = "welcome.tpl"
}