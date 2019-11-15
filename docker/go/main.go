package main

import (
	"github.com/f-masanori/my-nikki_dev/docker/go/conf"
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/router"
	_ "github.com/go-sql-driver/mysql"
)

// 初期化処理
func init() {
	// 設定ファイル読み取り
	conf.ReadConf()
}
func main() {
	router.Init()
	// defer database.CloseConn()
	// defer fmt.Println("deferrr")
}

// func addname(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("/addname")
// 	decoder := json.NewDecoder(r.Body)
// 	var users Users
// 	error := decoder.Decode(&users)
// 	if error != nil {
// 		w.Write([]byte("json decode error" + error.Error() + "\n"))
// 	}
// 	// // json -> struct
// 	// json.NewDecoder(r.Body).Decode(&users)

// 	ins, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ins.Exec(users.Name)
// 	fmt.Fprint(w, "succes insert name="+users.Name)
// }

// func showusers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("show users")
// 	rows, err := db.Query("SELECT * from users;")
// 	if err != nil {
// 		log.Print("error executing database query: ", err)
// 		return
// 	}
// 	defer rows.Close() // make sure rows is closed when the handler exits

// 	usersJson := make([]Users, 0)
// 	for rows.Next() {
// 		var users Users
// 		err := rows.Scan(&users.Id, &users.Name, &users.Created_at, &users.Updated_at)
// 		if err != nil {
// 			fmt.Println(err)
// 			panic(err.Error())
// 		}
// 		fmt.Printf(users.Name + "\n")
// 		usersJson = append(usersJson, users)
// 	}

// 	je, err := json.Marshal(usersJson)
// 	if err != nil {
// 		fmt.Printf("エラー:%s", err)
// 		os.Exit(1)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprint(w, string(je))
// 	fmt.Printf(usersJson[4].Name + "\n")

// }
