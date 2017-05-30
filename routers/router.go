package routers

import (
	"github.com/h-tko/beego-base/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
