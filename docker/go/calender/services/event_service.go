package services

type EventRepository interface {
	CreateEvent(string, int, string)
	// FindAll() (entities.Nikkis, error)
	// FindNikki(int, int) (entities.Nikki, error)
	// CreateNikki(int, int, string, string, int) (entities.Nikki, error)
	// DeleteNikki(int, int) (int, int, int, error)

}
type EventService struct {
	EventRepository EventRepository
}

func (e *EventService) CreateEvent(uid string, date int, event string) {
	e.EventRepository.CreateEvent(uid, date, event)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}
