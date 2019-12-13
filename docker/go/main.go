package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_docker/mynikki/conf"
	"go_docker/mynikki/infrastructure/router"
)

func init() {
	// 設定ファイル読み取り
	conf.Init()
}
func main() {
	router.Init()
	// defer database.CloseConn()
}
