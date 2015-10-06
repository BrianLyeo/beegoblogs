package controllers

import (
	//"beegoblogs/models"
	"github.com/astaxie/beego"
)

type BlogPageController struct {
	beego.Controller
}

func (c *BlogPageController) Get() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	
	c.Layout = "layout/layout_blog.html"
    c.TplNames = "layout/index.html"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "layout/html_head.html"
	c.LayoutSections["TopNav"] = "layout/top_nav.html"
    c.LayoutSections["Scripts"] = "layout/scripts.html"
    c.LayoutSections["Sidebar"] = ""
}