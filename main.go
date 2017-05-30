package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/h-tko/beego-base/routers"
)

func main() {
	// ログ設定
	logs.SetLogger(logs.AdapterFile, `{"filename": "logs/app.log"}`)

	beego.Run()
}
