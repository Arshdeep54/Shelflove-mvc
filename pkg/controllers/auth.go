package controllers

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	// "strings"
	// "github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logged in ")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Print("error parsing the form", err)
	}
	var (
		username  string
		email     string
		password  string
		password2 string
	)
	username = r.FormValue("username")
	email = r.FormValue("email")
	password = r.FormValue("password")
	password2 = r.FormValue("password2")

	if username == "" || email == "" || password == "" || password2 == "" {
		log.Fatal("Empty Fields")
		return
	}
	if password != password2 {
		log.Fatal("Passwords Didnt Match")
	}
	hashedPassword, err := utils.HashedPassword(password)
	if err != nil {
		log.Fatal("Error Hashing", err)
	}
	fmt.Println(hashedPassword)
	user := types.RegisterUserPayload{Username: username, Email: email, Password: hashedPassword}
	err = models.AddNewUser(&user)
	if err != nil {
		log.Fatal("Error adding new User", err)
	}

	fmt.Printf("Adding the user %s to the database\n", username)

}
