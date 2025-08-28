package main

import (
	"gotrading/app/controllers"
	"gotrading/config" //ここで init() が実行される
	"gotrading/utils"
	"log"
)

func main() {
	// 　グローバルロガーの出力先が「コンソール+ファイル」に設定される
	// 　すべてのパッケージの log.Printf() が同じ設定を使用
	utils.LoggingSettings(config.Config.LogFile)

	controllers.StreamIngestionData()

	log.Println(controllers.StartWebServer()) //エラー時のみログ出力
}
