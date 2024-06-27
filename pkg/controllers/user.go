package controllers
import (
	"fmt"
	"net/http"
  )
func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome from the main controller!")
}