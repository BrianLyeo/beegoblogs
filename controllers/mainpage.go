package controllers

import (
	"beegoblogs/models"
	"github.com/astaxie/beego"
)

type MainPageController struct {
	beego.Controller
}

func (c *MainPageController) Get() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	
	c.Data["BlogsList"] = models.GetAllBlogEntries()
	c.Data["User"] = sess.Get("user-email")
	c.TplNames = "mainpage.tpl"
}