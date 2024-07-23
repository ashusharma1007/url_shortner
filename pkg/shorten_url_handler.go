package pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./urls.db")
	if err != nil {
		fmt.Println(errors.New("problem in starting the db"), err)
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		fmt.Print(errors.New("url paramter is missing"))
		return
	}

	// Generate a unique shortened key for the original URL
	shortKey := GenerateShortKey()
	urls[shortKey] = originalURL

	// Construct the full shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:3030/short/%s", shortKey)
	// add short url with original url in the table
	AddUrl(db, originalURL, shortenedURL)
	// Serve the result page

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<p>Original URL: `, originalURL, `</p>
			<p>Shortened URL: <a href="`, shortenedURL, `">`, shortenedURL, `</a></p>
		</body>
		</html>
	`)
}
