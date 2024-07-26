package pkg

import (
	_ "github.com/mattn/go-sqlite3"
)

type UrlDb struct {
	Id          int
	OriginalUrl string
	ShortenUrl  string
}

func SetupDatabase() error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS URLS (original_url TEXT, short_url TEXT)")
	return err
}

func AddUrl(oU string, sU string) error {
	_, err := db.Exec(`INSERT INTO URLS (original_url, short_url) VALUES (?,?);`, oU, sU)
	return err
}

// additional code ofr extra purpose
// rows, err := db.Query("SELECT * FROM URLS")

// var originalUrl string
// var shortenUrl string

// for rows.Next() {
// 	rows.Scan(&originalUrl, &shortenUrl)
// 	fmt.Println(originalUrl + " " + shortenUrl)
// }

// func handleLinks(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	fmt.Print(r.FormValue("url"))
// 	// originalURL := r.FormValue("url")
// 	x, err := queryDbSearchShortUrl("https://mail.google.com/mail/u/0/#inbox")
// 	if err != nil {
// 		fmt.Print("found error", err)
// 	}

// 	fmt.Print("here is ", x)
// 	return
// }
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
