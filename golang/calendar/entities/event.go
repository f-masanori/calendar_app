package entities

type Event struct {
	ID      int
	UID     string
	EventID int
	Date    string
	Event   string
}

type Events []Event
