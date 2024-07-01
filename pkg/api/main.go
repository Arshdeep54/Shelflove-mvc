package api

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	// "github.com/gorilla/mux"
)

func Start(){
	// r:= mux.NewRouter()
	//render Routes
	http.HandleFunc("/",controllers.Home)
    http.HandleFunc("/books",controllers.Books)
    http.HandleFunc("/books/:id",controllers.Book)
	http.HandleFunc("/login",controllers.LoginPage)
	http.HandleFunc("/signup",controllers.SignUpPage)
	http.HandleFunc("/user",controllers.UserDashboard)
	http.HandleFunc("/admin",controllers.AdminDashboard)

	//api-routes

	//auth
	http.HandleFunc("/api/auth/login",controllers.Login)
	http.HandleFunc("/api/auth/signup",controllers.SignUp)

	//user
	http.HandleFunc("/api/user/issue/:bookid",controllers.IssueBook) 
	http.HandleFunc("/api/user/return",controllers.ReturnBook) 
	http.HandleFunc("/api/user/books",controllers.UserBooks) 
	http.HandleFunc("/api/user/adminrequest",controllers.AdminRequest) 

	//admin
	http.HandleFunc("/api/admin/addbook",controllers.AddBook) 
	http.HandleFunc("/api/admin/updatebook/:id",controllers.UpdateBook) 
	http.HandleFunc("/api/admin/deletebook/:bookId",controllers.DeleteBook) 
	http.HandleFunc("/api/admin/bookissues",controllers.IssuedBooks) 
	http.HandleFunc("/api/admin/approveissues",controllers.ApproveIssues) 
	http.HandleFunc("/api/admin/denyIssue/:id",controllers.DenyIssues) 
	http.HandleFunc("/api/admin/approvereturns",controllers.ApproveReturns) 
	http.HandleFunc("/api/admin/approveadmin",controllers.ApproveAdmin) 
	http.HandleFunc("/api/admin/denyadmin/:id",controllers.DenyAdmin) 


	

	err := http.ListenAndServe(":3000", nil)
	if(err!=nil){
		fmt.Printf("Error Starting the sever ...")
	}
}