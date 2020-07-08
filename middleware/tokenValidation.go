package middleware

import (
	"fmt"
	"log"
	"net/http"
	"warung_makan_gerin/utils/token"
)

func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getToken, err := r.Cookie("token")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if len(getToken.Value) == 0 {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		validity, userName, _ := token.VerifyToken(getToken.Value)
		fmt.Println(r.RequestURI + ` Accessed by ` + userName)
		if validity == true {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
		}
	})
}
