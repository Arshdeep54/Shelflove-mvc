package api

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/middlewares"
)

func Start() {

	//serve static css
	http.HandleFunc("/static/", controllers.StaticHandler)

	//render Routes
	http.HandleFunc("/", middlewares.Authenticate(controllers.Home))
	http.HandleFunc("/books", middlewares.Authenticate(controllers.Books))
	http.HandleFunc("/books/{id}", middlewares.Authenticate(middlewares.ExistingIssue(controllers.Book)))
	http.HandleFunc("/login", middlewares.Authenticate(controllers.LoginPage))
	http.HandleFunc("/signup", middlewares.Authenticate(controllers.SignUpPage))
	http.HandleFunc("/user", middlewares.Authenticate(middlewares.AdminRequestSent(controllers.UserDashboard)))
	http.HandleFunc("/admin", middlewares.Authenticate(controllers.AdminDashboard))
	http.HandleFunc("/error", controllers.Error)

	//api-routes

	//auth
	http.HandleFunc("/api/auth/login", controllers.Login)
	http.HandleFunc("/api/auth/signup", controllers.SignUp)
	http.HandleFunc("/api/auth/logout", controllers.Logout)

	//user
	http.HandleFunc("/api/user/issue/{bookid}", middlewares.Authenticate(middlewares.BookAvailable(controllers.IssueBook)))
	http.HandleFunc("/api/user/return/", middlewares.Authenticate(controllers.ReturnBook))
	http.HandleFunc("/api/user/adminrequest/", middlewares.Authenticate(controllers.AdminRequest))

	//admin
	http.HandleFunc("/api/admin/addbook/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.AddBook)))
	http.HandleFunc("/api/admin/updatebook/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.UpdateBook)))
	http.HandleFunc("/api/admin/deletebook/{bookId}", middlewares.Authenticate(middlewares.OnlyAdmin(middlewares.BookIssued(controllers.DeleteBook))))
	http.HandleFunc("/api/admin/approveissues/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.ApproveIssues)))
	http.HandleFunc("/api/admin/denyissue/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DenyIssues)))
	http.HandleFunc("/api/admin/approvereturns", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.ApproveReturns)))
	http.HandleFunc("/api/admin/denyreturn/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DenyReturn)))

	http.HandleFunc("/api/admin/approveadmin/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.ApproveAdmin)))
	http.HandleFunc("/api/admin/denyadmin/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DenyAdmin)))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error Starting the sever ...", err)
	}
}
