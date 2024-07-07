package utils

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func ParseBook(book *types.Book, r *http.Request) error {
	var layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, r.FormValue("publication_date"))
	if err != nil {
		return err
	}
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Description = r.FormValue("description")
	book.PublicationDate = &parsedDate
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



