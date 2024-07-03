package controllers

import (
	"fmt"

	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
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
	Username:          "",
	Email:             "",
}

var ErrorMessage string

func Home(w http.ResponseWriter, r *http.Request) {

	t := views.HomePage()

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
	err = t.Execute(w, Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func Book(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	t := views.BookPage()
	book, err := models.GetBook(bookId)
	if err != nil {
		fmt.Print(err.Error())
	}
	Data.Book = book
	err = t.Execute(w, Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func UserDashboard(w http.ResponseWriter, r *http.Request) {

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
	err := utils.UpdateHeaders(w, &Data)
	if err != nil {
		fmt.Println("Error updating Data", err)
	}
	t := views.LoginPage()

	Data.ErrorMessage = ErrorMessage

	err = t.Execute(w, Data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	Data.ErrorMessage = ErrorMessage
	t := views.SignUpPage()
	err := t.Execute(w, Data)
	if err != nil {
		fmt.Print(err.Error())
	}
}
