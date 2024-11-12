package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// for creation and verification of tokens
const secretKey = "superSecretKeyBBQ"

func GenerateToken(email string, userId int64) (string, error) {
	/*
		Create new jwt:
			HS256 signing method
			map of claims
				email/userId passed in
				token expires 2 hours from creation time
		returns generated token and/or error
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	// key has to be a slice of bytes despite the docs saying of type any.
	// be better devs...your docs are not good.
	return token.SignedString([]byte(secretKey))
}
