package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

func AdminRequestSent(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := models.AdminRequestSent(controllers.Data.UserId)
		if err != nil {
			fmt.Println(err)
			controllers.Data.AdminRequested = false
			next(w, r)
			return
		}
		if request {
			controllers.Data.AdminRequested = true
			next(w, r)
			return
		}
		controllers.Data.AdminRequested = false
		next(w, r)
	}
}
