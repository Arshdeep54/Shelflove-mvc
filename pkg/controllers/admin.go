package controllers

import (
	"encoding/json"
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

func ApproveIssues(w http.ResponseWriter, r *http.Request) {

	var payload types.RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	if payload.IssueIds == nil || len(payload.IssueIds) == 0 {
		http.Error(w, "Invalid request body: missing or invalid issue IDs", http.StatusBadRequest)
		return
	}
	err = models.UpdatebooksQuantity(&payload, false)
	if err != nil {
		fmt.Println("Error updating quantity:", err)
	}
	err = models.UpdateIssue(payload.IssueIds, utils.ISSUED)
	if err != nil {
		fmt.Println("Error updating issues:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(map[string]string{"message": "successfully approved"})
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func DenyIssues(w http.ResponseWriter, r *http.Request) {
	issueId := r.PathValue("id")
	id, err := strconv.ParseInt(issueId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing to int")
		return
	}
	err = models.DenyIssueRequest(int(id), utils.ISSUED)
	if err != nil {
		fmt.Println("Error Denying Request:", err)
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
func DenyReturn(w http.ResponseWriter, r *http.Request) {
	issueId := r.PathValue("id")
	id, err := strconv.ParseInt(issueId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing to int")
		return
	}
	err = models.DenyIssueRequest(int(id), utils.RETURNED)
	if err != nil {
		fmt.Println("Error Denying Request:", err)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
func ApproveReturns(w http.ResponseWriter, r *http.Request) {
	var payload types.RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	issueIds := payload.IssueIds
	if payload.IssueIds == nil || len(payload.IssueIds) == 0 {
		http.Error(w, "Invalid request body: missing or invalid issue IDs", http.StatusBadRequest)
		return
	}
	err = models.UpdatebooksQuantity(&payload, true)
	if err != nil {
		fmt.Println("Error updating quantity:", err)
	}
	err = models.UpdateIssue(issueIds, utils.RETURNED)
	if err != nil {
		fmt.Println("Error updating issues:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(map[string]string{"message": "successfully approved"})
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func ApproveAdmin(w http.ResponseWriter, r *http.Request) {
	type adminrequest struct {
		Ids []string
	}
	var request adminrequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if request.Ids == nil || len(request.Ids) == 0 {
		http.Error(w, "Invalid request body: missing or invalid  IDs", http.StatusBadRequest)
		return
	}

	err = models.ApproveAdmin(request.Ids)
	if err != nil {
		fmt.Println("Error approve admin:", err)
	}

}
func DenyAdmin(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing to int")
		return
	}
	err = models.DenyAdminRequest(int(id))
	if err != nil {
		fmt.Println("Error Denying Request:", err)
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
