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
func BookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/book.html", "templates/navbar.html"))
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

func AdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin.html", "templates/navbar.html"))
	return temp
}

func ErrorPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/error.html"))
	return temp
}
