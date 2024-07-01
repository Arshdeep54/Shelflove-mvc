package controllers

import (
	"fmt"

	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/views"
)

type Data struct {
	IsAdmin      bool
	IsLoggedIn   bool
	Books        []types.Book
	ErrorMessage string
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := views.HomePage()
	data := Data{
		IsAdmin:      false,
		IsLoggedIn:   true,
		Books:        nil,
		ErrorMessage: "",
	}
	err := t.Execute(w, data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func Books(w http.ResponseWriter, r *http.Request) {
	t := views.BooksPage()
	books, err := models.GetAllBooks()
	fmt.Println(len(books))
	if err != nil {
		fmt.Print(err.Error())
	}
	data := &Data{
		IsAdmin:      false,
		IsLoggedIn:   true,
		Books:        books,
		ErrorMessage: "",
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func Book(w http.ResponseWriter, r *http.Request) {

}
func UserDashboard(w http.ResponseWriter, r *http.Request) {

}
func AdminDashboard(w http.ResponseWriter, r *http.Request) {

}
func LoginPage(w http.ResponseWriter, r *http.Request) {
	// isLoggedIn: false,
	// errorMessage: null,
	t := views.LoginPage()
	data := &Data{
		IsAdmin:      false,
		IsLoggedIn:   false,
		Books:        nil,
		ErrorMessage: "",
	}
	err := t.Execute(w, data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	t := views.SignUpPage()
	data := &Data{
		IsAdmin:      false,
		IsLoggedIn:   false,
		Books:        nil,
		ErrorMessage: "",
	}
	err := t.Execute(w, data)
	if err != nil {
		fmt.Print(err.Error())
	}
}
