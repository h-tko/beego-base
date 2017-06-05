package controllers

import (
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

func (c *TopController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
