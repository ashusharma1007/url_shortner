package pkg

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandleForm(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML form
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<form method="post" action="/shorten">
				<input type="url" name="url" placeholder="Enter a URL" required>
				<input type="submit" value="Shorten">
			</form>
		</body>
		</html>
	`)
}
