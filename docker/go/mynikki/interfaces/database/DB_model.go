package database

type Users_table struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
type Nikkis_table struct {
	ID             int    `db:"id"`
	UserID         int    `db:"user_id"`
	Date           int    `db:"date"`
	Title          string `db:"title"`
	Content        string `db:"content"`
	NumberOfPhotos int    `db:"number_of_photos"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}
type Photos_table struct {
	ID        int    `db:"id"`
	NikkiID   int    `db:"nikki_id"`
	UserID    int    `db:"user_id"`
	Date      int    `db:"date"`
	PhotoId   int    `db:"photo_id"`
	Photo     string `db:"photo"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
