package database

import (
	"fmt"
	// "regexp"
	"github.com/DATA-DOG/go-sqlmock"
)

func NewMockDbHandler() *SqlHandler {
	mockDb, mock, err := sqlmock.New()
	if err != err {
		fmt.Println("error creating mock database")
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = mockDb

	// rows := sqlmock.NewRows([]string{"id", "title"}).
	// 	AddRow("1", "one").
	// 	AddRow("2", "two")

	prep := mock.ExpectPrepare("^INSERT INTO articles*")
    prep.ExpectExec().
        WithArgs("test", "test").
        WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "students"`)).
	// 	WillReturnRows(rows)
	// mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	fmt.Println("mockDB接続")
	return sqlHandler
}
