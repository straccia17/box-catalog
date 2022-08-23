package services

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB
var err error

func InitDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
	log.Println("Connection string to DB: ", connStr)
	DB, err = sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Connected to DB")
}
