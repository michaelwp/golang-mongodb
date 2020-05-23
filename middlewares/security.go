package middlewares

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GenHash(password string) string {
	inpPass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(inpPass, bcrypt.DefaultCost)
	if err != nil {log.Fatal(err)}
	return string(hash)
}
