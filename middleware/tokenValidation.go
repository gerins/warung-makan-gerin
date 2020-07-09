package middleware

import (
	"fmt"
	"net/http"
	"warung_makan_gerin/utils/token"
)

func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getCookies := r.Cookies()
		fmt.Println(getCookies)
		if len(getCookies) == 0 {
			http.Redirect(w, r, "127.0.0.1:5500/index.html", http.StatusMovedPermanently)
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
