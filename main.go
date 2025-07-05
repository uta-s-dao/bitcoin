package main

import (
	"fmt"
	"time"
	"gotrading/bitflyer"
	"gotrading/config"
	"gotrading/utils"
	// "log"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_USD")
    fmt.Println(ticker)
    fmt.Println(ticker.GetMidPrice())
    fmt.Println(ticker.DateTime())
    fmt.Println(ticker.TruncateDateTime(time.Hour)) 

    // log.Printf("test")
    // fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
}