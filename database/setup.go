package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "user=admin dbname=platform password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
