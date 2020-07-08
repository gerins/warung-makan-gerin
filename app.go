package main

import (
	"warung_makan_gerin/config/database"
	"warung_makan_gerin/config/router"
)

func main() {
	db := database.ConnectDB()
	r := router.CreateRouter()

	router.NewAppRouter(db, r).InitRouter()
	router.StartServer(r)
}
