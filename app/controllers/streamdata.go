package controllers

import (
	"gotrading/app/models"
	"gotrading/bitflyer"
	"gotrading/config"
	"log"
)

func StreamIngestionData() {
	c := config.Config
	ai := NewAI(c.ProductCode, c.TradeDuration, c.DataLimit, c.UsePercent, c.StopLimitPercent, c.BackTest)
	var tickerChannl = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	//取得されたデータがtickerChannlにわたってくる
	go func() {
		for ticker := range tickerChannl {
			//tickerChannlからTickerデータを継続的に受信
			//tickerChannl <- data1
			//tickerChannl <- data2
			//tickerChannl <- data3
			log.Printf("action=StreamIngestionData, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == config.Config.TradeDuration {
					// ← 新しいローソク足完成時のみ取引実行
					//10:30:15   既存ローソク足更新     false     ×
					// 10:30:32   既存ローソク足更新    false     ×
					// 10:30:58   既存ローソク足更新    false     ×
					// 10:31:00   新しいローソク足作成   true     ○ ← ai.Trade()実行！
					ai.Trade()
				}
			}
		}
	}()
}
