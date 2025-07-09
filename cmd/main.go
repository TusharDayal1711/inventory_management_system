package main

import (
	"fmt"
	"inventory_management_system/config"
	"inventory_management_system/database"
	"inventory_management_system/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadEnv()
	dbConnectionString := config.GetDatabaseString()
	database.Init(dbConnectionString)
	defer database.DB.Close()

	r := routes.GetRoutes()
	fmt.Println("Starting server on port " + os.Getenv("DB_PORT"))
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
