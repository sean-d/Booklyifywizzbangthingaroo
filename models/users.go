package models

import "github.com/sean-d/Booklyifywizzbangthingaroo/db"

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

	result, err := statement.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	u.ID = userID
	return err

}
