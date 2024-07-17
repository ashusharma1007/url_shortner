package pkg

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", HandleForm)
	http.HandleFunc("/shorten", HandleShorten)
	http.HandleFunc("/short/", HandleRedirect)
	// http.HandleFunc("/Getlinks", handleLinks)

	fmt.Println("URL Shortener is running on :3030")
	http.ListenAndServe(":3030", nil)
}
