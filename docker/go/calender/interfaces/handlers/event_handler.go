package handlers

import (
	"encoding/json"
	"fmt"
	Authentication "go_docker/calender/infrastructure"
	"go_docker/calender/infrastructure/database"
	sqlcmd "go_docker/calender/interfaces/database"
	"go_docker/calender/services"
	"log"
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
	type Request struct {
		Date       string `json:"Date"`
		InputEvent string `json:"InputEvent"`
	}
	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	// fmt.Println(r.Body)
	// fmt.Println(r.Method)
	// uid := Authentication.FirebaseUID
	e.Service.CreateEvent(Authentication.FirebaseUID, request.Date, request.InputEvent)
}
func (e *EventHandler) GetEventsByUID(w http.ResponseWriter, r *http.Request) {
	// type Request struct {
	// 	Date       string `json:"Date"`
	// 	InputEvent string `json:"InputEvent"`
	// }
	// decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)
	// request := new(Request)
	// err := decoder.Decode(&request)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(request)
	// fmt.Println(r)
	// fmt.Println(r.FormValue("date"))
	// uid := Authentication.FirebaseUID
	Events := e.Service.GetEventsByUID(Authentication.FirebaseUID)
	fmt.Println(Events)

	jsonEvents, err := json.Marshal(Events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonEvents)
}
