package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Layout = "base.html"
	c.TplName = "index.tpl"
}
func (c *TestController) Get() {
	c.Data["youthere"] = "Worlds"
	c.Layout = "base.html"
	c.TplName = "test.tpl"

}
