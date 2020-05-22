package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
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
	errMsg := "Data failed to register"

	w.Header().Set("Content-Type", "application/json")

	resp.Status = 1
	resp.Message = "Data successfully added"

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)

		resp.Status = 0
		resp.Message = errMsg
		w.WriteHeader(http.StatusBadRequest)
	} else {
		uRes := findEmail(u.Email)

		if uRes.Email != "" {
			resp.Status = 0
			resp.Message = "Email already registered"
			w.WriteHeader(http.StatusBadRequest)
		} else {
			mongoDB := mongoDB()

			pass := genHash(u.Password)

			inputData := bson.D{
				{"firstname", strings.ToLower(u.FirstName)},
				{"lastname", strings.ToLower(u.LastName)},
				{"email", strings.ToLower(u.Email)},
				{"password", pass},
			}

			_, err := mongoDB.Collection("tbl_user").InsertOne(context.TODO(), inputData)
			if err != nil {
				log.Println(err)

				resp.Status = 0
				resp.Message = errMsg
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}

	_ = json.NewEncoder(w).Encode(resp)
}

/*
	=====================================================
	View Users [GET]
	all users : http://localhost:8080/user
	spesific user : http://localhost:8080/user/{firstname}
	=====================================================
*/
func viewUser(w http.ResponseWriter, r *http.Request){
	var u []User
	var resp Response
	vars := mux.Vars(r)
	filter := bson.D{{"firstname", strings.ToLower(vars["firstname"])}}

	if len(vars) <= 0 {filter = bson.D{{}}}

	w.Header().Set("Content-type", "application/json")

	mongoDB := mongoDB()

	cur, err := mongoDB.Collection("tbl_user").Find(context.TODO(), filter)
	if err != nil {log.Fatal(err)}

	for cur.Next(context.TODO()) {
		var el = User{}
		err := cur.Decode(&el)
		if err != nil {
			log.Fatal(err)
		}

		u = append(u, el)
	}

	if err := cur.Err(); err != nil {log.Fatal(err)}

	_ = cur.Close(context.TODO())

	resp.Status = 1
	resp.Message = "List of User"
	resp.Data = u

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

/*
	=====================================================
	Find User [GET]
	http://localhost:8080/user/{email}
	=====================================================
*/
func findUser(w http.ResponseWriter, r *http.Request){
	var resp ResponseOne
	vars := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	uRes := findEmail(vars["email"])

	resp.Status = 1
	resp.Message = "User data"
	resp.Data = uRes

	if uRes.Email == "" {
		resp.Status = 0
		resp.Message = "Data not found"
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

/*
	=====================================================
	Find Email (for checking if email already registered)
	=====================================================
*/
func findEmail(email string) User {
	var uRes User
	mongoDB := mongoDB()

	filter := bson.M{"email": strings.ToLower(email)}
	err := mongoDB.Collection("tbl_user").FindOne(context.TODO(), filter).Decode(&uRes)
	if err != nil {log.Println(err)}

	return uRes
}

/*
	=====================================================
	Delete User [DELETE]
	http://localhost:8080/user/{id}
	=====================================================
*/
func deleteUser(w http.ResponseWriter, r *http.Request){
	var resp Response
	vars := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	resp.Status = 1
	resp.Message = "User successfully deleted"

	idPrimitive, err := primitive.ObjectIDFromHex(vars["id"])

	filter := bson.M{"_id": idPrimitive}

	mongoDB := mongoDB()
	delResult, err := mongoDB.Collection("tbl_user").DeleteOne(context.TODO(), filter)
	if err != nil {log.Fatal(err)}

	if delResult.DeletedCount <= 0 {
		resp.Status = 0
		resp.Message = "User failed to delete"
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

/*
	=====================================================
	Update User [PUT]
	http://localhost:8080/user/{id}
	request body : {
		"first_name": "John",
		"last_name": "Doe",
	}
	=====================================================
*/
func updateUser(w http.ResponseWriter, r *http.Request){
	var u User
	var resp Response
	vars := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {log.Fatal(err)}

	resp.Status = 1
	resp.Message = "User successfully updated"

	idPrimitive, err := primitive.ObjectIDFromHex(vars["id"])

	filter := bson.M{"_id": idPrimitive}
	update := bson.M{"$set": bson.M{
		"firstname": strings.ToLower(u.FirstName),
		"lastname":  strings.ToLower(u.LastName),
	},}

	mongoDB := mongoDB()
	updResult, err := mongoDB.Collection("tbl_user").UpdateOne(context.TODO(), filter, update)
	if err != nil {log.Fatal(err)}

	if updResult.ModifiedCount <= 0 {
		resp.Status = 0
		resp.Message = "User failed to update"
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(resp)
}
