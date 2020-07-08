package user

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitUserRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	UserController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", UserController.HandleGETAllUsers()).Methods("GET")
	p.HandleFunc("/login", UserController.HandleUserLogin()).Queries("user", "{user}", "password", "{password}").Methods("GET")
	p.HandleFunc("/register", UserController.HandleRegisterNewUser()).Methods("POST")
	p.HandleFunc("/{id}", UserController.HandleUPDATEUsers()).Methods("PUT")
	p.HandleFunc("/{id}", UserController.HandleDELETEUsers()).Methods("DELETE")
}

/*
Untuk login http://localhost:8080/user/login?user=gerin&password=admin
Untuk register http://localhost:8080/user/register
{
    "username": "gerin",
    "password": "admin"
}
*/
