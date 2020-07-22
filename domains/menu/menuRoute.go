package menu

import (
	"database/sql"
	"warung_makan_gerin/middleware"

	"github.com/gorilla/mux"
)

func InitMenuRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	MenuController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.Use(middleware.TokenValidation)
	p.HandleFunc("", MenuController.HandleGETAllMenus()).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods("GET")
	p.HandleFunc("/{id}", MenuController.HandleGETMenu()).Methods("GET")
	p.HandleFunc("", MenuController.HandlePOSTMenus()).Methods("POST")
	p.HandleFunc("/{id}", MenuController.HandleUPDATEMenus()).Methods("PUT")
	p.HandleFunc("/{id}", MenuController.HandleDELETEMenus()).Methods("DELETE")
}
