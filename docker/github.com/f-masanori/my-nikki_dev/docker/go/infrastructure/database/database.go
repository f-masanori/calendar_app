package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/f-masanori/my-nikki_dev/docker/go/conf"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var ConnectionError error

type SqlHandler struct {
	DB *sql.DB
}

func NewSqlHandler() *SqlHandler {

	//configからDBの読み取り
	connectionCmd := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.C.Database.User,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Dbname,
	)
	// // DB の接続
	// 	user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	// 	"root:mysql@tcp(mysql_container:3306)/app"

	Db, ConnectionError = sql.Open(conf.C.Database.Drivername, connectionCmd)
	fmt.Println("DB-information")
	fmt.Println("drivername: " + conf.C.Database.Drivername)
	fmt.Println("host: " + conf.C.Database.Host)
	fmt.Println("port: " + conf.C.Database.Port)
	fmt.Println("dbname: " + conf.C.Database.Dbname)
	fmt.Println("______________")
	if ConnectionError != nil {
		log.Fatal("error connecting to database: ", ConnectionError)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.DB = Db

	return sqlHandler
}

func Connect() *sql.DB {

	//configからDBの読み取り
	connectionCmd := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.C.Database.User,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Dbname,
	)

	Db, ConnectionError = sql.Open(conf.C.Database.Drivername, connectionCmd)
	fmt.Println("DB-information")
	fmt.Println("drivername: " + conf.C.Database.Drivername)
	fmt.Println("host: " + conf.C.Database.Host)
	fmt.Println("port: " + conf.C.Database.Port)
	fmt.Println("dbname: " + conf.C.Database.Dbname)
	fmt.Println("______________")
	if ConnectionError != nil {
		log.Fatal("error connecting to database: ", ConnectionError)
	}
	return Db
}

func CloseConn() {
	Db.Close()
}
