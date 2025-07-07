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
	err := database.Init(dbConnectionString)
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		os.Exit(1)
	}
	defer database.DB.Close()
	fmt.Println("postgres connected.")

	fmt.Println("Starting server on port " + os.Getenv("SERVER_PORT"))
}
