package routers

import (
	"github.com/astaxie/beego"
	"github.com/h-tko/beego-base/controllers"
)

func init() {
	beego.Router("/", &controllers.TopController{})
}
