package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
)

func OnlyAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("here")
		if controllers.Data.IsAdmin && r.Method == http.MethodPost {
			fmt.Print("also")
			next(w, r)

			return
		} else {
			controllers.ErrorData.Message = "You are not the admin "
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
	}
}
