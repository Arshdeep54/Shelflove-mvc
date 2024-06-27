package controllers

import (
	"fmt"
	"net/http"
  )
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}