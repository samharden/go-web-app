package routers

import (
	"github.com/heroku/testapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/test", &controllers.TestController{})
		beego.SetStaticPath("/css","css")
}
