package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

/*
	=====================================================
	Home [GET]
	http://localhost:8080
	=====================================================
*/
func home(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "home")
}

/*
	=====================================================
	Add User [POST]
	http://localhost:8080/user
	request body : {
		"first_name": "John",
		"last_name": "Doe",
		"email": "john.doe@mail.com",
		"password": "password"
	}
	=====================================================
*/
func addUser(w http.ResponseWriter, r *http.Request){
	var u User
	var resp Response

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)

		resp.Status = 0
		resp.Message = "Data failed added"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		mongoDB := mongoDB()

		_, err := mongoDB.Collection("tbl_user").InsertOne(context.TODO(), u)
		if err != nil {
			log.Println(err)

			resp.Status = 0
			resp.Message = "Data failed added"
			w.WriteHeader(http.StatusBadRequest)
		} else {
			resp.Status = 1
			resp.Message = "Data successfully added"
			w.WriteHeader(http.StatusOK)
		}
	}

	_ = json.NewEncoder(w).Encode(resp)
}

/*
	=====================================================
	View Users [GET]
	All users : http://localhost:8080/user
	spesific user : http://localhost:8080/user/{email}
	=====================================================
*/
func viewUser(w http.ResponseWriter, r *http.Request){
	var u []User
	var resp Response

	w.Header().Set("Content-type", "application/json")

	mongoDB := mongoDB()

	cur, err := mongoDB.Collection("tbl_user").Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var el = User{}
		err := cur.Decode(&el)
		if err != nil {
			log.Fatal(err)
		}

		u = append(u, el)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	resp.Status = 1
	resp.Message = "List of User"
	resp.Data = u

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
