package database

import (
	"fmt"
	"go_docker/calender/entities"
	"go_docker/calender/infrastructure/database"
	"log"
)

type EventRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *EventRepository) CreateEvent(uid string, date string, event string) {
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

//今は全部持って来てます
func (repo *EventRepository) GetEventsByUID(UID string) (entities.Events, error) {
	var events entities.Events
	fmt.Println(UID)
	rows, err := repo.SqlHandler.DB.Query("SELECT * from events WHERE uid = ?;", UID)
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	defer rows.Close()

	var events_table_colum Events_table
	for rows.Next() {
		var event entities.Event
		err := rows.Scan(
			&events_table_colum.ID,
			&events_table_colum.UID,
			&events_table_colum.Date,
			&events_table_colum.Event,
			&events_table_colum.CreatedAt,
			&events_table_colum.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		event.ID = events_table_colum.ID
		event.UID = events_table_colum.UID
		event.Date = events_table_colum.Date
		event.Event = events_table_colum.Event

		events = append(events, event)
	}
	return events, nil
}
