package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/user"
	"warung_makan_gerin/domains/kategorimenu"
	"warung_makan_gerin/domains/menu"
	"warung_makan_gerin/domains/transaction"
	"warung_makan_gerin/domains/user"

	"github.com/gorilla/mux"
)

const (
	MENUS_MAIN_ROUTE       = "/menus"
	CATEGORY_MAIN_ROUTE    = "/categorymenus"
	TRANSACTION_MAIN_ROUTE = "/transaction"
	USER_MAIN_ROUTE        = "/user"
)

type ConfigRouter struct {
	DB     *sql.DB
	Router *mux.Router
}

func (ar *ConfigRouter) InitRouter() {
	menu.InitMenuRoute(MENUS_MAIN_ROUTE, ar.DB, ar.Router)
	kategorimenu.InitKategoriMenuRoute(CATEGORY_MAIN_ROUTE, ar.DB, ar.Router)
	transaction.InitTransactionRoute(TRANSACTION_MAIN_ROUTE, ar.DB, ar.Router)
	user.InitUserRoute(USER_MAIN_ROUTE, ar.DB, ar.Router)
	ar.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

// NewAppRouter for creating new Route
func NewAppRouter(db *sql.DB, r *mux.Router) *ConfigRouter {
	return &ConfigRouter{
		DB:     db,
		Router: r,
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}
