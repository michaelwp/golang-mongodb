package routers

import (
	"github.com/gorilla/mux"
	"golang-mongodb/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/user", controllers.ViewUser).Methods("GET")
	router.HandleFunc("/user", controllers.AddUser).Methods("POST")
	router.HandleFunc("/user/{firstname}", controllers.ViewUser).Methods("GET")
	router.HandleFunc("/user/email/{email}", controllers.FindUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")

	return router
}
