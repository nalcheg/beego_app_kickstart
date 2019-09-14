package routers

import (
	"github.com/astaxie/beego"
	"github.com/nalcheg/beego_app_kickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/page-one", &controllers.PageOneController{})
	beego.Router("/page-two", &controllers.PageTwoController{})
}
