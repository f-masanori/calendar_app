package database

type Users_table struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Created_at string `db:"created_at"`
	Updated_at string `db:"updated_at"`
}
type Nikkis_table struct {
	Id         int    `db:"id"`
	User_id    int    `db:"user_id"`
	Created_at string `db:"created_at"`
	Updated_at string `db:"updated_at"`
	Date       int    `db:"date"`
	Content    string `db:"content"`
	Title      string `db:"title"`
}
