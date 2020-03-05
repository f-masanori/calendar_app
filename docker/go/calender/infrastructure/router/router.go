package router

import (
	"fmt"
	"log"
	"net/http"

	// "os"

	Authentication "go_docker/calender/infrastructure"
	"go_docker/calender/infrastructure/database"
	"go_docker/calender/interfaces/handlers"

	// gorillaHundler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == "OPTIONS" {
			fmt.Println("プリフライト")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}
func loginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	wweew := r.Header.Get("token")
	fmt.Println(wweew)
}
func Init() {
	router := mux.NewRouter()
	DBhandler := database.NewSqlHandler()
	router.Use(CORS)
	nikkiHandler := handlers.NewNikkiHandler(DBhandler)
	userHandler := handlers.NewUserHandler(DBhandler)
	platformHandler := handlers.NewPlatformHandler(DBhandler)
	router.HandleFunc("/login", loginController)
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
	// router.HandleFunc("/addEvent", nikkiHandler.AddEvent).Methods("POST")
	router.HandleFunc("/addEvent", nikkiHandler.AddEvent)
	/* userHandler.Userdelete で指定idのユーザーを削除したい */
	// router.HandleFunc("/user/delete/:id", userHandler.UserDelete).Methods("POST")
	/* userHandler.Userinfo　でそのユーザーに紐づけられているdevice、認証などの情報を取得したい(未実装) */
	// router.HandleFunc("/user/:id", userHandler.Userinfo).Methods("GET")
	router.HandleFunc("/initalize", platformHandler.Index).Methods("POST")

	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
