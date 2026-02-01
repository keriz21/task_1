package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// open database
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// test connection
	err = db.Ping()
	if err != nil {
		return nil , err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(0)
	// db.SetConnMaxLifetime(30 * time.Minute)


	log.Println("Database connected successfully")
	return db, nil

}