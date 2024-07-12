package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *sql.DB
var DbPath = false

func DbConnection() (*gorm.DB, error) {
	err := loadEnv()
	if err != nil {
		return nil, err
	}
	MYSQL_USERNAME := os.Getenv("MYSQL_USERNAME")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)
	Db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return Db, nil
}

func loadEnv() error {
	if DbPath {
		err := godotenv.Load("../../.env")
		if err != nil {
			return err
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
func CloseConnection(Db *gorm.DB) error {
	sqlDb, err := Db.DB()
	if err != nil {
		return err

	}
	sqlDb.Close()
	return nil
}
