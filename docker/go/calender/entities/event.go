package entities

type Event struct {
	ID    int
	UID   string
	Date  string
	Event string
}

type Events []Event
