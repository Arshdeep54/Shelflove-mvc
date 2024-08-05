package controllers

import (
	"fmt"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/views"
)

type ErrorDataType struct {
	Message string
}

var ErrorData = &ErrorDataType{
	Message: "",
}

func Error(w http.ResponseWriter, r *http.Request) {
	t := views.ErrorPage()

	err := t.Execute(w, ErrorData)
	if err != nil {
		fmt.Print(err.Error())
	}
}
