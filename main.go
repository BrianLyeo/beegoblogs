package main

import (
	_ "beegoblogs/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}

