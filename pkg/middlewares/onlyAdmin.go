package middlewares

import (
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
)

func OnlyAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if controllers.Data.IsAdmin && r.Method == http.MethodPost {
			next(w, r)
			return
		} else {
			controllers.ErrorData.Message = "You are not the admin "
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
	}
}
