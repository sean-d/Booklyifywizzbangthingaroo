package models

import (
	"errors"
	"github.com/sean-d/Booklyifywizzbangthingaroo/db"
	hash "github.com/sean-d/Booklyifywizzbangthingaroo/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() (err error) {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)
	defer statement.Close()

	if err != nil {
		return err
	}

	hashedPassword, err := hash.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	u.ID = userID
	return err

}

func (u *User) ValidateCredentials() (err error) {
	/*
		validate the credentials provided are ... valid

		we query a single row as we know each email is unique
		we take what is returned from the query and assign it to the helper variable, retrieved...

		Scan returns an error if no rows matched query

	*/

	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retreivedPassword string
	err = row.Scan(&retreivedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := hash.CheckPasswordHash(u.Password, retreivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil
}
