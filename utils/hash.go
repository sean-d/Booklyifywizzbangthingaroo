package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	/*
		Takes in string password, bytes it, salts/hashes it, returns string repr of this process and any errors
	*/

	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // cost is the complexity used

	return string(b), err

}
