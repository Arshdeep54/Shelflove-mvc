package controllers

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/views"
)

var Data = types.RenderData{
	IsAdmin:           false,
	IsLoggedIn:        false,
	UserId:            0,
	Books:             nil,
	Book:              nil,
	ErrorMessage:      "",
	IssueRequested:    false,
	IsIssued:          false,
	IsReturnRequested: false,
	AdminRequested:    false,
	Username:          "",
	Email:             "",
	HomeActive:        false,
	BooksActive:       false,
}

var ErrorMessage string

func Home(w http.ResponseWriter, r *http.Request) {
	t := views.HomePage()
	Data.HomeActive = true
	Data.BooksActive = false

	err := t.Execute(w, Data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func Books(w http.ResponseWriter, r *http.Request) {
	t := views.BooksPage()

	books, err := models.GetAllBooks()
	if err != nil {
		fmt.Print(err.Error())
	}

	Data.Books = books
	Data.HomeActive = false
	Data.BooksActive = true

	err = t.Execute(w, Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Book(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")

	db, err := config.DbConnection()
	if err != nil {
		fmt.Print("error connecting to Db: %w", err)
	}

	issue, err := models.GetIssue(db, bookId, Data.UserId)
	if err != nil {
		Data.IsIssued = false
		Data.IssueRequested = false
		Data.IsReturnRequested = false
	} else {
		Data.IsIssued = !issue.IssueRequested
		Data.IssueRequested = issue.IssueRequested
		Data.IsReturnRequested = issue.ReturnRequested
	}

	t := views.BookPage()

	book, err := models.GetBook(bookId)
	if err != nil {
		fmt.Print(err.Error())
	}

	Data.Book = book
	Data.HomeActive = false
	Data.BooksActive = false

	err = t.Execute(w, Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UserDashboard(w http.ResponseWriter, r *http.Request) {
	t := views.UserDashboardPage()

	request, err := models.AdminRequestSent(Data.UserId)
	if err != nil {
		fmt.Println(err)

		Data.AdminRequested = false
	}

	if request {
		Data.AdminRequested = true
	}

	userIssues, err := models.GetUserIssues(Data.UserId)
	if err != nil {
		fmt.Print(err.Error())
	}

	Data.HomeActive = false
	Data.BooksActive = false

	Data.RequestedIssues = userIssues

	err = t.Execute(w, Data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if Data.IsAdmin {
		t := views.AdminPage()

		requestAdmin, err := models.GetAdminRequest()
		if err != nil {
			fmt.Print(err.Error())
		}

		requestedReturns, requestedIssues, err := models.GetRequestedAll()
		if err != nil {
			fmt.Print(err.Error())
		}

		Data.RequestedAdmins = requestAdmin
		Data.RequestedIssues = requestedIssues
		Data.RequestedReturns = requestedReturns
		Data.HomeActive = false
		Data.BooksActive = false

		err = t.Execute(w, Data)
		if err != nil {
			fmt.Print(err.Error())
		}
	} else {
		ErrorData.Message = "You are not the admin "

		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t := views.LoginPage()

	Data.ErrorMessage = ErrorMessage
	Data.HomeActive = false
	Data.BooksActive = false

	err := t.Execute(w, Data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	Data.ErrorMessage = ErrorMessage
	Data.HomeActive = false
	Data.BooksActive = false
	t := views.SignUpPage()

	err := t.Execute(w, Data)
	if err != nil {
		fmt.Print(err.Error())
	}
}
