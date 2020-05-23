package main

import (
	"fmt"
	"golang-mongodb/db"
	"golang-mongodb/routers"
	"log"
	"net/http"
)

func main(){
	port := ":8080"
	router := routers.Router()

	db.MongoDB()

	fmt.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}