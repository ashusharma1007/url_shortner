package pkg

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", HandleForm)
	http.HandleFunc("POST /shorten", HandleShorten)
	http.HandleFunc("/short/", HandleRedirect)
	// http.HandleFunc("/Getlinks", handleLinks)

	fmt.Println("URL Shortener is running on :3030")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		// TODO: Handle specific errors.
		fmt.Println("server is not runnning", err)
	}
}
