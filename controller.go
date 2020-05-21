package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "home")
}

func addUser(w http.ResponseWriter, r *http.Request){

}
