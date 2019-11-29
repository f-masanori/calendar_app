package services

import (
	"fmt" 
	"testing"
	"go_docker/mynikki/conf"
	"go_docker/mynikki/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)
func TestGetAllSuccess(t *testing.T) {
	/* テストのためのconfig実体作成 */
	conf.Test()

	DBhandler := database.TestNewSqlHandler()

	NewUserService := NewUserService(DBhandler)
	users, err := NewUserService.GetAll()
	fmt.Println("test")
	fmt.Println(users)
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
