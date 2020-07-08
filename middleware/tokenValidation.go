package middleware

import (
	"fmt"
	"net/http"
	"warung_makan_gerin/utils/token"
)

func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getCookies := r.Cookies()
		if len(getCookies) == 0 {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		validity, userName, _ := token.VerifyToken(getCookies[0].Value)
		fmt.Println(r.RequestURI + ` Accessed by ` + userName)
		if validity == true && userName == getCookies[1].Value {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Sessions", http.StatusUnauthorized)
		}
	})
}
