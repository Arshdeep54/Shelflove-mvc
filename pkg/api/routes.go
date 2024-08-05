package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/middlewares"
)

func Start() {
	// serve static css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// render Routes
	http.HandleFunc("/", middlewares.Authenticate(controllers.Home))
	http.HandleFunc("/books", middlewares.Authenticate(controllers.Books))
	http.HandleFunc("/books/{id}", middlewares.Authenticate(controllers.Book))
	http.HandleFunc("/login", controllers.LoginPage)
	http.HandleFunc("/signup", controllers.SignUpPage)
	http.HandleFunc("/user", middlewares.Authenticate(controllers.UserDashboard))
	http.HandleFunc("/admin", middlewares.Authenticate(controllers.AdminDashboard))
	http.HandleFunc("/error", controllers.Error)

	// api-routes

	// auth
	http.HandleFunc("/api/auth/login", controllers.Login)
	http.HandleFunc("/api/auth/signup", controllers.SignUp)
	http.HandleFunc("/api/auth/logout", controllers.Logout)

	// user
	http.HandleFunc("/api/user/issue/{bookid}", middlewares.Authenticate(controllers.IssueBook))
	http.HandleFunc("/api/user/return/", middlewares.Authenticate(controllers.ReturnBook))
	http.HandleFunc("/api/user/adminrequest/", middlewares.Authenticate(controllers.AdminRequest))

	// admin
	http.HandleFunc("/api/admin/addbook/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.AddBook)))
	http.HandleFunc("/api/admin/updatebook/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.UpdateBook)))
	http.HandleFunc("/api/admin/deletebook/{bookId}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DeleteBook)))
	http.HandleFunc("/api/admin/approveissues/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.ApproveIssues)))
	http.HandleFunc("/api/admin/denyissue/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DenyIssue)))
	http.HandleFunc("/api/admin/approvereturns", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.ApproveReturns)))
	http.HandleFunc("/api/admin/denyreturn/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DenyReturn)))

	http.HandleFunc("/api/admin/approveadmin/", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.ApproveAdmin)))
	http.HandleFunc("/api/admin/denyadmin/{id}", middlewares.Authenticate(middlewares.OnlyAdmin(controllers.DenyAdmin)))

	const three = 3
	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: three * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Error Starting the sever ... %s", err))
	}
}
