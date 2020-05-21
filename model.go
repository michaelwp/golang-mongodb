package main

type User struct {
	_Id       int    `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}

type ResponseOne struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User
}