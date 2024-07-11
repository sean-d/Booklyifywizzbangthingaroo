package db

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("DB Connection Failed")
	}
	DB.SetMaxIdleConns(10) // creates a pool of persistent connections to prevent revolving door of opening and closing connections
	DB.SetMaxIdleConns(5)  // number of connections kept open if nothing is going on
	fmt.Println("DB CREATED")
}

func CreateTables() {
	createEventsTable := `
	    CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    date_time DATETIME NOT NULL,
	    user_id INTEGER NOT NULL
	)`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}
}
