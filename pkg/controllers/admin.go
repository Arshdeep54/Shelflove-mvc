package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Print("error parsing the form", err)
	}

	var book types.Book
	err = utils.ParseBook(&book, r)
	if err != nil {
		fmt.Print("error parsing the book", err)
	}
	if book.Title == "" || book.Address == " " || book.Description == "" || book.Quantity == 0 || book.PublicationDate == " " || book.Rating == 0 || book.Address == "" || book.Genre == "" {
		ErrorData.Message = "Empty fields"
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	id, already, err := models.BookStatus(book.Title)
	if err != nil {
		fmt.Println(err.Error())
		ErrorData.Message = "Error Adding book"
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	if already {
		payload := &types.RequestPayload{
			IssueIds:      []string{strconv.Itoa(id)},
			SelectedBooks: map[string]int{strconv.Itoa(id): int(book.Quantity)},
		}
		err = models.UpdatebooksQuantity(payload, true)
		if err != nil {
			fmt.Println(err.Error())
			ErrorData.Message = err.Error()
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	err = models.AddNewBook(&book)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing to int")
		return
	}
	var book types.Book
	err = utils.ParseBook(&book, r)
	if err != nil {
		fmt.Print("error parsing the book", err)
	}
	err = models.Updatebook(&book, int(id))
	if err != nil {
		fmt.Print("error updating the book", err)

	}
	http.Redirect(w, r, fmt.Sprintf("/books/%d", id), http.StatusSeeOther)
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
		fmt.Println("Error parsing to int", err)
		return
	}
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}


