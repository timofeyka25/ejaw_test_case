package database

import (
	"database/sql"
	"ejaw_test_case/pkg/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	cfg := config.GetDB()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Successfully connected to the database")
	return db, nil
}
