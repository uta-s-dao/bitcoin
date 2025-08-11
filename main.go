package main

import (
	"fmt"
	"gotrading/app/controllers"
	"gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData() // ← データ取り込み処理
	controllers.StartWebServer()      // ← Webサーバー起動
}
