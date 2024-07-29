package pkg

import (
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/", HandleForm)
	http.HandleFunc("POST /shorten", HandleShortenMaker(db))
	http.HandleFunc("/short/", HandleRedirect)
	// http.HandleFunc("/Getlinks", handleLinks)

	fmt.Println("URL Shortener is running on :3030")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		// TODO: Handle specific errors.
		fmt.Println("server is not runnning", err)
	}
}
