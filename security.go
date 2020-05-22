package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func genHash(password string) string {
	inpPass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(inpPass, bcrypt.DefaultCost)
	if err != nil {log.Fatal(err)}
	return string(hash)
}
