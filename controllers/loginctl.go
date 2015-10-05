package controllers

import (
	"beegoblogs/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplNames = "login.tpl"
}

func (c *LoginController) Post() {
	email := c.GetString("email")
	passwd := c.GetString("password")
	
	account, err := models.QueryAccountByEmail(email)
	if err != nil || account == nil || account.Password != passwd {
		c.Data["Timeout"] = "2"
		c.Data["URL"] = "/static/page/login_failed.html"
	} else {
		c.Data["Timeout"] = "2"
		c.Data["URL"] = "/login_success"
	}
	c.TplNames = "login_redirect.tpl"
}