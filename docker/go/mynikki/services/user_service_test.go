package services

import (
	"fmt" 
	"testing"
	"go_docker/mynikki/conf"
	// "go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)
func TestGetAllSuccess(t *testing.T) {
	/* テストのためのconfig実体作成 */
	conf.Test()

	/* テスト用データベースに接続(configでテストDBを決める) */
	DBhandler := database.TestNewSqlHandler()

	NewUserService := NewUserService(DBhandler)
	
	users, err := NewUserService.GetAll()

	expected := "entities.Users"
	autual := reflect.TypeOf(users).String()
	
	if autual != expected{
		t.Fatalf("failed test %#v", "返り値型エラー")
	}
	if err !=nil{
		t.Fatalf("failed test %#v", err)
	}
}
// func TestStoreNewUserSuccess(t *testing.T) {
// 	conf.Init()
// 	DBhandler := database.NewSqlHandler()
// 	rows, err := DBhandler.DB.Query("SELECT * from users;")
// 	fmt.Println(rows)
// 	if err !=nil{
// 		t.Fatalf("failed test %#v", err)
// 	}
// }
