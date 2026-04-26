package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "user=postgres password=1608 dbname=goLang_practice host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error opening DB", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting DB", err)
		panic(err)
	}

	fmt.Println("Successfully connected to DB")
	DB = db
}
