package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init(connectionStr string) error {
	fmt.Println("Connecting to:", connectionStr)
	var err error
	DB, err = sqlx.Connect("postgres", connectionStr)
	if err != nil {
		return fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	fmt.Println("Connected to PostgreSQL")
	return nil
}
