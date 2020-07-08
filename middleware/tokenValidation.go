package middleware

import (
	"net/http"
	"strings"
	"warung_makan_gerin/utils/token"

	"github.com/gorilla/mux"
)

func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokens := r.Header.Get("Authorization")
		if len(tokens) == 0 {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		param := mux.Vars(r)
		validity, userName, _ := token.VerifyToken(strings.Split(tokens, " ")[1])

		if validity == true && param["id"] == userName {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
		}
	})
}
