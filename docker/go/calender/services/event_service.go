package services

import (
	"fmt"
	"go_docker/calender/entities"
)

type EventRepository interface {
	CreateEvent(string, string, string)
	GetEventsByUID(string) (entities.Events, error)
	// FindAll() (entities.Nikkis, error)
	// FindNikki(int, int) (entities.Nikki, error)
	// CreateNikki(int, int, string, string, int) (entities.Nikki, error)
	// DeleteNikki(int, int) (int, int, int, error)

}
type EventService struct {
	EventRepository EventRepository
}

func (e *EventService) CreateEvent(uid string, date string, event string) {
	e.EventRepository.CreateEvent(uid, date, event)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}
func (e *EventService) GetEventsByUID(uid string) entities.Events {
	events, err := e.EventRepository.GetEventsByUID(uid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(events)
	return events
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}
