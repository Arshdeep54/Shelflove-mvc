package main

import (
	"fmt"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	"log"
	"os"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("Initillising Mysql database ")
	db, err := config.DbConnection()
	if err != nil {
		fmt.Println("%w", err)
		os.Exit(1)
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

	if err := migrateTables(db); err != nil {
		log.Fatal(err)
	}
	if err := MigrateDummyData(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrations completed.")
}

func migrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(&types.User{}, &types.Book{}, &types.Issue{})
	if err != nil {
		return fmt.Errorf("error creating user table: %w", err)
	}
	return nil

}
func MigrateDummyData(db *gorm.DB) error {
	books := utils.GetDummyBooks()
	tx := db.Create(&books)
	if tx.Error != nil {
		return fmt.Errorf("error adding dummy books data: %w", tx.Error)
	}
	return nil
}
func cleanDB(db *gorm.DB) error {
	db.Migrator().DropTable(&types.User{})
	db.Migrator().DropTable(&types.Book{})
	db.Migrator().DropTable(&types.Issue{})
	return nil
}
