package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error
	DB , err = sql.Open("sqlite3","api.db" )

	if err!= nil{
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
	
}

func createTables(){

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create Users table")
	}




	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT NOT NULL,
        date TEXT NOT NULL,
        location TEXT NOT NULL,
        user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE  --if user is deleted, then all his events are also deleted.
   
        );
	`
	_, err = DB.Exec(createEventsTable)

	if err!=nil{
		panic("Error creating events table")
	}


	createRegistrationsTable :=`
	CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        event_id INTEGER,
        user_id INTEGER,
        FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE  --if user or event is deleted, then all registrations for that user or event are also deleted.

	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	
	if err != nil{
		panic("Error creating registrations table")
	}
}

