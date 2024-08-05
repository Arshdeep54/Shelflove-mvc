package utils

import (
	"net/http"
	"strconv"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func ParseBook(book *types.Book, r *http.Request) error {
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Description = r.FormValue("description")
	book.PublicationDate = r.FormValue("publication_date")
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

func IssueStatus(request types.IssueWithDetails) string {
	switch {
	case !request.Issue.IssueRequested && !request.Issue.IsReturned && !request.Issue.ReturnRequested:
		return "Issued"
	case request.Issue.IssueRequested:
		return "Issue Requested"
	case request.Issue.ReturnRequested:
		return "Return Requested"
	case request.Issue.IsReturned:
		return "Returned"

	default:
		return " "
	}
}
