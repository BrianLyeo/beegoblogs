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
	
	c.Layout = "layout/layout_blog.html"
    c.TplNames = "mainpage.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "layout/html_head.html"
	c.LayoutSections["TopNav"] = "layout/top_nav.html"
    c.LayoutSections["Scripts"] = "layout/scripts.html"
    c.LayoutSections["Sidebar"] = ""
	
	c.Data["BlogsList"] = models.GetAllBlogEntries()
	c.Data["User"] = sess.Get("user-email")
}