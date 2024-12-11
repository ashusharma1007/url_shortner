package pkg

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandleShortenMaker(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		originalURL := r.FormValue("url")
		if originalURL == "" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "url is missing")
			return
		}

		// Generate a unique shortened key for the original URL
		shortKey := GenerateShortKey()
		urls[shortKey] = originalURL

		// Construct the full shortened URL
		shortenedURL := fmt.Sprintf("http://localhost:3030/short/%s", shortKey)
		// add short url with original url in the table
		err := AddUrl(db, originalURL, shortenedURL)
		if err != nil {
			fmt.Println("problem in making entries in the table", err)
			return
		}
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
}
