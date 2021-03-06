package router

import (
	"fmt"

	Authentication "golang/calendar/infrastructure"
	"golang/calendar/infrastructure/database"
	"golang/calendar/interfaces/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CORS対応 Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CORS")
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

func Init() {
	router := mux.NewRouter()
	DBhandler := database.NewSqlHandler()
	router.Use(CORS)
	userHandler := handlers.NewUserHandler(DBhandler)
	eventHandler := handlers.NewEventHandler(DBhandler)
	allInOneHandler := handlers.NewAllInOneHandler(DBhandler)
	todoHandler := handlers.NewTodoHandler(DBhandler)

	router.HandleFunc("/addEvent", Authentication.AuthMiddleware(eventHandler.AddEvent))
	router.HandleFunc("/getEventsByUID", Authentication.AuthMiddleware(eventHandler.GetEventsByUID))
	router.HandleFunc("/registerUser", userHandler.NewUser)
	router.HandleFunc("/deleteEvent", Authentication.AuthMiddleware(eventHandler.DeleteEvent))
	router.HandleFunc("/editEvent", Authentication.AuthMiddleware(eventHandler.EditEvent))
	router.HandleFunc("/getNextEventID", Authentication.AuthMiddleware(eventHandler.GetNextEventID))

	router.HandleFunc("/addScript", Authentication.AuthMiddleware(allInOneHandler.AddScript))

	router.HandleFunc("/addTodo", Authentication.AuthMiddleware(todoHandler.AddTodo))
	router.HandleFunc("/deleteTodo", Authentication.AuthMiddleware(todoHandler.DeleteTodo))

	router.HandleFunc("/getTodosByUID", Authentication.AuthMiddleware(todoHandler.GetTodosByUID))

	/* 以下はカレンダーアプリでは使用していません */
	// nikkiHandler := handlers.NewNikkiHandler(DBhandler)
	// platformHandler := handlers.NewPlatformHandler(DBhandler)
	// router.HandleFunc("/nikkis", nikkiHandler.Index).Methods("GET")
	// router.HandleFunc("/nikki", nikkiHandler.CreateNikki).Methods("POST")
	// router.HandleFunc("/nikki/{userID}/{date}", nikkiHandler.GetNikki).Methods("GET")
	// router.HandleFunc("/nikki/delete", nikkiHandler.DeleteNikki).Methods("POST")
	// router.HandleFunc("/nikki/edit", nikkiHandler.EditNikki).Methods("POST")
	// router.HandleFunc("/registerphoto", nikkiHandler.RegisterPhoto).Methods("POST")
	// router.HandleFunc("/users", userHandler.Index).Methods("GET")
	// router.HandleFunc("/test", userHandler.Test).Methods("GET")
	// router.HandleFunc("/testauth", Authentication.AuthMiddleware(userHandler.Index)).Methods("GET")
	// router.HandleFunc("/user", userHandler.NewUser).Methods("POST")
	// router.HandleFunc("/user/delete", userHandler.DeleteUser).Methods("POST")
	/* userHandler.Userdelete で指定idのユーザーを削除したい */
	// router.HandleFunc("/user/delete/:id", userHandler.UserDelete).Methods("POST")
	/* userHandler.Userinfo　でそのユーザーに紐づけられているdevice、認証などの情報を取得したい(未実装) */
	// router.HandleFunc("/user/:id", userHandler.Userinfo).Methods("GET")
	// router.HandleFunc("/initalize", platformHandler.Index).Methods("POST")
	/**/
	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
