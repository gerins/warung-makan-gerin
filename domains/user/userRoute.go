package user

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitUserRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	UserController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", UserController.HandleGETAllUsers()).Methods("GET")
	p.HandleFunc("/login", UserController.HandleUserLogin()).Methods("POST")
	p.HandleFunc("/register", UserController.HandleRegisterNewUser()).Methods("POST")
	p.HandleFunc("/logout", UserController.UserLogOut()).Methods("GET")
	p.HandleFunc("/{id}", UserController.HandleUPDATEUsers()).Methods("PUT")
	p.HandleFunc("/{id}", UserController.HandleDELETEUsers()).Methods("DELETE")
}
