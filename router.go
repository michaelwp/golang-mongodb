package main

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user", viewUser).Methods("GET")
	router.HandleFunc("/user", addUser).Methods("POST")

	return router
}
