package router

import (
	"fmt"
	"log"
	"net/http"
	// "os"

	Authentication "github.com/f-masanori/my-nikki_dev/docker/go/infrastructure"
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
	"github.com/f-masanori/my-nikki_dev/docker/go/interfaces/handlers"

	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()

	DBhandler := database.NewSqlHandler()

	// MockDBhandler := database.NewMockDbHandler()
	// sss,err:=MockDBhandler.DB.Exec("SELECT * FROM students")
	// fmt.Println(sss)
	// fmt.Println(err)
	// ins, err := MockDBhandler.DB.Prepare("INSERT INTO articles(title,content) VALUES(?,?)")
    // if err != nil {
    //     log.Fatal(err)
    // }
	// wwww,err:= MockDBhandler.DB.Query("SELECT * from users")
	// fmt.Println(wwww)
	// if err != nil {
    //     log.Fatal(err)
    // }

	userHandler := handlers.NewUserHandler(DBhandler)
	platformHandler := handlers.NewPlatformHandler(DBhandler)

	router.HandleFunc("/", userHandler.Index).Methods("GET")
	router.HandleFunc("/test", userHandler.Test).Methods("GET")
	router.HandleFunc("/testauth", Authentication.AuthMiddleware(userHandler.Index)).Methods("GET")
	router.HandleFunc("/user", userHandler.NewUser).Methods("POST")
	router.HandleFunc("/initalize", platformHandler.Index).Methods("POST")

	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
