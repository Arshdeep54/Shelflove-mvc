package models

import (
	"fmt"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"strconv"
	"strings"
)

func GetAllBooks() ([]types.Book, error) {
	var books []types.Book
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	// rows, err := db.Query(`SELECT * FROM book WHERE quantity >=0 ORDER BY quantity DESC;`)
	rows, err := db.Table("books").Where("quantity >=0").Order("quantity desc").Rows()
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
	err = config.CloseConnection(db)
	if err != nil {
		return nil, err
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
	var book types.Book
	tx := db.First(&book, "id=?", id)

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, err
	}
	err = config.CloseConnection(db)
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
	tx := db.Model(&types.Book{}).Create(&book)
	if tx.Error != nil {
		return err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}

func Updatebook(book *types.Book, id int) error {
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	tx := db.Model(types.Book{}).Save(&types.Book{
		Id:              int32(id),
		Title:           book.Title,
		Author:          book.Author,
		Description:     book.Description,
		Quantity:        book.Quantity,
		PublicationDate: book.PublicationDate,
		Rating:          book.Rating,
		Genre:           book.Genre,
		Address:         book.Address})
	if tx.Error != nil {
		return tx.Error
	}

	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int64) error {
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	tx := db.Model(&types.Book{}).Where("id=?", id).Update("quantity", -1)
	if tx.Error != nil {
		return tx.Error
	}

	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}

func IssuedBookCount(id int64) (int, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}

	row := db.Model(&types.Issue{}).Where("book_id = ? and is_returned=false and issue_requested=false", id).Select("Count(book_id)").Row()
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error scaning: %w", err)
	}
	err = config.CloseConnection(db)
	if err != nil {
		return 0,err
	}
	return count, nil
}
func BookCount(id int64) (int, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}
	row := db.Model(&types.Book{}).Select("id", "quantity").Where("id = ?", id).Row()
	var userid, count int
	err = row.Scan(&userid, &count)
	if err != nil {
		return 0, fmt.Errorf("error scaning: %w", err)
	}
	err = config.CloseConnection(db)
	if err != nil {
		return 0, err
	}
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

	rowsAffected := db.Raw(query).RowsAffected
	if rowsAffected == 0 {
		fmt.Println("no updation")
		return fmt.Errorf("no Updation")
	}
	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil

}

func BookStatus(title string) (int, bool, error) {
	db, err := config.DbConnection()
	if err != nil {
		return 0, false, fmt.Errorf("error connecting to Db: %w", err)
	}
	row := db.Model(&types.Book{}).Where("title=?", title).Select("id").Row()
	var bookid int
	err = row.Scan(&bookid)
	if err != nil {
		return 0, false, nil
	}
	err = config.CloseConnection(db)
	if err != nil {
		return 0, false, err
	}
	return bookid, true, nil
}
