package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
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

		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationDate, &book.Quantity, &book.Genre, &book.Address, &book.Rating, &book.Description)
		if err != nil {
			return nil, fmt.Errorf("error scanning book data: %w", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	db.Close()

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

	var publication_date *time.Time

	err = row.Scan(&book.Id, &book.Title, &book.Author, &publication_date, &book.Quantity, &book.Genre, &book.Description, &book.Rating, &book.Address)
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	book.PublicationDate = publication_date.Format(utils.LAYOUT)

	db.Close()

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

	db.Close()

	return nil
}

func Updatebook(book *types.Book, id int) error {
	query := `UPDATE book
        SET title = COALESCE(?, title),
            author = COALESCE(?, author),
            description = COALESCE(?, description),
            quantity = COALESCE(?, quantity),
            publication_date = COALESCE(?, publication_date),
            rating = COALESCE(?, rating),
            genre = COALESCE(?, genre),
            address = COALESCE(?, address)
        WHERE id = ?;`

	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}

	result, err := db.Exec(query, book.Title, book.Author, book.Description, book.Quantity, book.PublicationDate, book.Rating, book.Genre, book.Address, id)
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

	db.Close()

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

	db.Close()

	return nil
}

func IssuedBookCount(id int64) (int, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}

	query := `
    SELECT Count(book_id)
    FROM issue
    WHERE book_id = ? and isReturned=false and issueRequested=false
  `
	row := db.QueryRow(query, id)

	var count int

	err = row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error scaning: %w", err)
	}

	db.Close()

	return count, nil
}

func BookCount(id int64) (int, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}

	query := ` SELECT id, quantity
        FROM book
        WHERE id = ?
      `
	row := db.QueryRow(query, id)

	var userid, count int

	err = row.Scan(&userid, &count)
	if err != nil {
		return 0, fmt.Errorf("error scaning: %w", err)
	}

	db.Close()

	return count, nil
}

func UpdatebooksQuantity(payload *types.RequestPayload, increase bool) error {
	var sign string
	if increase {
		sign = "+"
	} else {
		sign = "-"
	}

	query := `UPDATE book SET quantity= (CASE `
	for key, value := range payload.SelectedBooks {
		query += fmt.Sprintf(" WHEN id= %s THEN quantity %s %d", key, sign, value)
	}

	var keyString string
	for key := range payload.SelectedBooks {
		keyString += key
		keyString += ","
	}

	keyString = strings.Trim(keyString, ",")
	query += fmt.Sprintf(` ELSE (quantity) END ) WHERE id IN (%s)`, keyString)
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}

	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("no updation")

		return fmt.Errorf("no Updation")
	}

	db.Close()

	return nil
}

func BookStatus(title string) (int, bool, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, false, fmt.Errorf("error connecting to Db: %w", err)
	}

	query := ` SELECT id
        FROM book
        WHERE title = ?
		LIMIT 1;
      `
	row := db.QueryRow(query, title)

	var bookid int

	err = row.Scan(&bookid)
	if err != nil {
		return 0, false, err
	}

	db.Close()

	return bookid, true, nil
}
