package main

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user", viewUser).Methods("GET")
	router.HandleFunc("/user", addUser).Methods("POST")
	router.HandleFunc("/user/{firstname}", viewUser).Methods("GET")
	router.HandleFunc("/user/email/{email}", findUser).Methods("GET")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")

	return router
}
