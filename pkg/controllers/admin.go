package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sf")
	err := r.ParseForm()
	if err != nil {
		fmt.Print("error parsing the form", err)
	}
	var book types.Book
	var layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, r.FormValue("publication_date"))
	if err != nil {
		fmt.Println("Error parsing publication date:", err)
		return
	}
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Description = r.FormValue("description")
	book.PublicationDate = &parsedDate
	book.Genre = r.FormValue("genre")
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	book.Rating = float32(rating)
	quantity, err := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	book.Quantity = int32(quantity)
	book.Address = r.FormValue("address")
	if book.Title == "" || book.Address == " " || book.Description == "" || book.Quantity == 0 || book.PublicationDate == nil || book.Rating == 0 || book.Address == "" || book.Genre == "" {
		ErrorData.Message = "Empty fields"
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	err = models.AddNewBook(&book)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	fmt.Println(book)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("bookId")
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing to int")
		return
	}
	err = models.DeleteBook(id)
	if err != nil {
		fmt.Println("Error parsing to int",err)
		return
	}
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
func IssuedBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func ApproveIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func DenyIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func ApproveReturns(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func ApproveAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func DenyAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
