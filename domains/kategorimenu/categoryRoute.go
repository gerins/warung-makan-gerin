package kategorimenu

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitKategoriMenuRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	KategoriMenuController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", KategoriMenuController.HandleGETAllKategoriMenus()).Methods("GET")
	p.HandleFunc("/{id}", KategoriMenuController.HandleGETKategoriMenu()).Methods("GET")
	p.HandleFunc("", KategoriMenuController.HandlePOSTKategoriMenus()).Methods("POST")
	p.HandleFunc("/{id}", KategoriMenuController.HandleUPDATEKategoriMenus()).Methods("PUT")
	p.HandleFunc("/{id}", KategoriMenuController.HandleDELETEKategoriMenus()).Methods("DELETE")
}
