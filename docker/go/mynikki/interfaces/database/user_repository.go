package database

import (
	"fmt"
	"log"

	"go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
)

type UserRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *UserRepository) FindAll() (entities.Users, error) {
	var users entities.Users

	fmt.Println("show users")
	rows, err := repo.SqlHandler.DB.Query("SELECT * from users;")
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	defer rows.Close() // make sure rows is closed when the handler exits
	defer fmt.Println("どこで終了かの確認")
	// type users_table struct {
	// 	Id         int    `db:"id"`
	// 	Name       string `db:"name"`
	// 	Created_at string `db:"created_at"`
	// 	Updated_at string `db:"updated_at"`
	// }
	var users_table_colum Users_table
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&users_table_colum.Id, &users_table_colum.Name, &users_table_colum.Created_at, &users_table_colum.Updated_at)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		user.Id = users_table_colum.Id
		user.Name = users_table_colum.Name
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) CreateUser(name string) (entities.User, error) {
	var user entities.User
	statement := "INSERT INTO users(name) VALUES(?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("error1")
		return user, err
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(name)
	if err != nil {
		fmt.Println("error2")
		return user, err
	}
	lastInsertID, err := result.LastInsertId()
	user.Id = int(lastInsertID)
	user.Name = name
	return user, nil
}

// func (repo *UserRepository) Find(id int) (entities.User, error) {
// 	var user entities.User

// 	if err := repo.SqlHandler.Conn.Where("id = ?", id).First(&user).Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
func NewUserRepository(sqlHandler *database.SqlHandler) *UserRepository {
	return &UserRepository{
		SqlHandler: sqlHandler,
	}
}
