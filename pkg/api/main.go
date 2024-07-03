package api

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/middlewares"
	// "github.com/gorilla/mux"
)

func Start() {
	// r:= mux.NewRouter()
	//render Routes
	http.HandleFunc("/", middlewares.Authenticate(controllers.Home))
	http.HandleFunc("/books", middlewares.Authenticate(controllers.Books))
	http.HandleFunc("/books/{id}", middlewares.Authenticate(middlewares.ExistingIssue(controllers.Book)))
	http.HandleFunc("/login", middlewares.Authenticate(controllers.LoginPage))
	http.HandleFunc("/signup", middlewares.Authenticate(controllers.SignUpPage))
	http.HandleFunc("/user", middlewares.Authenticate(controllers.UserDashboard))
	http.HandleFunc("/admin", middlewares.Authenticate(controllers.AdminDashboard))
	http.HandleFunc("/error", controllers.Error)

	//api-routes

	//auth
	http.HandleFunc("/api/auth/login", controllers.Login)
	http.HandleFunc("/api/auth/signup", controllers.SignUp)
	http.HandleFunc("/api/auth/logout", controllers.Logout)

	//user
	http.HandleFunc("/api/user/issue/:bookid", controllers.IssueBook)
	http.HandleFunc("/api/user/return", controllers.ReturnBook)
	http.HandleFunc("/api/user/books", controllers.UserBooks)
	http.HandleFunc("/api/user/adminrequest", controllers.AdminRequest)

	//admin
	http.HandleFunc("/api/admin/addbook/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.AddBook)))
	http.HandleFunc("/api/admin/updatebook/{id}", middlewares.OnlyAdmin(controllers.UpdateBook))
	http.HandleFunc("/api/admin/deletebook/{bookId}", middlewares.Authenticate(middlewares.OnlyAdmin(middlewares.BookIsued(controllers.DeleteBook))))
	http.HandleFunc("/api/admin/bookissues", middlewares.OnlyAdmin(controllers.IssuedBooks))
	http.HandleFunc("/api/admin/approveissues", middlewares.OnlyAdmin(controllers.ApproveIssues))
	http.HandleFunc("/api/admin/denyIssue/{id}", middlewares.OnlyAdmin(controllers.DenyIssues))
	http.HandleFunc("/api/admin/approvereturns", middlewares.OnlyAdmin(controllers.ApproveReturns))
	http.HandleFunc("/api/admin/approveadmin", middlewares.OnlyAdmin(controllers.ApproveAdmin))
	http.HandleFunc("/api/admin/denyadmin/{id}", middlewares.OnlyAdmin(controllers.DenyAdmin))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error Starting the sever ...")
	}
}
