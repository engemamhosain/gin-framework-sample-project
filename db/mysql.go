package db

import (
	"database/sql"
	"fmt"
	"gin-sample-project/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitializeMySQL initializes a MySQL connection pool
func InitializeMySQL() {
	// dsn := "user:password@tcp(localhost:3306)/dbname?parseTime=true"
	var err error

	DB, err = sql.Open("mysql", config.Config.GetString("app.master-url"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Configure the connection pool
	DB.SetMaxOpenConns(25)                 // Maximum number of open connections to the database
	DB.SetMaxIdleConns(25)                 // Maximum number of connections in the idle connection pool
	DB.SetConnMaxLifetime(5 * time.Minute) // Maximum amount of time a connection may be reused

	// Test the connection to ensure everything is set up properly
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Database connection established with pooling.")
}
