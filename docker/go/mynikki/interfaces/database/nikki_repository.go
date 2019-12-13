package database

import (
	"fmt"
	"go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
	"log"
	"strconv"
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
		err := rows.Scan(
			&nikkis_table_colum.Id,
			&nikkis_table_colum.User_id,
			&nikkis_table_colum.Date,
			&nikkis_table_colum.Title,
			&nikkis_table_colum.Content,
			&nikkis_table_colum.NumberOfPhotos,
			&nikkis_table_colum.Created_at,
			&nikkis_table_colum.Updated_at)
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

func (repo *NikkiRepository) GetNikki(UserId int, Date int) {
	var nikkis_table_colum Nikkis_table
	var nikki entities.Nikki
	var nikkiRecordFlag = false

	rows, err := repo.SqlHandler.DB.Query("SELECT * FROM nikkis WHERE user_id = ? and date = ? LIMIT 1;", UserId, Date)
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	for rows.Next() {
		err := rows.Scan(
			&nikkis_table_colum.Id,
			&nikkis_table_colum.User_id,
			&nikkis_table_colum.Date,
			&nikkis_table_colum.Title,
			&nikkis_table_colum.Content,
			&nikkis_table_colum.NumberOfPhotos,
			&nikkis_table_colum.Created_at,
			&nikkis_table_colum.Updated_at)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		nikkiRecordFlag = true
		nikki.Id = nikkis_table_colum.Id
		nikki.User_id = nikkis_table_colum.User_id
		nikki.Title = nikkis_table_colum.Title
		nikki.Date = nikkis_table_colum.Date
		nikki.Content = nikkis_table_colum.Content
		fmt.Println("find nikki userId = " + strconv.Itoa(nikkis_table_colum.User_id) + " date = " + strconv.Itoa(nikkis_table_colum.Date))
		// nikkis = append(nikkis, nikki)
	}
	if nikkiRecordFlag && nikkis_table_colum.NumberOfPhotos != 0 {
		var Photos entities.Photos

		rows, err := repo.SqlHandler.DB.Query("SELECT id,photo FROM photos WHERE nikki_id = ? LIMIT ?;",
			 nikkis_table_colum.Id, nikkis_table_colum.NumberOfPhotos)
		if err != nil {
			log.Print("error executing database query: ", err)
		}

		for rows.Next() {
			var photo entities.Photo
			err := rows.Scan(&photo.Id, &photo.Photo)
			if err != nil {
				fmt.Println(err)
				panic(err.Error())
			}
			Photos = append(Photos, photo)
		}
		nikki.Photos = Photos
	}
	// err := repo.SqlHandler.DB.QueryRow("SELECT * FROM nikkis WHERE user_id = ? and date = ? LIMIT 1;", UserId, Date).Scan(
	// 		&nikkis_table_colum.Id,
	// 		&nikkis_table_colum.User_id,
	// 		&nikkis_table_colum.Date,
	// 		&nikkis_table_colum.Title,
	// 		&nikkis_table_colum.Content,
	// 		&nikkis_table_colum.NumberOfPhotos,
	// 		&nikkis_table_colum.Created_at,
	// 		&nikkis_table_colum.Updated_at)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(nikki)
	fmt.Println(nikkis_table_colum)
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
func (repo *NikkiRepository) CreateNikki(UserId int, Date int, Title string, Content string, NumberOfPhotos int) (entities.Nikki, error) {
	var nikki entities.Nikki
	statement := "INSERT INTO nikkis(user_id,date,title,content,number_of_photos) VALUES(?,?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare(statement) error")
		return nikki, err
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(UserId, Date, Title, Content, NumberOfPhotos)
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

	result, err := stmtEdit.Exec(Title, Content, 2, 20191128)
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

func (repo *NikkiRepository) RegisterPhoto(UserId int, Date int) {

}
