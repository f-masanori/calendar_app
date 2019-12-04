package entities

type Nikki struct {
	Id      int
	User_id int
	Title   string
	Content string
	Date    int
}

type Nikkis []Nikki
