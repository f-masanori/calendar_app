package main

import (
	"github.com/f-masanori/my-nikki_dev/docker/go/conf"
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/router"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 設定ファイル読み取り
	conf.Init()
}
func main() {
	router.Init()
	// defer database.CloseConn()
}
