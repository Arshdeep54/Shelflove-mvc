package utils

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func ParseBook(book *types.Book, r *http.Request) error {
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Description = r.FormValue("description")
	publication_date, err := time.Parse(LAYOUT, r.FormValue("publication_date"))
	if err != nil {
		return nil
	}
	book.PublicationDate = publication_date
	book.Genre = r.FormValue("genre")
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	if err != nil {
		return err
	}
	book.Rating = float32(rating)
	quantity, err := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
	if err != nil {
		return err
	}
	book.Quantity = int32(quantity)
	book.Address = r.FormValue("address")
	return nil
}
func ParseTime(str string) time.Time {
	date, err := time.Parse(LAYOUT, str)
	if err != nil {
		return time.Now()
	}
	return date
}
func ParseInt(str string) int {
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int(integer)
}
func ParseFloat(str string) float32 {
	float, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(float)

}
func GetDummyBooks() []types.Book {
	var books []types.Book
	for _, value := range DummyData {
		var book types.Book
		book.Title = value[0]
		book.Author = value[1]
		book.PublicationDate = ParseTime(value[2])
		book.Quantity = int32(ParseInt(value[3]))
		book.Genre = value[4]
		book.Description = value[5]
		book.Rating = ParseFloat(value[6])
		book.Address = value[7]
		books = append(books, book)
	}
	return books

}
