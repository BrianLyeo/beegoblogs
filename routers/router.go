package routers

import (
	"beegoblogs/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login_success", &controllers.LoginSuccessController{})
	beego.Router("/main", &controllers.MainPageController{})
}
