package pkg

import (
	"fmt"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var urls = make(map[string]string)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := urls[shortKey]

	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "url not found")
		return
	}

	// Redirect the user to the original URLa
	// http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
