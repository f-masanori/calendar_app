package database

import (
	"fmt"
	"go_docker/calender/infrastructure/database"
)

type EventRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *EventRepository) CreateEvent(uid string, date int, event string) {
	statement := "INSERT INTO events(uid,date,event) VALUES(?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare(statement) error")
		// return nikki, err
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(uid, date, event)
	fmt.Println(result)
	if err != nil {
		fmt.Println("stmtInsert.Exec　error")
		// return nikki, err
	}
}
