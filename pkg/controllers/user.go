package controllers

import (
	"fmt"
	"net/http"
)

func IssueBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func ReturnBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func UserBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
func AdminRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}
