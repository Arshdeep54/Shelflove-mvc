package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

func BookAvailable(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("error parsing form")
			return
		}

		bookId := r.FormValue("bookid")
		id, err := strconv.ParseInt(bookId, 10, 64)
		if err != nil {
			fmt.Println("Error parsing to int")
			return
		}
		count, err := models.BookCount(id)
		if err != nil {
			fmt.Println("Error getting book count", err)
			return
		}
		if count > 0 {
			next(w, r)
			return
		}
		controllers.ErrorData.Message = "Not Enough books"
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}
