package middleware

import (
	"net/http"
)

func AccessCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		next.ServeHTTP(w, r)
	})
}
