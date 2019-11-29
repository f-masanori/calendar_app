package database

import (
	"fmt"
	// "log"

	"go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
)

type NikkiRepository struct {
	SqlHandler *database.SqlHandler
}
func (repo *NikkiRepository) FindAll()  {
	aaa := new(entities.Nikki)
	fmt.Println(aaa)
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
func (repo *NikkiRepository) CreateNikki() (entities.Nikki, error)  {
	var nikki entities.Nikki
	statement := "INSERT INTO nikkis(user_id,date,title,content) VALUES(?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare(statement) error")
		return nikki, err
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec("2","20191128","testtilt","testcontenstttt")
	if err != nil {
		fmt.Println("stmtInsert.Exec　error")
		return nikki, err
	}
	lastInsertID, err := result.LastInsertId()
	nikki.Id = int(lastInsertID)
	nikki.Date = 27001211
	return nikki, nil
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
