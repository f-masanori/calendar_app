package database

import (
	"fmt"
	"log"

	"go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
)

type NikkiRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *NikkiRepository) FindAll() (entities.Nikkis, error) {
	var nikkis entities.Nikkis
	fmt.Println("show nikkis")
	rows, err := repo.SqlHandler.DB.Query("SELECT * from nikkis;")
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	defer rows.Close()

	var nikkis_table_colum Nikkis_table
	for rows.Next() {
		var nikki entities.Nikki
		err := rows.Scan(&nikkis_table_colum.Id, &nikkis_table_colum.User_id,
			&nikkis_table_colum.Created_at, &nikkis_table_colum.Updated_at,
			&nikkis_table_colum.Date, &nikkis_table_colum.Content,
			&nikkis_table_colum.Title)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		nikki.Id = nikkis_table_colum.Id
		nikki.User_id = nikkis_table_colum.User_id
		nikki.Title = nikkis_table_colum.Title
		nikki.Date = nikkis_table_colum.Date
		nikki.Content = nikkis_table_colum.Content
		nikkis = append(nikkis, nikki)
	}
	return nikkis, nil
}

// func (repo *NikkiRepository) DeleteNikki()  {
// 	var nikki entities.Nikki
// 	statement := "DELETE FROM nikkis WHERE id = ?"
// 	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
// 	if err != nil {
// 		fmt.Println("Prepare(statement) error")
// 		return nikki, err
// 	}
// }
func (repo *NikkiRepository) CreateNikki(UserId int, Date int, Title string, Content string) (entities.Nikki, error) {
	var nikki entities.Nikki
	statement := "INSERT INTO nikkis(user_id,date,title,content) VALUES(?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare(statement) error")
		return nikki, err
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(UserId, Date, Title, Content)
	if err != nil {
		fmt.Println("stmtInsert.Exec　error")
		return nikki, err
	}
	lastInsertID, err := result.LastInsertId()

	err = repo.SqlHandler.DB.QueryRow("SELECT id,user_id,date,title,content FROM nikkis WHERE id = ?", lastInsertID).Scan(&nikki.Id, &nikki.User_id,
		&nikki.Date, &nikki.Title, &nikki.Content)
	if err != nil {
		log.Fatal(err)
	}

	return nikki, nil
}
func (repo *NikkiRepository) DeleteNikki(UserId int, Date int) (int, int, int, error) {
	stmtDelete, err := repo.SqlHandler.DB.Prepare("DELETE FROM nikkis WHERE user_id = ? and date = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(UserId, Date)
	if err != nil {
		panic(err.Error())
	}

	_rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(_rowsAffect)
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		fmt.Println("削除できません")
	} else if rowsAffect == 1 {
		fmt.Println("削除完了")
	} else {
		fmt.Println("DB table エラー") //削除データが2個以上は起らないはず
	}
	return UserId, Date, rowsAffect, err
}
func (repo *NikkiRepository) EditNikki(UserId int, Date int, Title string, Content string) {
	stmtEdit, err := repo.SqlHandler.DB.Prepare("UPDATE nikkis SET title = ?,content = ? WHERE user_id = ? and date = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtEdit.Close()

	result, err := stmtEdit.Exec(Title, Content,2,20191128)
	if err != nil {
		panic(err.Error())
	}
	_rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(_rowsAffect)
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		fmt.Println("削除できません")
	} else if rowsAffect == 1 {
		fmt.Println("削除完了")
	} else {
		fmt.Println("DB table エラー") //削除データが2個以上は起らないはず
	}
}

// func (repo *UserRepository) FindAll() (entities.Users, error) {
// 	var users entities.Users

// 	fmt.Println("show users")
// 	rows, err := repo.SqlHandler.DB.Query("SELECT * from users;")
// 	if err != nil {
// 		log.Print("error executing database query: ", err)
// 	}
// 	defer rows.Close() // make sure rows is closed when the handler exits
// 	defer fmt.Println("どこで終了かの確認")
// 	// type users_table struct {
// 	// 	Id         int    `db:"id"`
// 	// 	Name       string `db:"name"`
// 	// 	Created_at string `db:"created_at"`
// 	// 	Updated_at string `db:"updated_at"`
// 	// }
// 	var users_table_colum Users_table
// 	for rows.Next() {
// 		var user entities.User
// 		err := rows.Scan(&users_table_colum.Id, &users_table_colum.Name, &users_table_colum.Created_at, &users_table_colum.Updated_at)
// 		if err != nil {
// 			fmt.Println(err)
// 			panic(err.Error())
// 		}
// 		user.Id = users_table_colum.Id
// 		user.Name = users_table_colum.Name
// 		users = append(users, user)
// 	}

// 	return users, nil
// }
