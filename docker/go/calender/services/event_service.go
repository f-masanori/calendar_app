package services

import (
	"fmt"
	"go_docker/calender/entities"
)

type EventRepository interface {
	CreateEvent(string, int, string, string)
	GetEventsByUID(string) (entities.Events, int, error)
	DeleteEvent(string, int)
	GetNextEventID(string) int
}
type EventService struct {
	EventRepository EventRepository
}

func (e *EventService) CreateEvent(uid string, eventID int, date string, event string) {
	/* Event作成時にNextEventIDを更新する必要あり
	Event作成時には必ず必要な動作なのでe.EventRepository.CreateEventに
	入れ込む(トランザクション処理も可能になるため) */
	e.EventRepository.CreateEvent(uid, eventID, date, event)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}
func (e *EventService) GetEventsByUID(uid string) (entities.Events, int) {
	events, nextEventID, err := e.EventRepository.GetEventsByUID(uid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("events, nextEventID : ")
	fmt.Println(events, nextEventID)
	return events, nextEventID
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}
func (e *EventService) DeleteEvent(uid string, eventID int) {

	e.EventRepository.DeleteEvent(uid, eventID)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}
func (e *EventService) GetNextEventID(uid string) int {
	NextEventID := e.EventRepository.GetNextEventID(uid)
	return NextEventID
}
