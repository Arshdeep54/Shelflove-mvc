package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	"github.com/joho/godotenv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Print("error parsing the form", err)
	}

	var (
		username string
		password string
	)

	username = r.FormValue("username")
	password = r.FormValue("password")
	u, err := models.GetUserbyUserName(username)
	err = signupRedirect(w, r, err, "/login", "Incorrect credentials")

	if err == nil {
		return
	}

	if !utils.ComparePasswords(u.Password, []byte(password)) {
		err = fmt.Errorf("incorrect credentials")
	} else {
		err = nil
	}

	err = signupRedirect(w, r, err, "/login", "Incorrect credentials")
	if err == nil {
		return
	}

	token, err := utils.JwtToken(types.JwtPayload{Id: u.Id, Email: u.Email, Username: u.Username, IsAdmin: u.IsAdmin, AdminRequest: u.AdminRequest})
	err = signupRedirect(w, r, err, "/login", "Error creating Jwt Token")

	if err == nil {
		return
	}

	cookie := http.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Path:     "/",
		MaxAge:   utils.Jwt_Expiration_Int(os.Getenv("JWT_EXPIRATION")),
		HttpOnly: true,
		Secure:   os.Getenv("NODE_ENV") == "development",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	if u.IsAdmin {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

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

	if password != password2 {
		err = fmt.Errorf("passwords don't match")
	} else {
		err = nil
	}

	err = signupRedirect(w, r, err, "/signup", "Passwords Didnt matched")
	if err == nil {
		return
	}

	if len(strings.Split(password, "")) <= utils.MINIMUM_PASSWORD_LENGTH {
		err = fmt.Errorf("password must be of more than %d characters", utils.MINIMUM_PASSWORD_LENGTH)
	} else {
		err = nil
	}

	err = signupRedirect(w, r, err, "/signup", fmt.Sprintf("password must be of more than %d characters", utils.MINIMUM_PASSWORD_LENGTH))
	if err == nil {
		return
	}

	hashedPassword, err := utils.HashedPassword(password)
	err = signupRedirect(w, r, err, "/signup", "Error Hashing Passwords")

	if err == nil {
		return
	}

	user := types.RegisterUserPayload{Username: username, Email: email, Password: hashedPassword}
	err = models.AddNewUser(&user)
	err = signupRedirect(w, r, err, "/signup", "Username or email already used")

	if err == nil {
		return
	}

	u, err := models.GetUserbyUserName(username)
	err = signupRedirect(w, r, err, "/signup", "Error Creating token")

	if err == nil {
		return
	}

	token, err := utils.JwtToken(types.JwtPayload{Id: u.Id, Email: u.Email, Username: u.Username, IsAdmin: u.IsAdmin, AdminRequest: u.AdminRequest})
	err = signupRedirect(w, r, err, "/login", "Error creating Jwt Token")

	if err == nil {
		return
	}

	cookie := http.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Path:     "/",
		MaxAge:   utils.Jwt_Expiration_Int(os.Getenv("JWT_EXPIRATION")),
		HttpOnly: true,
		Secure:   os.Getenv("NODE_ENV") == "development",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	if u.IsAdmin {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(map[string]string{"message": "Successfully logged out"})
	if err != nil {
		fmt.Println("Error encoding JSON response:", err)
	}
}

func signupRedirect(w http.ResponseWriter, r *http.Request, err error, redirectUrl string, message string) error {
	if err != nil {
		ErrorMessage = message

		fmt.Println(err)
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)

		return nil
	} else {
		ErrorMessage = ""
	}

	return fmt.Errorf("no Error to Redirect")
}
