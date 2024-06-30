package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	"github.com/joho/godotenv"
)

func main()  {
	fmt.Println("Initillising Mysql database ")
    db,err:=config.DbConnection()
	if err != nil {
        fmt.Println("%w", err)
		return 
    }
	exists, err := dbExists(db)
    if err != nil {
        fmt.Println("error checking database existence: %w", err)
		return 
    }
	cleandbFlag := os.Getenv("CLEANDB")
    if cleandbFlag == "true" {
        fmt.Println("Cleaning database...")
        err := cleanDB(db)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Database cleaned successfully")
        return
    }
	if !exists {
        _, err = config.Db.Exec("CREATE DATABASE shelflove;") 
        if err != nil {
            fmt.Println("error creating database: %w", err)
			return 
        }
        fmt.Println("Database created successfully")
    } else {
        fmt.Println("Database already exists")
    }
    if err := migrateTables(db); err != nil {
        log.Fatal(err)
    }
	if err := migrateDummyData(db); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Migrations completed.")
}
func dbExists(Db *sql.DB) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    dbName := os.Getenv("MYSQL_DATABASE")
    query := `SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?;`
    row := Db.QueryRow(query, dbName)
    var schemaName string
    err = row.Scan(&schemaName)
    if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("error checking database existence: %w", err)
    }
    return schemaName == dbName, nil 
}

func migrateTables(db *sql.DB) error {
    _, err := db.Exec(utils.Create_User_Table)
    if err != nil {
        return fmt.Errorf("error creating user table: %w", err)
    }
	_, err = db.Exec(utils.Create_Book_Table)
    if err != nil {
        return fmt.Errorf("error creating book table: %w", err)
    }
	_, err = db.Exec(utils.Create_Issue_Table)
    if err != nil {
        return fmt.Errorf("error creating issue table: %w", err)
    }
    return nil

}
func migrateDummyData(db *sql.DB) error {
	_, err := db.Exec(utils.Add_Dummy_Books)
    if err != nil {
        return fmt.Errorf("error adding dummy books data: %w", err)
    }
	return nil
}
func cleanDB(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM issue;")
    if err != nil {
        return err
    }
    _, err = db.Exec("DELETE FROM book;")
    if err != nil {
        return err
    }
    _, err = db.Exec("DELETE FROM user;")
    if err != nil {
        return err
    }
    return nil
}