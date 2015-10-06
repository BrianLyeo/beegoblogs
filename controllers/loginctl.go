package controllers

import (
	"github.com/astaxie/beego/session"
	"beegoblogs/models"
	"github.com/astaxie/beego"
)

var globalSessions *session.Manager

func init() {
    globalSessions, _ = session.NewManager("memory",
		`{
			"cookieName":"gosessionid",
			"enableSetCookie,omitempty": true,
			"gclifetime":3600,
			"maxLifetime": 3600,
			"secure": false,
			"sessionIDHashFunc": "sha1",
			"sessionIDHashKey": "",
			"cookieLifeTime": 3600,
			"providerConfig": ""
		}`)
    go globalSessions.GC()
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	
	action := c.GetString("action")
	if action == "logout" {
		sess.Delete("user-id")
		sess.Delete("user-email")
		sess.Delete("user-name")
		c.Redirect("/main", 302)
	} else {
		if sess.Get("user-email") != nil {
			c.TplNames = "welcome.tpl"
		} else {
			c.Layout = "layout/layout_blog.html"
		    c.TplNames = "login.tpl"
		    c.LayoutSections = make(map[string]string)
		    c.LayoutSections["HtmlHead"] = "layout/html_head.html"
			c.LayoutSections["TopNav"] = "layout/top_nav.html"
		    c.LayoutSections["Scripts"] = "layout/scripts.html"
		    c.LayoutSections["Sidebar"] = ""
			c.TplNames = "login.tpl"
		}
	}
	
}

func (c *LoginController) Post() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	
	email := c.GetString("email")
	passwd := c.GetString("password")
	
	account, err := models.QueryAccountByEmail(email)
	if err != nil || account == nil || account.Password != passwd {
		c.Data["Timeout"] = "2"
		c.Data["URL"] = "/static/page/login_failed.html"
	} else {
		c.Data["Timeout"] = "2"
		c.Data["URL"] = "/login_success"
		sess.Set("user-id", account.Id)
		sess.Set("user-email", account.EMail)
		sess.Set("user-name", account.DisplayName)
	}
	c.TplNames = "login_redirect.tpl"
}