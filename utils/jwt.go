package utils

import (
	"errors"
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

func VerifyToken(t string) (int64, error) {
	/*
		verify supplied token is valid.
		returns userId and any errors.

		jwt.Parse:
			takes in the supplied token, t
			takes in a function that automatically receives the token from jwt package, does its magic, and returns the secret
			key constant to the jwt.Parse function

			we check to make sure that the method used to sign the token is of type *jwt.SigningMethodHMAC
			jwt.NewWithClaim() above is using HS256 so we are simply verifying that the token being provided
			is signed by the expected method type.


			returns the parsedToken and any errors

		we check to ensure the parsedToken is valid.\
		we check to ensure the claims are of type jwt.MapClaims.
			since we created the tokens with that as the claim type, this assertion needs to be true.
			returns the claims and a bool for success

		since claims is simply a map, we can easily get the email and userId from it.
		mapClaims: values are of type any. so we type check to ensure email and userId are the correct types


	*/

	parsedToken, err := jwt.Parse(t, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		// return a slice of bytes...again, would be nice if the documentation let you know this.
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token")
	}

	//email, ok := claims["email"].(string)
	//
	//if !ok {
	//	return errors.New("Unable to extract email from token")
	//}

	// because it's so much fun to get a float64 back for a number that can never be anything other than an int....
	// whoever designed this should be kneecapped.
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
