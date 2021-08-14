package database

import "database/sql"

var Db *sql.DB
var Err error

func Connect(connectionString string) {
	// Db, Err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/dbmusic")
	Db, Err = sql.Open("mysql", connectionString)
	if Err != nil {
		panic(Err.Error())
	}
}
