package menu

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitMenuRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	MenuController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", MenuController.HandleGETAllMenus()).Methods("GET")
	p.HandleFunc("/{id}", MenuController.HandleGETMenu()).Methods("GET")
	p.HandleFunc("", MenuController.HandlePOSTMenus()).Methods("POST")
	p.HandleFunc("/{id}", MenuController.HandleUPDATEMenus()).Methods("PUT")
	p.HandleFunc("/{id}", MenuController.HandleDELETEMenus()).Methods("DELETE")
}
