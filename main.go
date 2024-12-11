package main

import (
	"database/sql"
	"log"
	"url-short/pkg"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	var db *sql.DB
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal(err)
	}
	pkg.SetupDatabase(db)
	pkg.RegisterRoutes(db)
}
