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

// func queryDbSearchShortUrl(u string) (string, error) {
// 	db, err := sql.Open("sqlite3", "./urls.db")
// 	if err != nil {
// 		fmt.Print("Error Found", err)
// 	}
// 	rows := db.QueryRow("SELECT short_url FROM URLS WHERE original_url = ?", u)

// 	var shortenUrl string
// 	rows.Scan(&shortenUrl)
// 	return shortenUrl, nil
// }
