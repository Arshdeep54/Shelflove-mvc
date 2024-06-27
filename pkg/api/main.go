package api

import (
	"fmt"
	"net/http"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/gorilla/mux"
)

func Start(){
	r:= mux.NewRouter()
	r.HandleFunc("/",controllers.Home)

	err := http.ListenAndServe(":3000", r)
	if(err!=nil){
		fmt.Printf("Error Starting the sever ...")
	}
}