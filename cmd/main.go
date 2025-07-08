package main

import (
	"fmt"
	"inventory_management_system/config"
	"inventory_management_system/database"
	"os"
)

func main() {
	config.LoadEnv()
	dbConnectionString := config.GetDatabaseString()
	database.Init(dbConnectionString)
	defer database.DB.Close()
	fmt.Println("Starting server on port " + os.Getenv("SERVER_PORT"))
}
