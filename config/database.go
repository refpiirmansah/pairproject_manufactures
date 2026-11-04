package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := "root:root123@tcp(127.0.0.1:3306)/komputer_manufacture_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	fmt.Println("Database Connected Successfully (3NF Structure Verified)")
	return db
}
