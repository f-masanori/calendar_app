package database

import (
	"fmt"
	"log"

	"github.com/f-masanori/my-nikki_dev/docker/go/entities"
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
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
	type users_table struct {
		Id         int    `db:"id"`
		Name       string `db:"name"`
		Created_at string `db:"created_at"`
		Updated_at string `db:"updated_at"`
	}
	var users_table_colum users_table
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&users_table_colum.Id, &users_table_colum.Name, &users_table_colum.Created_at, &users_table_colum.Updated_at)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		user.Id = users_table_colum.Id
		user.Email = users_table_colum.Name
		users = append(users, user)
	}

	return users, nil
}

// func (repo *UserRepository) SaveUser(entities.Users, error){

// }
// func (repo *UserRepository) Find(id int) (entities.User, error) {
// 	var user entities.User

// 	if err := repo.SqlHandler.Conn.Where("id = ?", id).First(&user).Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
