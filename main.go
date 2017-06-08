package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/h-tko/beego-base/helpers"
	_ "github.com/h-tko/beego-base/routers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

// TemplateHelperの登録処理
func initTemplate() {
	helpers.RegisterTemplateHelpers()
}

// .env読み込み
func envLoad() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
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

// .envの内容に設定を書き換える
func updateSettings() {
	beego.BConfig.RunMode = os.Getenv("APP_RUNMODE")
	beego.BConfig.Listen.HTTPAddr = os.Getenv("APP_DEVHOST_NAME")
}

// アプリケーションエントリーポイント
func main() {
	err := envLoad()

	if err != nil {
		panic(err)
	}

	updateSettings()

	// 初期化処理
	initDB()
	initTemplate()

	if beego.BConfig.RunMode != "prod" {
		orm.Debug = true
	}

	// ログ設定
	logs.SetLogger(logs.AdapterFile, `{"filename": "logs/app.log"}`)

	beego.Run()
}
