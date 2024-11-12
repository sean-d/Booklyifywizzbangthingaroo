package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	/*
		Takes in string password, bytes it, salts/hashes it,
		returns string repr of this process and any errors
	*/

	// cost is the complexity used
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(b), err

}

func CheckPasswordHash(password, hashedPassword string) bool {
	/*
		compared the saved password with the user submitted one
	*/
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	// will be true if no error, false if error. more terse approach
	// return err == nil
	if err != nil {
		return false
	}
	return true
}
