package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./myblogs.db")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("database connection failed: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL ,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp TEXT NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("failed to create blogs table: %v", err)
	}

	log.Println("Successfully connected to the blogs database and ensured the Schema exists.")
	return nil
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database not Initialized!")
	}
	return DB
}
