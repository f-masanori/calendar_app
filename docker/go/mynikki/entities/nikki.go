package entities

type Nikki struct {
	ID      int
	UserID int
	Date    int
	Title   string
	Content string
	Photos  Photos
}

type Nikkis []Nikki

type Photo struct {
	ID    int
	Photo string
}
type Photos []Photo
