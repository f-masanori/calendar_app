package database

type Users_table struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Created_at string `db:"created_at"`
	Updated_at string `db:"updated_at"`
}
