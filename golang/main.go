package main

import (
	"fmt"
	"golang/calendar/infrastructure/router"
	"golang/conf"
)

func init() {
	// 設定ファイル読み取り
	//read config file
	conf.Init()
}
func main() {
	fmt.Println("hello world")
	router.Init()
	// defer database.CloseConn()
}
