package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func DbConnection() (*sql.DB, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	MYSQL_USERNAME := os.Getenv("MYSQL_USERNAME")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
	Db.SetConnMaxLifetime(time.Minute * 5)
	err = Db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return Db, nil
}
