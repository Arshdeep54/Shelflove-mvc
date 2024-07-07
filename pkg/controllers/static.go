package controllers

import (
	"fmt"
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("static/style.css")
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "text/css")
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
	}
}
