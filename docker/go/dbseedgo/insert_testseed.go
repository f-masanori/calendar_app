package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/f-masanori/my-nikki_dev/docker/go/conf"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/* conf.Initで設定構造体の実体を作成 */
	conf.Init()

	DBhandler :=  NewSqlHandler()

	Create_seed_data_to_users(DBhandler)

}

func Create_seed_data_to_users(DB *sql.DB){
	statement := "INSERT INTO users(name) VALUES(?)"
	stmtInsert ,err := DB.Prepare(statement)
	if err != nil{
		fmt.Println("error1")
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec("testname")
	if err != nil{
		fmt.Println("error2")
	}
	lastInsertID, err := result.LastInsertId()
	fmt.Println(lastInsertID)
}

func NewSqlHandler() *sql.DB {

	//configからDBの読み取り
	connectionCmd := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Dbname,
	)
	// // DB の接続
	// 	user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	// 	"root:mysql@tcp(mysql_container:3306)/app"
	var Db *sql.DB
	var ConnectionError error

	Db, ConnectionError = sql.Open(conf.Database.Drivername, connectionCmd)
	fmt.Println("DB-information")
	fmt.Println("drivername: " + conf.Database.Drivername)
	fmt.Println("host: " + conf.Database.Host)
	fmt.Println("port: " + conf.Database.Port)
	fmt.Println("dbname: " + conf.Database.Dbname)
	fmt.Println("______________")
	if ConnectionError != nil {
		log.Fatal("error connecting to database: ", ConnectionError)
	}
	return Db
}
