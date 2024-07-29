package middlewares

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
// 	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
// )

// func BookIssued(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		bookId := r.PathValue("bookId")
// 		id, err := strconv.ParseInt(bookId, 10, 64)
// 		if err != nil {
// 			fmt.Println("Error parsing to int")
// 			return
// 		}
// 		count,err:=models.IssuedBookCount(id)
// 		if err!=nil{
// 			fmt.Println("Error geeting book count",err)
// 			return
// 		}
// 		if count>0{
// 			controllers.ErrorData.Message="Book is Issued by some users, can't delete"
// 			http.Redirect(w,r,"/error",http.StatusSeeOther)
// 			return
// 		}
// 		next(w,r)
// 	}
// }
