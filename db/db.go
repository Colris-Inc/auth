package db

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/Colris-Inc/auth/configurations"
	_ "github.com/lib/pq"
)

var (
	dbHost     = config.AppConfig.DBHost
	dbPort     = config.AppConfig.DBPort
	dbUser     = config.AppConfig.DBUser
	dbPassword = config.AppConfig.DBPassword
	dbName     = config.AppConfig.DBName
)

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	log.Println("opening DB connection")
	sqlconn := ``
	sqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", sqlconn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")

	return db, err
}
