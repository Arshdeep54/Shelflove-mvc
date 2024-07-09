package controllers

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
)

func IssueBook(w http.ResponseWriter, r *http.Request) {
	existingIssuesCount, err := models.ExistingIssueCount(Data.UserId)
	if err != nil {
		fmt.Println("Error fetching existing issues", err)
		return
	}
	if existingIssuesCount >= utils.BOOK_ISSUE_LIMIT {
		ErrorData.Message = fmt.Sprintf("You have already issued %d books",existingIssuesCount)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	err = models.AddNewIssue(Data.UserId, int(Data.Book.Id))
	if err != nil {
		fmt.Println("Error fetching existing issues", err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/books/%d", Data.Book.Id), http.StatusSeeOther)
}
func ReturnBook(w http.ResponseWriter, r *http.Request) {
	err := models.ReturnRequest(Data.UserId, int(Data.Book.Id))
	if err != nil {
		fmt.Println("Error Updating issue",err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/books/%d", Data.Book.Id), http.StatusSeeOther)
}

func AdminRequest(w http.ResponseWriter, r *http.Request) {
	err := models.AdminRequest(Data.UserId)
	if err != nil {
		ErrorData.Message = "Failed to send Request or Already sent"
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)

}
