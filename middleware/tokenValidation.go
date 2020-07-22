package middleware

import (
	"log"
	"net/http"
	"warung_makan_gerin/utils/token"
)

// Validate Token from cookies
func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getCookies := r.Cookies()
		if len(getCookies) == 0 {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
			return
		}

		validity, userName, _ := token.VerifyToken(getCookies[0].Value)
		log.Println(userName + " accessing " + r.RequestURI)
		if validity == true && userName == getCookies[1].Value {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Sessions", http.StatusUnauthorized)
		}
	})
}
