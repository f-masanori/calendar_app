package firebaseauth

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("テスト中")
		next.ServeHTTP(w, r)
	}
}
