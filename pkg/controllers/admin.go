package controllers

import (
	"fmt"
	"net/http"
  )
func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
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