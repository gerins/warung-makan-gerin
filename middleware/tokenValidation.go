package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"warung_makan_gerin/utils/token"
)

func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokens := r.Header.Get("Authorization")

		if len(tokens) == 0 {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		validity, userName, _ := token.VerifyToken(strings.Split(tokens, " ")[1])
		fmt.Println(r.RequestURI + ` Accessed by ` + userName)
		if validity == true {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
		}
	})
}
