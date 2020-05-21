package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	port := ":8080"
	router := router()

	mongoDB()

	fmt.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}