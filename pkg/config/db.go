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

// var DbPath = false

func DbConnection() (*sql.DB, error) {
	// err:=loadEnv()
	// if err!=nil{
	// 	return nil,err
	// }
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	MYSQL_USERNAME := os.Getenv("MYSQL_USERNAME")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)

	Db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	const MAX_OPEN_CONNECTIONS = 20

	const MAX_IDLE_CONNECTIONS = 20

	const MAX_LIFETIME = 5 // minutes

	Db.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	Db.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	Db.SetConnMaxLifetime(time.Minute * MAX_LIFETIME)

	err = Db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return Db, nil
}

// func loadEnv() error {
// 	if DbPath {
// 		err := godotenv.Load("../../.env")
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		err := godotenv.Load()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
