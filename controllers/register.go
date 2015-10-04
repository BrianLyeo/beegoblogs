package controllers

import (
	"log"
	"errors"
	"beegoblogs/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplNames = "register.tpl"
}

func (c *RegisterController) Post() {
	
	var account models.Account
	err := parseAccount(&c.Controller, &account)
	if err != nil {
		c.Redirect("/login_failed2.html", 302)
		log.Fatal(err)
		return
	}
	
	err = models.CreateNewAccount(&account)
	if err != nil {
		c.Redirect("/login_failed3.html", 302)
		log.Fatal(err)
		return
	}
}

func parseAccount(ctx *beego.Controller, account *models.Account) error {
	email := ctx.GetString("email")
	password := ctx.GetString("password")
	password2 := ctx.GetString("password2")
	name := ctx.GetString("name")
	showemail := ctx.GetString("showemail")
	
	var err error
	err = nil
	if email == "" {
		err = errors.New("email miss")
		return err
	}
	
	if password == "" || password != password2 {
		err = errors.New("password")
		return err
	}
	
	if name == "" {
		err = errors.New("name miss")
		return err
	}
	
	if showemail != "show" && showemail != "hide" {
		err = errors.New("showemail miss")
		return err
	}
	
	return err
}
