package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_docker/calender/conf"
	"go_docker/calender/infrastructure/router"
)

func init() {
	// 設定ファイル読み取り
	//read config file
	conf.Init()
}
func main() {
	router.Init()
	// defer database.CloseConn()
}
