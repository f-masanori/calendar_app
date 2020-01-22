package router

import (
	"fmt"
	"log"
	"net/http"
	// "os"

	Authentication "go_docker/calender/infrastructure"
	"go_docker/calender/infrastructure/database"
	"go_docker/calender/interfaces/handlers"

	"github.com/gorilla/mux"
	gorillaHundler "github.com/gorilla/handlers"
)

func Init() {
	router := mux.NewRouter()
	DBhandler := database.NewSqlHandler()
	cors := gorillaHundler.CORS(
		gorillaHundler.AllowedHeaders([]string{"content-type"}),
		gorillaHundler.AllowedOrigins([]string{"*"}),
		gorillaHundler.AllowCredentials(),
	)
	router.Use(cors)
	
	nikkiHandler := handlers.NewNikkiHandler(DBhandler)
	userHandler := handlers.NewUserHandler(DBhandler)
	platformHandler := handlers.NewPlatformHandler(DBhandler)

	router.HandleFunc("/nikkis", nikkiHandler.Index).Methods("GET")
	router.HandleFunc("/nikki", nikkiHandler.CreateNikki).Methods("POST")
	router.HandleFunc("/nikki/{userID}/{date}", nikkiHandler.GetNikki).Methods("GET")
	router.HandleFunc("/nikki/delete", nikkiHandler.DeleteNikki).Methods("POST")
	router.HandleFunc("/nikki/edit", nikkiHandler.EditNikki).Methods("POST")

	router.HandleFunc("/registerphoto", nikkiHandler.RegisterPhoto).Methods("POST")

	router.HandleFunc("/users", userHandler.Index).Methods("GET")
	router.HandleFunc("/test", userHandler.Test).Methods("GET")
	router.HandleFunc("/testauth", Authentication.AuthMiddleware(userHandler.Index)).Methods("GET")
	router.HandleFunc("/user", userHandler.NewUser).Methods("POST")
	router.HandleFunc("/user/delete", userHandler.DeleteUser).Methods("POST")

	/* userHandler.Userdelete で指定idのユーザーを削除したい */
	// router.HandleFunc("/user/delete/:id", userHandler.UserDelete).Methods("POST")
	/* userHandler.Userinfo　でそのユーザーに紐づけられているdevice、認証などの情報を取得したい(未実装) */
	// router.HandleFunc("/user/:id", userHandler.Userinfo).Methods("GET")
	router.HandleFunc("/initalize", platformHandler.Index).Methods("POST")

	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
