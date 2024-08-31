package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDatabase(User string, Passwd string, DBName string) *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   User,
		Passwd: Passwd,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: DBName,
	}
	// Get a database handle.
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Check that the database is accessible.
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database :", DBName)

	return DB
}
