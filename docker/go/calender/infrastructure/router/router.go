package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	// "os"

	Authentication "go_docker/calender/infrastructure"
	"go_docker/calender/infrastructure/database"
	"go_docker/calender/interfaces/handlers"

	// gorillaHundler "github.com/gorilla/handlers"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
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
	opt := option.WithCredentialsFile("./calender-9275a-firebase-adminsdk-c09d7-691abcc199.json")
	// fmt.Print(opt)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// return nil, fmt.Errorf("error initializing app: %v", err)
	}
	// fmt.Print(app)
	// // // Access auth service from the default app
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	authHeader := r.Header.Get("Authorization")

	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	// fmt.Println(r.Body)

	// // JWT の検証
	token, err := auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {

		u := fmt.Sprintf("error verifying ID token: %v\n", err)
		fmt.Print(u)
		// return c.JSON(http.StatusBadRequest, u)
	}
	uid := token.Claims["user_id"]
	fmt.Println(uid)
}
func Init() {

	router := mux.NewRouter()
	DBhandler := database.NewSqlHandler()
	router.Use(CORS)
	nikkiHandler := handlers.NewNikkiHandler(DBhandler)
	userHandler := handlers.NewUserHandler(DBhandler)
	// platformHandler := handlers.NewPlatformHandler(DBhandler)
	eventHandler := handlers.NewEventHandler(DBhandler)
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
	router.HandleFunc("/addEvent", Authentication.AuthMiddleware(eventHandler.AddEvent))
	router.HandleFunc("/getEventsByUID", Authentication.AuthMiddleware(eventHandler.GetEventsByUID))
	/* userHandler.Userdelete で指定idのユーザーを削除したい */
	// router.HandleFunc("/user/delete/:id", userHandler.UserDelete).Methods("POST")
	/* userHandler.Userinfo　でそのユーザーに紐づけられているdevice、認証などの情報を取得したい(未実装) */
	// router.HandleFunc("/user/:id", userHandler.Userinfo).Methods("GET")
	// router.HandleFunc("/initalize", platformHandler.Index).Methods("POST")

	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
