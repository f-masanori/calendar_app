package handlers

import (

	// "go_docker/calender/entities"

	// "fmt"
	"go_docker/calender/infrastructure/database"
	sqlcmd "go_docker/calender/interfaces/database"
	"go_docker/calender/services"
	"net/http"
)

type EventHandler struct {
	Service *services.EventService
}

func NewEventHandler(sqlHandler *database.SqlHandler) *EventHandler {
	return &EventHandler{
		Service: &services.EventService{
			EventRepository: &sqlcmd.EventRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (e *EventHandler) AddEvent(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(e.Body)
	// fmt.Println(e.Method)
	// uid := Authentication.UID
	// e.Service.CreateEvent("UIDDD", 20200322, "llll")
}
