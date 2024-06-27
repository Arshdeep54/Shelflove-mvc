package controllers

import (
	"net/http"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/views"
)
type Data struct{
	IsLoggedIn bool
	IsAdmin bool
}
func Home(w http.ResponseWriter, r *http.Request) {
	t := views.HomePage()
	data := Data{
		IsLoggedIn: true, // Replace with your logic to determine logged in state
		IsAdmin:    false, // Replace with your logic to determine admin status
	}
	t.Execute(w, data)
}