package models

import (
	"fmt"
	"strconv"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func GetAllBooks() ([]types.Book, error) {
	var books []types.Book
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	rows, err := db.Query(`SELECT * FROM book WHERE quantity >=0 ORDER BY quantity DESC;`)
	if err != nil {
		return nil, fmt.Errorf("error querying books: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var book types.Book
		// Scan each column of the current row into the Book struct fields
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationDate, &book.Quantity, &book.Genre, &book.Address, &book.Rating, &book.Description)
		if err != nil {
			return nil, fmt.Errorf("error scanning book data: %w", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return books, nil
}
func GetBook(bookId string) (*types.Book, error) {
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		return nil, err
	}
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	query := `SELECT * FROM book WHERE id = ? `
	row := db.QueryRow(query, id)
	var book types.Book
	err = row.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationDate, &book.Quantity, &book.Genre, &book.Address, &book.Rating, &book.Description)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func AddNewBook(book *types.Book) error {
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	query := `
        INSERT INTO book ( title, author, description, quantity ,genre,publication_date,rating, address)
        VALUES (?, ?, ?, ?,?,?,?,?)
      `
	_, err = db.Exec(query, book.Title, book.Author, book.Description, book.Quantity, book.Genre, book.PublicationDate, book.Rating, book.Address)
	if err != nil {
		return err
	}
	return nil
}
func DeleteBook(id int64) error {
	query := `UPDATE book SET quantity = -1 WHERE id= ? ;`
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("book Not Found")
	}
	return nil
}
func BookCount(id int64) (int, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}
	query := `
    SELECT Count(bookid)
    FROM issue
    WHERE bookid = ? and isReturned=false and issueRequested=false
  `
	row := db.QueryRow(query, id)
	var count int
	err=row.Scan(&count)
	if err!=nil{
		return 0, fmt.Errorf("error scaning: %w", err)
	}
	return count, nil
}
