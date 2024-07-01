package views

import (
	"html/template"
)



func HomePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/home.html", "templates/navbar.html"))
	return temp

}
func BooksPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/books.html", "templates/navbar.html"))
	return temp
}
func LoginPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/login.html", "templates/navbar.html"))
	return temp
}
func SignUpPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/signup.html", "templates/navbar.html"))
	return temp
}
