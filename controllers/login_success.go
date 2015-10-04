package controllers

import (
	"github.com/astaxie/beego"
)

type LoginSuccessController struct {
	beego.Controller
}

func (c *LoginSuccessController)Get() {
	c.TplNames = "welcome.tpl"
}