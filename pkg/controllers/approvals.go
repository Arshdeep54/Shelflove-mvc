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

func DenyIssue(w http.ResponseWriter, r *http.Request) {
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
		Ids []string `json:"ids"`
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

	err = models.BalanceIssues(request.Ids)
	if err != nil {
		fmt.Println(err.Error())
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
