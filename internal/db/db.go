package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(url string) {
	var err error
	DB, err = sqlx.Connect("postgres", url)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	log.Println("Database connection established")
}
