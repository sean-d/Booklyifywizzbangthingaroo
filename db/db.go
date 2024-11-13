package db

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite", "api.db?_foreign_keys=on") //force foreign key as it's off by default in sqlite3

	if err != nil {
		panic("DB Connection Failed")
	}
	DB.SetMaxIdleConns(10) // creates a pool of persistent connections to prevent revolving door of opening and closing connections
	DB.SetMaxIdleConns(5)  // number of connections kept open if nothing is going on
	fmt.Println("DB CREATED")
}

func CreateTables() {
	createRegistrationTable := `
	    CREATE TABLE IF NOT EXISTS registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    event_id INTEGER NOT NULL,
	    user_id INTEGER NOT NULL,
	    FOREIGN KEY(event_id) REFERENCES events(id),
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create registration table")
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email TEXT NOT NULL UNIQUE,
    	password TEXT NOT NULL
)`

	_, err = DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	    CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    date_time DATETIME NOT NULL,
	    user_id INTEGER NOT NULL,
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

}
