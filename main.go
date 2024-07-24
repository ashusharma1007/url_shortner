package main

import (
	"database/sql"
	"url-short/pkg"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./urls.db")
	if err != nil {
		panic(err)
	}
	pkg.SetupDatabase(db)
	pkg.RegisterRoutes()
}
