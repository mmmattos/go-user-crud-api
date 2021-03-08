package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// type Utils struct {
// }

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
