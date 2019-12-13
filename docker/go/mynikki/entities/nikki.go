package entities

type Nikki struct {
	Id      int
	User_id int
	Date    int
	Title   string
	Content string
	Photos  Photos
}

type Nikkis []Nikki

type Photo struct {
	Id    int
	Photo string
}
type Photos []Photo
