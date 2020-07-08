package main

import (
	"warung_makan_gerin/config/database"
	"warung_makan_gerin/config/router"
	"warung_makan_gerin/middleware"
)

func main() {
	db := database.ConnectDB()
	r := router.CreateRouter()

	r.Use(middleware.LoggingMiddleware)

	router.NewAppRouter(db, r).InitRouter()
	router.StartServer(r)
}
