package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
	"github.com/f-masanori/my-nikki_dev/docker/go/interfaces/handlers"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	userHandler := handlers.NewUserHandler(database.NewSqlHandler())
	router.HandleFunc("/", userHandler.Index).Methods("GET")
	// router.HandleFunc("/users", showusers).Methods("GET")
	router.HandleFunc("/user", userHandler.CreateUser).Methods("POST")

	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
