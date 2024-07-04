package middlewares

import (
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

func ExistingIssue(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookId := r.PathValue("id")
		issue, err := models.GetIssue(bookId, controllers.Data.UserId)
		if err != nil {
			controllers.Data.IsIssued = false
			controllers.Data.IssueRequested = false
			controllers.Data.IsReturnRequested = false
			next(w, r)
			return
		}
		controllers.Data.IsIssued = !issue.IssueRequested
		controllers.Data.IssueRequested = issue.IssueRequested
		controllers.Data.IsReturnRequested = issue.ReturnRequested

		next(w, r)
	}
}
