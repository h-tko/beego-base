package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/h-tko/beego-base/helpers"
	_ "github.com/h-tko/beego-base/routers"
	_ "github.com/lib/pq"
)

// TemplateHelperの登録処理
func initTemplate() {
	helpers.RegisterTemplateHelpers()
}

// DB接続初期化
func initDB() {
	dbLink := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		beego.AppConfig.String("db_user"),
		beego.AppConfig.String("db_password"),
		beego.AppConfig.String("db_name"),
		beego.AppConfig.String("db_host"),
		beego.AppConfig.String("db_port"),
	)

	logs.Info("%s", dbLink)

	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		panic(err)
	}

	if err := orm.RegisterDataBase("default", "postgres", dbLink); err != nil {
		panic(err)
	}
}

func init() {
	initDB()
}

// アプリケーションエントリーポイント
func main() {
	if beego.AppConfig.String("runmode") != "prod" {
		orm.Debug = true
	}

	initTemplate()

	// ログ設定
	logs.SetLogger(logs.AdapterFile, `{"filename": "logs/app.log"}`)

	beego.Run()
}
