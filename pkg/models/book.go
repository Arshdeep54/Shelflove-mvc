package models

import (
	"fmt"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func GetAllBooks() ([]types.Book, error) {
	var books []types.Book
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("Error connecting to Db: %w", err)
	}
	rows, err := db.Query(`SELECT * FROM book WHERE quantity >=0 ORDER BY quantity DESC;`)
	if err != nil {
		return nil, fmt.Errorf("Error querying books: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var book types.Book
		// Scan each column of the current row into the Book struct fields
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationDate, &book.Quantity, &book.Genre, &book.Address, &book.Rating,&book.Description)
		if err != nil {
			return nil, fmt.Errorf("Error scanning book data: %w", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}
	return books, nil
}
