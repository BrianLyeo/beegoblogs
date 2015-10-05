package controllers

import (
	"time"
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
		//c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["Operation"] = "Register"
		c.Data["Reason"] = "Input Someting Invalid"
		c.TplNames = "common_failed.tpl"
		return
	}
	
	err = models.CreateNewAccount(&account)
	if err != nil {
		//c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["Operation"] = "Register"
		existAccount, _ := models.QueryAccountByEmail(account.EMail)
		if existAccount != nil {
			c.Data["Reason"] = "Your Email is Used by Others"
		} else {
			c.Data["Reason"] = "Something Error, Try Later"
		}
		c.TplNames = "common_failed.tpl"
		
		return
	}
	
	c.Ctx.ResponseWriter.Write([]byte("Success"))
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
	
	account.EMail = email
	account.Password = password
	account.DisplayName = name
	account.CreateTime = time.Now()
	account.LastLoginTime = time.Now()
	if showemail == "show" {
		account.ShowEmail = true
	} else {
		account.ShowEmail = false
	}
	account.Level = 0
	
	return err
}
