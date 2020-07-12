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
			// http.Redirect(w, r, "http://localhost:5500/index.html", http.StatusSeeOther)
			http.Error(w, "Token Expired", http.StatusUnauthorized)
			return
		}

		validity, userName, _ := token.VerifyToken(getCookies[0].Value)
		fmt.Println(userName + " accessing " + r.RequestURI)
		if validity == true && userName == getCookies[1].Value {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Sessions", http.StatusUnauthorized)
		}
	})
}
