package main

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")

	return router
}
