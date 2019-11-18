package router

import (
	"fmt"
	"log"
	"net/http"

	Authentication "github.com/f-masanori/my-nikki_dev/docker/go/infrastructure"
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
	"github.com/f-masanori/my-nikki_dev/docker/go/interfaces/handlers"

	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()

	DBhandler := database.NewSqlHandler()

	userHandler := handlers.NewUserHandler(DBhandler)
	platformHandler := handlers.NewPlatformHandler(DBhandler)

	router.HandleFunc("/", userHandler.Index).Methods("GET")
	router.HandleFunc("/test", userHandler.Test).Methods("GET")
	router.HandleFunc("/testauth", Authentication.AuthMiddleware(userHandler.Index)).Methods("GET")

	// router.HandleFunc("/users", showusers).Methods("GET")
	router.HandleFunc("/user", userHandler.NewUser).Methods("POST")
	router.HandleFunc("/initalize", platformHandler.Index).Methods("POST")
	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
