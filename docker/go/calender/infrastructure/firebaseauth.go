package firebaseauth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// "reflect"
	"strings"

	firebase "firebase.google.com/go"
	// "golang.org/x/net/context"
	"google.golang.org/api/option"
)

var FirebaseUID string

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		opt := option.WithCredentialsFile("/go/src/go_docker/calender-9275a-firebase-adminsdk-c09d7-691abcc199.json")
		// fmt.Print(opt)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Errorf("error initializing app: %v", err)
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
		// fmt.Print(token)
		uid := token.Claims["user_id"]
		// fmt.Println(uid)
		FirebaseUID = uid.(string)
		// fmt.Println(reflect.TypeOf(FirebaseUID))
		// Firebase SDK のセットアップ
		// opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
		// app, err := firebase.NewApp(context.Background(), nil, opt)
		// if err != nil {
		//     fmt.Printf("error: %v\n", err)
		//     os.Exit(1)
		// }
		// auth, err := app.Auth(context.Background())
		// if err != nil {
		//     fmt.Printf("error: %v\n", err)
		//     os.Exit(1)
		// }

		// クライアントから送られてきた JWT 取得
		// authHeader := r.Header.Get("Authorization")
		// idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		// token, err := auth.VerifyIDToken(context.Background(), idToken)
		// if err != nil {
		//     // JWT が無効なら Handler に進まず別処理
		//     fmt.Printf("error verifying ID token: %v\n", err)
		//     w.WriteHeader(http.StatusUnauthorized)
		//     w.Write([]byte("error verifying ID token\n"))
		//     return
		// }
		// log.Printf("Verified ID token: %v\n", token)
		// log.Printf("テスト中")
		next.ServeHTTP(w, r)
	}
}
