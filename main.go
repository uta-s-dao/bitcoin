package main

import (
	"fmt"
	"gotrading/bitflyer"
	"gotrading/config"
	"gotrading/utils"
	// "log"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
    fmt.Println(apiClient.GetBalance())

    // log.Printf("test")
    // fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
}